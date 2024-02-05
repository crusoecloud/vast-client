# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HostsDiscover**](HostsApi.md#HostsDiscover) | **Get** /hosts/discover | Discover Hosts
[**HostsList**](HostsApi.md#HostsList) | **Get** /hosts/ | List Hosts
[**HostsRead**](HostsApi.md#HostsRead) | **Get** /hosts/{id}/ | Return Details of Discovered Host

# **HostsDiscover**
> AsyncTaskInResponse HostsDiscover(ctx, )
Discover Hosts

Discover hosts

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostsList**
> []Host HostsList(ctx, optional)
List Hosts

This endpoint lists discovered hosts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HostsApiHostsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HostsApiHostsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **auto** | **optional.Bool**| Specify auto or manually hosts | 
 **ipList** | **optional.String**| Comma-separated list of nodes IPs | 
 **loopback** | **optional.Bool**| Loopback nodes | 
 **ip** | **optional.String**|  | 
 **state** | **optional.String**|  | 
 **clusterName** | **optional.String**|  | 
 **clusterId** | **optional.Int32**|  | 
 **installState** | **optional.String**|  | 
 **nodeType** | **optional.String**|  | 
 **swVersion** | **optional.String**|  | 
 **build** | **optional.String**|  | 

### Return type

[**[]Host**](Host.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostsRead**
> Host HostsRead(ctx, id)
Return Details of Discovered Host

This endpoint returns details of a discovered host.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Host**](Host.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

