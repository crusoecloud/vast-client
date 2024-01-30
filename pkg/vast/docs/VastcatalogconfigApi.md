# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VastcatalogconfigColumns**](VastcatalogconfigApi.md#VastcatalogconfigColumns) | **Get** /bigcatalogconfig/columns/ | Receives Vast Catalog columns
[**VastcatalogconfigDelete**](VastcatalogconfigApi.md#VastcatalogconfigDelete) | **Delete** /bigcatalogconfig/{id}/ | Delete Vast Catalog Config and Vast Catalog Policy.
[**VastcatalogconfigList**](VastcatalogconfigApi.md#VastcatalogconfigList) | **Get** /bigcatalogconfig/ | List of Vast Catalog Config
[**VastcatalogconfigModify**](VastcatalogconfigApi.md#VastcatalogconfigModify) | **Patch** /bigcatalogconfig/{id}/ | Modify Vast Catalog Config
[**VastcatalogconfigQueryData**](VastcatalogconfigApi.md#VastcatalogconfigQueryData) | **Post** /bigcatalogconfig/query_data/ | Receives information according to the sent rules
[**VastcatalogconfigRead**](VastcatalogconfigApi.md#VastcatalogconfigRead) | **Get** /bigcatalogconfig/{id}/ | Return Vast Catalog Config
[**VastcatalogconfigStats**](VastcatalogconfigApi.md#VastcatalogconfigStats) | **Get** /bigcatalogconfig/stats/ | Receives information about Vast Catalog stats

# **VastcatalogconfigColumns**
> []Column VastcatalogconfigColumns(ctx, )
Receives Vast Catalog columns

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Column**](Column.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VastcatalogconfigDelete**
> VastcatalogconfigDelete(ctx, id)
Delete Vast Catalog Config and Vast Catalog Policy.

This endpoint deletes Vast Catalog Config.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Vastcatalogconfig ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VastcatalogconfigList**
> []VastCatalogConfig VastcatalogconfigList(ctx, optional)
List of Vast Catalog Config

This endpoint returns list of Vast Catalog Config.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VastcatalogconfigApiVastcatalogconfigListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VastcatalogconfigApiVastcatalogconfigListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**[]VastCatalogConfig**](VastCatalogConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VastcatalogconfigModify**
> VastcatalogconfigModify(ctx, id, optional)
Modify Vast Catalog Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Vast Catalog Config Policy ID | 
 **optional** | ***VastcatalogconfigApiVastcatalogconfigModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VastcatalogconfigApiVastcatalogconfigModifyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of BigcatalogconfigIdBody**](BigcatalogconfigIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VastcatalogconfigQueryData**
> QueryDataOutputSchema VastcatalogconfigQueryData(ctx, optional)
Receives information according to the sent rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VastcatalogconfigApiVastcatalogconfigQueryDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VastcatalogconfigApiVastcatalogconfigQueryDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VastCatalogQueryDataInputSchema**](VastCatalogQueryDataInputSchema.md)|  | 
 **page** | **optional.**|  | 
 **limit** | **optional.**|  | [default to 1000]

### Return type

[**QueryDataOutputSchema**](QueryDataOutputSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VastcatalogconfigRead**
> VastCatalogConfig VastcatalogconfigRead(ctx, id)
Return Vast Catalog Config

This endpoint returns Vast Catalog Config details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Vastcatalogconfig ID | 

### Return type

[**VastCatalogConfig**](VastCatalogConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VastcatalogconfigStats**
> Stats VastcatalogconfigStats(ctx, )
Receives information about Vast Catalog stats

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Stats**](Stats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

