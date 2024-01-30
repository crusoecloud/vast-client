package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	swagger "gitlab.com/crusoeenergy/island/vast/pkg/vast"
)

func main() {
	t := http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	tt := NewAuthenticatingTransport(&t,
		"admin",
		"123456",
		"https://172.25.51.191/api/token/",
	)

	client := swagger.NewAPIClient(&swagger.Configuration{
		BasePath: "https://172.25.51.191/api/v2",
		HTTPClient: &http.Client{
			Transport: tt,
		},
	})

	list, _, err := client.TenantsApi.TenantsList(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(list)

	return
}
