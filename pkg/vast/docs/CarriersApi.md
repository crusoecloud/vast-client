# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CarriersControlLed**](CarriersApi.md#CarriersControlLed) | **Patch** /carriers/{id}/control_led/ | Control Slot LED
[**CarriersHighlight**](CarriersApi.md#CarriersHighlight) | **Patch** /carriers/{id}/highlight/ | This endpoint highlights the control Slot
[**CarriersList**](CarriersApi.md#CarriersList) | **Get** /carriers/ | List Carriers
[**CarriersPartialUpdate**](CarriersApi.md#CarriersPartialUpdate) | **Patch** /carriers/{id}/ | Activate or Deactivate a Carrier
[**CarriersRead**](CarriersApi.md#CarriersRead) | **Get** /carriers/{id}/ | Return Details of One Carrier
[**ResetPci**](CarriersApi.md#ResetPci) | **Patch** /carriers/{id}/reset_pci/ | This endpoint power cycles the given Slot

# **CarriersControlLed**
> CarriersControlLed(ctx, id, optional)
Control Slot LED

This endpoint controls a slot LED

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Slot ID | 
 **optional** | ***CarriersApiCarriersControlLedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CarriersApiCarriersControlLedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdControlLedBody5**](IdControlLedBody5.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CarriersHighlight**
> AsyncTaskInResponse CarriersHighlight(ctx, id)
This endpoint highlights the control Slot

This endpoint highlights the control Slot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| slot ID | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CarriersList**
> []Carrier CarriersList(ctx, optional)
List Carriers

This endpoint lists carriers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CarriersApiCarriersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CarriersApiCarriersListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **optional.String**| Filter by carrier state | 
 **dboxName** | **optional.String**| Filter by parent DBox name | 
 **dboxId** | **optional.String**| Filter by parent DBox ID | 

### Return type

[**[]Carrier**](Carrier.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CarriersPartialUpdate**
> AsyncTaskInResponse CarriersPartialUpdate(ctx, id, optional)
Activate or Deactivate a Carrier

This endpoint activates or deactivates a carrier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Carrier ID | 
 **optional** | ***CarriersApiCarriersPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CarriersApiCarriersPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CarriersIdBody**](CarriersIdBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CarriersRead**
> Carrier CarriersRead(ctx, id)
Return Details of One Carrier

This endpoint returns details of a specific carrier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Carrier**](Carrier.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetPci**
> AsyncTaskInResponse ResetPci(ctx, id)
This endpoint power cycles the given Slot

This endpoint power cycles the given Slot

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| slot ID | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

