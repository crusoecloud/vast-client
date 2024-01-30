# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CboxesControlLed**](CboxesApi.md#CboxesControlLed) | **Patch** /cboxes/{id}/control_led/ | Control CBox LEDs
[**CboxesList**](CboxesApi.md#CboxesList) | **Get** /cboxes/ | List CBoxes
[**CboxesPartialUpdate**](CboxesApi.md#CboxesPartialUpdate) | **Patch** /cboxes/{id}/ | Modify CBox
[**CboxesRead**](CboxesApi.md#CboxesRead) | **Get** /cboxes/{id}/ | Return Details of a CBox
[**RefreshUid**](CboxesApi.md#RefreshUid) | **Patch** /cboxes/{id}/refresh_uid | Refreshes cbox uid to match it&#x27;s cnodes chassis serial

# **CboxesControlLed**
> AsyncTaskInResponse CboxesControlLed(ctx, id, optional)
Control CBox LEDs

This endpoint controls CBox LEDs (on/off)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| CBox ID | 
 **optional** | ***CboxesApiCboxesControlLedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CboxesApiCboxesControlLedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdControlLedBody3**](IdControlLedBody3.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CboxesList**
> []CBox CboxesList(ctx, optional)
List CBoxes

This endpoint lists the CBoxes that belong to the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CboxesApiCboxesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CboxesApiCboxesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 

### Return type

[**[]CBox**](CBox.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CboxesPartialUpdate**
> CboxesPartialUpdate(ctx, id, optional)
Modify CBox

This endpoint modifies a CBox description.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| CBox ID | 
 **optional** | ***CboxesApiCboxesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CboxesApiCboxesPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CboxesIdBody**](CboxesIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CboxesRead**
> CBox CboxesRead(ctx, id)
Return Details of a CBox

This endpoint returns details of a CBox.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| CBox ID | 

### Return type

[**CBox**](CBox.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshUid**
> RefreshUid(ctx, id)
Refreshes cbox uid to match it's cnodes chassis serial

This endpoint refreshes cbox uid to match it's cnodes chassis serial.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| CBox ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

