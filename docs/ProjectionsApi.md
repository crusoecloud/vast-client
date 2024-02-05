# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectionsCreate**](ProjectionsApi.md#ProjectionsCreate) | **Post** /projections/ | Create a Database Table Semi-sorted Projection
[**ProjectionsDelete**](ProjectionsApi.md#ProjectionsDelete) | **Delete** /projections/delete/ | Delete a Database Table Semi-sorted Projection
[**ProjectionsList**](ProjectionsApi.md#ProjectionsList) | **Get** /projections/ | List Database Table Semi-sorted Projections
[**ProjectionsRename**](ProjectionsApi.md#ProjectionsRename) | **Patch** /projections/rename/ | Rename a Database Table Semi-sorted Projection
[**ProjectionsShow**](ProjectionsApi.md#ProjectionsShow) | **Get** /projections/show/ | Show a Database Table Semi-sorted Projection

# **ProjectionsCreate**
> ProjectionsCreate(ctx, optional)
Create a Database Table Semi-sorted Projection

This endpoint creates a Database Table Semi-sorted Projection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectionsApiProjectionsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectionsApiProjectionsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ProjectionCreateParams**](ProjectionCreateParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectionsDelete**
> ProjectionsDelete(ctx, optional)
Delete a Database Table Semi-sorted Projection

This endpoint deletes a Database Table Semi-sorted Projection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectionsApiProjectionsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectionsApiProjectionsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ProjectionDeleteParams**](ProjectionDeleteParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectionsList**
> []Projection ProjectionsList(ctx, databaseName, schemaName, tableName, optional)
List Database Table Semi-sorted Projections

This endpoint lists Database Table Semi-sorted Projections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **databaseName** | **string**| Getting list of objects by database_name | 
  **schemaName** | **string**| Getting list of objects by schema_name | 
  **tableName** | **string**| Getting list of objects by table_name | 
 **optional** | ***ProjectionsApiProjectionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectionsApiProjectionsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | [default to 999]
 **name** | **optional.String**| Getting list of objects by exact match | 
 **nameStartswith** | **optional.String**| Getting list of objects by prefix match | 
 **countOnly** | **optional.Bool**| Whether to only return count of objects | [default to false]

### Return type

[**[]Projection**](Projection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectionsRename**
> ProjectionsRename(ctx, optional)
Rename a Database Table Semi-sorted Projection

This endpoint renames a Database Table Semi-sorted Projection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectionsApiProjectionsRenameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectionsApiProjectionsRenameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ProjectionRenameParams**](ProjectionRenameParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectionsShow**
> DetailedProjection ProjectionsShow(ctx, databaseName, schemaName, tableName, name)
Show a Database Table Semi-sorted Projection

This endpoint shows a Database Table Semi-sorted Projection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **databaseName** | **string**| Getting list of objects by database_name | 
  **schemaName** | **string**| Getting list of objects by schema_name | 
  **tableName** | **string**| Getting list of objects by table_name | 
  **name** | **string**| Getting object by exact match | 

### Return type

[**DetailedProjection**](DetailedProjection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

