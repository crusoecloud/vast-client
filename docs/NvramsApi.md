# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NvramFormat**](NvramsApi.md#NvramFormat) | **Patch** /nvrams/{id}/format/ | Format NVRAM
[**NvramsList**](NvramsApi.md#NvramsList) | **Get** /nvrams/ | List NVRAMs
[**NvramsPartialUpdate**](NvramsApi.md#NvramsPartialUpdate) | **Patch** /nvrams/{id}/ | Activate or Deactivate an NVRAM
[**NvramsRead**](NvramsApi.md#NvramsRead) | **Get** /nvrams/{id}/ | Return Details of an NVRAM

# **NvramFormat**
> AsyncTaskInResponse NvramFormat(ctx, id)
Format NVRAM

This endpoint formats the NVRAM.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| NVRAM ID | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvramsList**
> []Nvram NvramsList(ctx, optional)
List NVRAMs

This endpoint lists NVRAMs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NvramsApiNvramsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvramsApiNvramsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **state** | **optional.String**| Filter by NVRAM state | 
 **sn** | **optional.String**| Filter by NVRAM serial number | 
 **model** | **optional.String**| Filter by NVRAM model | 
 **dboxName** | **optional.String**| Filter by parent DBox name | 
 **dboxId** | **optional.String**| Filter by parent DBox ID | 
 **fwVersion** | **optional.String**| Filter by NVRAM firmware version | 

### Return type

[**[]Nvram**](NVRAM.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvramsPartialUpdate**
> AsyncTaskInResponse NvramsPartialUpdate(ctx, id, optional)
Activate or Deactivate an NVRAM

This endpoint activates or deactivates an NVRAM.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| NVRAM ID | 
 **optional** | ***NvramsApiNvramsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NvramsApiNvramsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NvramsIdBody**](NvramsIdBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NvramsRead**
> Nvram NvramsRead(ctx, id)
Return Details of an NVRAM

This endpoint returns information about an NVRAM.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Nvram**](NVRAM.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

