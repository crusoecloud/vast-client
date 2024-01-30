# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalPathsCreate**](GlobalpathsApi.md#GlobalPathsCreate) | **Post** /globalpaths/ | Create Global Path
[**GlobalPathsList**](GlobalpathsApi.md#GlobalPathsList) | **Get** /globalpaths/ | Return Global Paths
[**GlobalPathsPartialUpdate**](GlobalpathsApi.md#GlobalPathsPartialUpdate) | **Patch** /globalpaths/{id}/ | Modify a Global Path
[**GlobalPathsRead**](GlobalpathsApi.md#GlobalPathsRead) | **Get** /globalpaths/{id}/ | Return Details of a Global Path
[**GlobalpathsDelete**](GlobalpathsApi.md#GlobalpathsDelete) | **Delete** /globalpaths/{id}/ | Delete Global Path

# **GlobalPathsCreate**
> GlobalPath GlobalPathsCreate(ctx, optional)
Create Global Path

This endpoint creates a global path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GlobalpathsApiGlobalPathsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalpathsApiGlobalPathsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GlobalpathsBody**](GlobalpathsBody.md)|  | 

### Return type

[**GlobalPath**](GlobalPath.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalPathsList**
> []GlobalPath GlobalPathsList(ctx, optional)
Return Global Paths

This endpoint lists global paths.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GlobalpathsApiGlobalPathsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalpathsApiGlobalPathsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **path** | **optional.String**| Filter by global path | 

### Return type

[**[]GlobalPath**](GlobalPath.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalPathsPartialUpdate**
> GlobalPathsPartialUpdate(ctx, id, optional)
Modify a Global Path

This endpoint modifies a global path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Global Path ID | 
 **optional** | ***GlobalpathsApiGlobalPathsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalpathsApiGlobalPathsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GlobalpathsIdBody**](GlobalpathsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalPathsRead**
> GlobalPath GlobalPathsRead(ctx, id)
Return Details of a Global Path

This endpoint returns details of a specified global path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Global Path ID | 

### Return type

[**GlobalPath**](GlobalPath.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalpathsDelete**
> GlobalpathsDelete(ctx, id)
Delete Global Path

This endpoint deletes a global path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Global Path ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

