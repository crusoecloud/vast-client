# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectioncolumnsList**](ProjectioncolumnsApi.md#ProjectioncolumnsList) | **Get** /projectioncolumns/ | List Database Table Semi-sorted Projection Columns
[**ProjectioncolumnsShow**](ProjectioncolumnsApi.md#ProjectioncolumnsShow) | **Get** /projectioncolumns/show/ | Show a Database Table Semi-sorted Projection Column

# **ProjectioncolumnsList**
> []ProjectionColumn ProjectioncolumnsList(ctx, databaseName, schemaName, tableName, projectionName, optional)
List Database Table Semi-sorted Projection Columns

This endpoint lists the Database Table Semi-sorted Projection Columns.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **databaseName** | **string**| Getting list of objects by database_name | 
  **schemaName** | **string**| Getting list of objects by schema_name | 
  **tableName** | **string**| Getting list of objects by table_name | 
  **projectionName** | **string**| Getting list of objects by projection_name | 
 **optional** | ***ProjectioncolumnsApiProjectioncolumnsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectioncolumnsApiProjectioncolumnsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | [default to 999]
 **name** | **optional.String**| Getting list of objects by exact match | 
 **nameStartswith** | **optional.String**| Getting list of objects by prefix match | 
 **countOnly** | **optional.Bool**| Whether to only return count of objects | [default to false]

### Return type

[**[]ProjectionColumn**](ProjectionColumn.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectioncolumnsShow**
> ProjectionColumn ProjectioncolumnsShow(ctx, databaseName, schemaName, tableName, projectionName, name)
Show a Database Table Semi-sorted Projection Column

This endpoint shows a Database Table Semi-sorted Projection Column.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **databaseName** | **string**| Getting list of objects by database_name | 
  **schemaName** | **string**| Getting list of objects by schema_name | 
  **tableName** | **string**| Getting list of objects by table_name | 
  **projectionName** | **string**| Getting list of objects by projection_name | 
  **name** | **string**| Getting object by exact match | 

### Return type

[**ProjectionColumn**](ProjectionColumn.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

