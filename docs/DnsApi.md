# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsCreate**](DnsApi.md#DnsCreate) | **Post** /dns/ | Create a VAST-DNS Server Configuration
[**DnsDelete**](DnsApi.md#DnsDelete) | **Delete** /dns/{id}/ | Delete VAST-DNS Server Configuration.
[**DnsList**](DnsApi.md#DnsList) | **Get** /dns/ | List VAST-DNS Configuration
[**DnsPartialUpdate**](DnsApi.md#DnsPartialUpdate) | **Patch** /dns/{id}/ | Modify VAST-DNS Server Configuration
[**DnsRead**](DnsApi.md#DnsRead) | **Get** /dns/{id}/ | Return Details of VAST-DNS Configuration

# **DnsCreate**
> Dns DnsCreate(ctx, optional)
Create a VAST-DNS Server Configuration

This endpoint creates a configuration for the VAST DNS service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DnsApiDnsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DnsApiDnsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DnsBody**](DnsBody.md)|  | 

### Return type

[**Dns**](DNS.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DnsDelete**
> DnsDelete(ctx, id)
Delete VAST-DNS Server Configuration.

This endpoint deletes a VAST-DNS server configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DNS ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DnsList**
> []Dns DnsList(ctx, optional)
List VAST-DNS Configuration

This endpoint lists the VAST-DNS server configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DnsApiDnsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DnsApiDnsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 

### Return type

[**[]Dns**](DNS.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DnsPartialUpdate**
> DnsPartialUpdate(ctx, id, optional)
Modify VAST-DNS Server Configuration

This endpoint modifies the VAST-DNS Server configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DNS ID | 
 **optional** | ***DnsApiDnsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DnsApiDnsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DnsIdBody**](DnsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DnsRead**
> Dns DnsRead(ctx, id)
Return Details of VAST-DNS Configuration

This endpoint returns details of a VAST-DNS server configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DNS ID | 

### Return type

[**Dns**](DNS.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

