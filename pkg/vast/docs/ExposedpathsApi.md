# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExposedPathsCreate**](ExposedpathsApi.md#ExposedPathsCreate) | **Post** /exposedpaths/ | Create Exposed Path
[**ExposedPathsDelete**](ExposedpathsApi.md#ExposedPathsDelete) | **Delete** /exposedpaths/{id}/ | Delete Exposed Path
[**ExposedPathsList**](ExposedpathsApi.md#ExposedPathsList) | **Get** /exposedpaths/ | Return Exposed Paths
[**ExposedPathsPartialUpdate**](ExposedpathsApi.md#ExposedPathsPartialUpdate) | **Patch** /exposedpaths/{id}/ | Modify a Exposed Path
[**ExposedPathsRead**](ExposedpathsApi.md#ExposedPathsRead) | **Get** /exposedpaths/{id}/ | Return Details of a Exposed Path

# **ExposedPathsCreate**
> ExposedPath ExposedPathsCreate(ctx, optional)
Create Exposed Path

This endpoint creates a exposed path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExposedpathsApiExposedPathsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExposedpathsApiExposedPathsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ExposedpathsBody**](ExposedpathsBody.md)|  | 

### Return type

[**ExposedPath**](ExposedPath.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExposedPathsDelete**
> ExposedPathsDelete(ctx, id)
Delete Exposed Path

This endpoint deletes a exposed path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Exposed Path ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExposedPathsList**
> []ExposedPath ExposedPathsList(ctx, optional)
Return Exposed Paths

This endpoint lists exposed paths.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExposedpathsApiExposedPathsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExposedpathsApiExposedPathsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **path** | **optional.String**| Filter by exposed path | 

### Return type

[**[]ExposedPath**](ExposedPath.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExposedPathsPartialUpdate**
> ExposedPathsPartialUpdate(ctx, id, optional)
Modify a Exposed Path

This endpoint modifies a exposed path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Exposed Path ID | 
 **optional** | ***ExposedpathsApiExposedPathsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExposedpathsApiExposedPathsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ExposedpathsIdBody**](ExposedpathsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExposedPathsRead**
> ExposedPath ExposedPathsRead(ctx, id)
Return Details of a Exposed Path

This endpoint returns details of a specified exposed path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Exposed Path ID | 

### Return type

[**ExposedPath**](ExposedPath.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

