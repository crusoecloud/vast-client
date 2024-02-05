# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ControlLed**](DboxesApi.md#ControlLed) | **Patch** /dboxes/{id}/control_led/ | Control DBox LEDs
[**DboxesAdd**](DboxesApi.md#DboxesAdd) | **Post** /dboxes/add/ | Add DBox
[**DboxesDelete**](DboxesApi.md#DboxesDelete) | **Delete** /dboxes/{id}/ | Delete a DBox
[**DboxesList**](DboxesApi.md#DboxesList) | **Get** /dboxes/ | List DBoxes
[**DboxesPartialUpdate**](DboxesApi.md#DboxesPartialUpdate) | **Patch** /dboxes/{id}/ | Modify a DBox Description
[**DboxesRead**](DboxesApi.md#DboxesRead) | **Get** /dboxes/{id}/ | Return Details of a DBox.

# **ControlLed**
> AsyncTaskInResponse ControlLed(ctx, id, optional)
Control DBox LEDs

This endpoint controls DBox LEDs (on/off)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DBox ID | 
 **optional** | ***DboxesApiControlLedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DboxesApiControlLedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdControlLedBody2**](IdControlLedBody2.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DboxesAdd**
> AsyncTaskInResponse DboxesAdd(ctx, optional)
Add DBox

This endpoint adds a DBox to the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DboxesApiDboxesAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DboxesApiDboxesAddOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DboxesAddBody**](DboxesAddBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DboxesDelete**
> AsyncTaskInResponse DboxesDelete(ctx, id)
Delete a DBox

This endpoint deletes a DBox.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DBox ID | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DboxesList**
> []DBox DboxesList(ctx, optional)
List DBoxes

This endpoint lists the DBoxes that belong to the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DboxesApiDboxesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DboxesApiDboxesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 

### Return type

[**[]DBox**](DBox.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DboxesPartialUpdate**
> DboxesPartialUpdate(ctx, id, optional)
Modify a DBox Description

This endpoint modifies a DBox description.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DBox ID | 
 **optional** | ***DboxesApiDboxesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DboxesApiDboxesPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DboxesIdBody**](DboxesIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DboxesRead**
> DBox DboxesRead(ctx, id)
Return Details of a DBox.

This endpoint returns details of a DBox.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DBox ID | 

### Return type

[**DBox**](DBox.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

