# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VastcatalogindexedcolumnsAdd**](VastcatalogindexedcolumnsApi.md#VastcatalogindexedcolumnsAdd) | **Patch** /bigcatalogindexedcolumns/add/ | Add new indexed column for vast catalog
[**VastcatalogindexedcolumnsList**](VastcatalogindexedcolumnsApi.md#VastcatalogindexedcolumnsList) | **Get** /bigcatalogindexedcolumns/ | List Vast Catalog Tags
[**VastcatalogindexedcolumnsRemove**](VastcatalogindexedcolumnsApi.md#VastcatalogindexedcolumnsRemove) | **Delete** /bigcatalogindexedcolumns/remove/ | Remove Vast Catalog indexed column

# **VastcatalogindexedcolumnsAdd**
> VastcatalogindexedcolumnsAdd(ctx, optional)
Add new indexed column for vast catalog

This endpoint adds the Vast Catalog Indexed Column.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VastcatalogindexedcolumnsApiVastcatalogindexedcolumnsAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VastcatalogindexedcolumnsApiVastcatalogindexedcolumnsAddOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BigcatalogindexedcolumnsAddBody**](BigcatalogindexedcolumnsAddBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VastcatalogindexedcolumnsList**
> []VastCatalogIndexedColumn VastcatalogindexedcolumnsList(ctx, optional)
List Vast Catalog Tags

This endpoint lists the Vast Catalog Tags.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VastcatalogindexedcolumnsApiVastcatalogindexedcolumnsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VastcatalogindexedcolumnsApiVastcatalogindexedcolumnsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **columnType** | **optional.String**|  | 

### Return type

[**[]VastCatalogIndexedColumn**](VastCatalogIndexedColumn.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VastcatalogindexedcolumnsRemove**
> VastcatalogindexedcolumnsRemove(ctx, optional)
Remove Vast Catalog indexed column

This endpoint removes the Vast Catalog Indexed Column.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VastcatalogindexedcolumnsApiVastcatalogindexedcolumnsRemoveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VastcatalogindexedcolumnsApiVastcatalogindexedcolumnsRemoveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BigcatalogindexedcolumnsRemoveBody**](BigcatalogindexedcolumnsRemoveBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

