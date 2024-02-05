# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FansList**](FansApi.md#FansList) | **Get** /fans/ | List Fans
[**FansRead**](FansApi.md#FansRead) | **Get** /fans/{id}/ | Return Details of Fan

# **FansList**
> []Fan FansList(ctx, optional)
List Fans

This endpoint lists fans.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FansApiFansListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FansApiFansListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **state** | **optional.String**| Filter by fan state | 
 **sn** | **optional.String**| Filter by fan serial number | 
 **model** | **optional.String**| Filter by fan model | 
 **boxName** | **optional.String**| Filter by parent CBox or DBox name | 
 **boxId** | **optional.String**| Filter by parent CBox or DBox ID | 

### Return type

[**[]Fan**](Fan.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FansRead**
> Fan FansRead(ctx, id)
Return Details of Fan

This endpoint returns details of a specified fan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Fan ID | 

### Return type

[**Fan**](Fan.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

