# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ColumnsDelete**](ColumnsApi.md#ColumnsDelete) | **Delete** /columns/delete/ | Delete a Database Table Column
[**ColumnsList**](ColumnsApi.md#ColumnsList) | **Get** /columns/ | List Database Table Columns
[**ColumnsRename**](ColumnsApi.md#ColumnsRename) | **Patch** /columns/rename/ | Rename a Database Table Column
[**ColumnsShow**](ColumnsApi.md#ColumnsShow) | **Get** /columns/show/ | Show a Database Table Column

# **ColumnsDelete**
> ColumnsDelete(ctx, optional)
Delete a Database Table Column

This endpoint deletes a Database Table Column.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ColumnsApiColumnsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ColumnsApiColumnsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ColumnDeleteParams**](ColumnDeleteParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ColumnsList**
> []Column ColumnsList(ctx, databaseName, schemaName, tableName, optional)
List Database Table Columns

This endpoint lists Database Table Columns.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **databaseName** | **string**| Getting list of objects by database_name | 
  **schemaName** | **string**| Getting list of objects by schema_name | 
  **tableName** | **string**| Getting list of objects by table_name | 
 **optional** | ***ColumnsApiColumnsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ColumnsApiColumnsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | [default to 999]
 **name** | **optional.String**| Getting list of objects by exact match | 
 **nameStartswith** | **optional.String**| Getting list of objects by prefix match | 
 **countOnly** | **optional.Bool**| Whether to only return count of objects | [default to false]

### Return type

[**[]Column**](Column.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ColumnsRename**
> ColumnsRename(ctx, optional)
Rename a Database Table Column

This endpoint renames column.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ColumnsApiColumnsRenameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ColumnsApiColumnsRenameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ColumnRenameParams**](ColumnRenameParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ColumnsShow**
> Column ColumnsShow(ctx, databaseName, schemaName, tableName, name)
Show a Database Table Column

This endpoint shows a Database Table Column.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **databaseName** | **string**| Getting list of objects by database_name | 
  **schemaName** | **string**| Getting list of objects by schema_name | 
  **tableName** | **string**| Getting list of objects by table_name | 
  **name** | **string**| Getting object by exact match | 

### Return type

[**Column**](Column.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

