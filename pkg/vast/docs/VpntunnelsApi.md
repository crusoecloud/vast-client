# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VpntunnelsCreate**](VpntunnelsApi.md#VpntunnelsCreate) | **Post** /vpntunnels/ | Create VPN Tunnel
[**VpntunnelsDelete**](VpntunnelsApi.md#VpntunnelsDelete) | **Delete** /vpntunnels/{id}/ | Close and Delete VPN Tunnel
[**VpntunnelsDeleteAll**](VpntunnelsApi.md#VpntunnelsDeleteAll) | **Delete** /vpntunnels/delete_all/ | Close and Delete all available VPN Tunnels
[**VpntunnelsList**](VpntunnelsApi.md#VpntunnelsList) | **Get** /vpntunnels/ | List VPN Tunnels
[**VpntunnelsRead**](VpntunnelsApi.md#VpntunnelsRead) | **Get** /vpntunnels/{id}/ | Return Details of a VPN Tunnel

# **VpntunnelsCreate**
> VpnTunnel VpntunnelsCreate(ctx, optional)
Create VPN Tunnel

This endpoint creates a VPN tunnel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VpntunnelsApiVpntunnelsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VpntunnelsApiVpntunnelsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VpntunnelsBody**](VpntunnelsBody.md)|  | 

### Return type

[**VpnTunnel**](VpnTunnel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VpntunnelsDelete**
> VpntunnelsDelete(ctx, id)
Close and Delete VPN Tunnel

This endpoint closes and deletes an open VPN tunnel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| VPN tunnel ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VpntunnelsDeleteAll**
> VpntunnelsDeleteAll(ctx, )
Close and Delete all available VPN Tunnels

This endpoint closes and deletes all open VPN tunnels

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VpntunnelsList**
> []VpnTunnel VpntunnelsList(ctx, )
List VPN Tunnels

This endpoint lists VPN tunnels.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]VpnTunnel**](VpnTunnel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VpntunnelsRead**
> VpnTunnel VpntunnelsRead(ctx, id)
Return Details of a VPN Tunnel

This endpoint returns details of a VPN tunnel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| VPN tunnel ID | 

### Return type

[**VpnTunnel**](VpnTunnel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

