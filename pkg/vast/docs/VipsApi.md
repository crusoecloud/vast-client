# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VipsList**](VipsApi.md#VipsList) | **Get** /vips/ | List VIPs
[**VipsRead**](VipsApi.md#VipsRead) | **Get** /vips/{id}/ | Return Details of a VIP

# **VipsList**
> []Vip VipsList(ctx, optional)
List VIPs

This endpoint lists VIPs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VipsApiVipsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VipsApiVipsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **ip** | **optional.String**| Filter by IP | 
 **ipIcontains** | **optional.String**| Filter IP by case-insensitive containment | 
 **vippoolId** | **optional.String**| Filter by VIP Pool ID | 
 **vippoolName** | **optional.String**| Filter by VIP Pool name | 
 **cnodeName** | **optional.String**| Filter by CNode name | 
 **cnodeId** | **optional.String**| Filter by CNode ID | 
 **clusterName** | **optional.String**|  | 
 **clusterId** | **optional.String**|  | 

### Return type

[**[]Vip**](VIP.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VipsRead**
> Vip VipsRead(ctx, id)
Return Details of a VIP

This endpoint returns information about a specified VIP.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| VIP ID | 

### Return type

[**Vip**](VIP.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

