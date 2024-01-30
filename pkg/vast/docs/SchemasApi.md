# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SchemasCreate**](SchemasApi.md#SchemasCreate) | **Post** /schemas/ | Create a Database Schema
[**SchemasDelete**](SchemasApi.md#SchemasDelete) | **Delete** /schemas/delete/ | Delete a Database Schema
[**SchemasList**](SchemasApi.md#SchemasList) | **Get** /schemas/ | List Database Schemas
[**SchemasRename**](SchemasApi.md#SchemasRename) | **Patch** /schemas/rename/ | Rename a Database Schema
[**SchemasShow**](SchemasApi.md#SchemasShow) | **Get** /schemas/show/ | Show a Database Schema

# **SchemasCreate**
> SchemasCreate(ctx, optional)
Create a Database Schema

This endpoint creates a Database Schema

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SchemasApiSchemasCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SchemasApiSchemasCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SchemaCreateParams**](SchemaCreateParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SchemasDelete**
> SchemasDelete(ctx, optional)
Delete a Database Schema

This endpoint deletes a Database Schema.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SchemasApiSchemasDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SchemasApiSchemasDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SchemaDeleteParams**](SchemaDeleteParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SchemasList**
> []Schema SchemasList(ctx, optional)
List Database Schemas

This endpoint lists Database Schemas.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SchemasApiSchemasListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SchemasApiSchemasListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | [default to 999]
 **databaseName** | **optional.String**| Getting list of objects by database_name | 
 **schemaName** | **optional.String**| Getting list of objects by schema_name | 
 **name** | **optional.String**| Getting list of objects by exact match | 
 **nameStartswith** | **optional.String**| Getting list of objects by prefix match | 
 **countOnly** | **optional.Bool**| Whether to only return count of objects | [default to false]
 **schemaNameStartswith** | **optional.String**| Getting list of objects starting with by schema_name__startswith (deprecated since 5.0) | 
 **byLevel** | **optional.String**| Can be true or false. If this by_level is true, the VMS will provide a list of all schemas of only one level. to get the next level, need to make another request with the name of the parent scheme  | 

### Return type

[**[]Schema**](Schema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SchemasRename**
> SchemasRename(ctx, optional)
Rename a Database Schema

This endpoint renames a Database Schema.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SchemasApiSchemasRenameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SchemasApiSchemasRenameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SchemaRenameParams**](SchemaRenameParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SchemasShow**
> Schema SchemasShow(ctx, databaseName, name)
Show a Database Schema

This endpoint shows a Database Schema.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **databaseName** | **string**| Getting list of objects by database_name | 
  **name** | **string**| Getting object by exact match | 

### Return type

[**Schema**](Schema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

