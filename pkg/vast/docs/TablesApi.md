# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TablesAddColumns**](TablesApi.md#TablesAddColumns) | **Patch** /tables/add_columns/ | Add Columns to a Database Table
[**TablesCreate**](TablesApi.md#TablesCreate) | **Post** /tables/ | Create a Database Table
[**TablesDelete**](TablesApi.md#TablesDelete) | **Delete** /tables/delete/ | Delete a Database Table
[**TablesList**](TablesApi.md#TablesList) | **Get** /tables/ | List Database Tables
[**TablesRename**](TablesApi.md#TablesRename) | **Patch** /tables/rename/ | Rename a Database Table
[**TablesShow**](TablesApi.md#TablesShow) | **Get** /tables/show/ | Show a Database Table

# **TablesAddColumns**
> TablesAddColumns(ctx, optional)
Add Columns to a Database Table

This endpoint adds Columns to a Database Table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TablesApiTablesAddColumnsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiTablesAddColumnsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TableAddColumnsParams**](TableAddColumnsParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TablesCreate**
> TablesCreate(ctx, optional)
Create a Database Table

This endpoint creates a Database Table

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TablesApiTablesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiTablesCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TableCreateParams**](TableCreateParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TablesDelete**
> TablesDelete(ctx, optional)
Delete a Database Table

This endpoint deletes a Database Table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TablesApiTablesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiTablesDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TableDeleteParams**](TableDeleteParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TablesList**
> []Table TablesList(ctx, databaseName, schemaName, optional)
List Database Tables

This endpoint lists Database Tables.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **databaseName** | **string**| Getting list of objects by database_name | 
  **schemaName** | **string**| Getting list of objects by schema_name | 
 **optional** | ***TablesApiTablesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiTablesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | [default to 999]
 **name** | **optional.String**| Getting list of objects by exact match | 
 **nameStartswith** | **optional.String**| Getting list of objects by prefix match | 
 **countOnly** | **optional.Bool**| Whether to only return count of objects | [default to false]
 **tableNameStartswith** | **optional.String**| Getting list of objects starting with by table_name__startswith (deprecated since 5.0) | 

### Return type

[**[]Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TablesRename**
> TablesRename(ctx, optional)
Rename a Database Table

This endpoint renames a Database Table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TablesApiTablesRenameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TablesApiTablesRenameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TableRenameParams**](TableRenameParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TablesShow**
> Table TablesShow(ctx, databaseName, schemaName, name)
Show a Database Table

This endpoint shows a Database Table.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **databaseName** | **string**| Getting list of objects by database_name | 
  **schemaName** | **string**| Getting list of objects by schema_name | 
  **name** | **string**| Getting object by exact match | 

### Return type

[**Table**](Table.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

