package swagger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// arbitrary buffer for fetching a new access token slightly before expiration.
	expirationDurationBuffer = time.Minute
	// access token validity duration, per VAST Data's docs
	accessTokenDuration = time.Hour
)

// AuthenticatingTransport is an http.Transport that will authenticate requests by inserting an access token header,
// fetching new access tokens as needed.
type AuthenticatingTransport struct {
	username       string        // username to use for authentication
	password       string        // password to use for authentication
	authAPIPath    string        // timestamp after which last auth token should be expired
	accessToken    string        // most recently fetched auth token
	tokenDuration  time.Duration // how long until each new token expires
	tokenExpiresAt time.Time     // time after which accessToken is expired
	http.RoundTripper
}

type tokenResponseBody struct {
	Access  string
	Refresh string
}

func NewAuthenticatingTransport(r http.RoundTripper, username, password, authAPIPath string) *AuthenticatingTransport {
	if r == nil {
		r = http.DefaultTransport
	}

	return &AuthenticatingTransport{
		username:      username,
		password:      password,
		authAPIPath:   authAPIPath,
		tokenDuration: accessTokenDuration,
		RoundTripper:  r,
	}
}

func (t *AuthenticatingTransport) authenticate() error {
	body, err := json.Marshal(map[string]string{
		"username": t.username,
		"password": t.password,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal auth token request body: %w", err)
	}

	bodyBuf := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, t.authAPIPath, bodyBuf)
	if err != nil {
		return fmt.Errorf("failed to create auth token request: %w", err)
	}
	req.Header.Set("content-type", "application/json")

	resp, err := t.RoundTrip(req)
	if err != nil {
		return fmt.Errorf("auth token request failed: %w", err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body from auth token request response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("auth token request received status code %d. body: %s", resp.StatusCode, string(respBody))
	}

	var tokenResp tokenResponseBody
	err = json.Unmarshal(respBody, &tokenResp)
	if err != nil {
		return fmt.Errorf("failed to unmarshal auth token response: %w", err)
	}

	t.accessToken = tokenResp.Access
	t.tokenExpiresAt = time.Now().Add(t.tokenDuration - expirationDurationBuffer)

	return nil
}

func (t *AuthenticatingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	// fetch new token if likely expired
	if time.Now().After(t.tokenExpiresAt) && r.URL.String() != t.authAPIPath {
		err := t.authenticate()
		if err != nil {
			return nil, fmt.Errorf("failed to fetch auth token: %w", err)
		}
	}

	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.accessToken))

	//nolint:wrapcheck // error should be forwarded here.
	return t.RoundTripper.RoundTrip(r)
}
