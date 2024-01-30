# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SsdFormat**](SsdsApi.md#SsdFormat) | **Patch** /ssds/{id}/format/ | Format SSD
[**SsdsList**](SsdsApi.md#SsdsList) | **Get** /ssds/ | List SSDs
[**SsdsPartialUpdate**](SsdsApi.md#SsdsPartialUpdate) | **Patch** /ssds/{id}/ | Activate/Deactivate an SSD
[**SsdsRead**](SsdsApi.md#SsdsRead) | **Get** /ssds/{id}/ | Return Details of an SSD

# **SsdFormat**
> AsyncTaskInResponse SsdFormat(ctx, id)
Format SSD

This endpoint formats an SSD.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| SSD ID | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SsdsList**
> []Ssd SsdsList(ctx, optional)
List SSDs

This endpoint lists SSDs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SsdsApiSsdsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SsdsApiSsdsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.String**|  | 
 **page** | **optional.String**|  | 
 **state** | **optional.String**| Filter by SSD state | 
 **sn** | **optional.String**| Filter by SSD serial number | 
 **model** | **optional.String**| Filter by SSD model | 
 **dboxName** | **optional.String**| Filter by parent DBox name | 
 **dboxId** | **optional.String**| Filter by parent DBox ID | 
 **fwVersion** | **optional.String**| Filter by SSD firmware version | 

### Return type

[**[]Ssd**](SSD.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SsdsPartialUpdate**
> AsyncTaskInResponse SsdsPartialUpdate(ctx, id, optional)
Activate/Deactivate an SSD

This endpoint activates or deactivates an SSD.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| SSD ID | 
 **optional** | ***SsdsApiSsdsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SsdsApiSsdsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SsdsIdBody**](SsdsIdBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SsdsRead**
> Ssd SsdsRead(ctx, id)
Return Details of an SSD

This endpoint returns information about an SSD.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Ssd**](SSD.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

