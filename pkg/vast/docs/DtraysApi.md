# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtraysControlLed**](DtraysApi.md#DtraysControlLed) | **Patch** /dtrays/{id}/control_led/ | Control DTray (DNode) LEDs
[**DtraysList**](DtraysApi.md#DtraysList) | **Get** /dtrays/ | List DTrays
[**DtraysPartialUpdate**](DtraysApi.md#DtraysPartialUpdate) | **Patch** /dtrays/{id}/ | Activate/Deactivate DTray
[**DtraysRead**](DtraysApi.md#DtraysRead) | **Get** /dtrays/{id}/ | Return Details of a DTray
[**DtraysRename**](DtraysApi.md#DtraysRename) | **Patch** /dtrays/{id}/rename/ | Rename Dtray

# **DtraysControlLed**
> DtraysControlLed(ctx, id, optional)
Control DTray (DNode) LEDs

This endpoint controls the DTray (DNode) LEDs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DTray ID | 
 **optional** | ***DtraysApiDtraysControlLedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DtraysApiDtraysControlLedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdControlLedBody4**](IdControlLedBody4.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DtraysList**
> []DTray DtraysList(ctx, optional)
List DTrays

This endpoint lists the all DTrays.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DtraysApiDtraysListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DtraysApiDtraysListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **state** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **enabled** | **optional.Bool**|  | 

### Return type

[**[]DTray**](DTray.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DtraysPartialUpdate**
> AsyncTaskInResponse DtraysPartialUpdate(ctx, id, optional)
Activate/Deactivate DTray

This endpoint activates, deactivates a DTray.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DTray ID | 
 **optional** | ***DtraysApiDtraysPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DtraysApiDtraysPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DtraysIdBody**](DtraysIdBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DtraysRead**
> DTray DtraysRead(ctx, id)
Return Details of a DTray

This endpoint returns details of a DTray.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DTray ID | 

### Return type

[**DTray**](DTray.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DtraysRename**
> DtraysRename(ctx, id, optional)
Rename Dtray

This endpoint renames a Dtray.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Dtray ID | 
 **optional** | ***DtraysApiDtraysRenameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DtraysApiDtraysRenameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdRenameBody2**](IdRenameBody2.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

