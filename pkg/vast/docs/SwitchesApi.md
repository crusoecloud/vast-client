# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SwitchesCreate**](SwitchesApi.md#SwitchesCreate) | **Post** /switches/ | Add Switch
[**SwitchesDelete**](SwitchesApi.md#SwitchesDelete) | **Delete** /switches/{id}/ | Delete Switch
[**SwitchesList**](SwitchesApi.md#SwitchesList) | **Get** /switches/ | List Switches
[**SwitchesPartialUpdate**](SwitchesApi.md#SwitchesPartialUpdate) | **Patch** /switches/{id}/ | Change Switch User Credentials
[**SwitchesRead**](SwitchesApi.md#SwitchesRead) | **Get** /switches/{id}/ | Return Details of a Switch.

# **SwitchesCreate**
> ModelSwitch SwitchesCreate(ctx, optional)
Add Switch

This endpoint adds a switch to the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwitchesApiSwitchesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwitchesApiSwitchesCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SwitchesBody**](SwitchesBody.md)|  | 

### Return type

[**ModelSwitch**](Switch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchesDelete**
> SwitchesDelete(ctx, id)
Delete Switch

This endpoint deletes a switch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Switch ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchesList**
> []ModelSwitch SwitchesList(ctx, optional)
List Switches

This endpoint lists switches.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwitchesApiSwitchesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwitchesApiSwitchesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **optional.String**| Filter by switch state | 

### Return type

[**[]ModelSwitch**](Switch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchesPartialUpdate**
> ModelSwitch SwitchesPartialUpdate(ctx, id, optional)
Change Switch User Credentials

This endpoint modifies a switch username and password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Switch ID | 
 **optional** | ***SwitchesApiSwitchesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwitchesApiSwitchesPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SwitchesIdBody**](SwitchesIdBody.md)|  | 

### Return type

[**ModelSwitch**](Switch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwitchesRead**
> ModelSwitch SwitchesRead(ctx, id)
Return Details of a Switch.

This endpoint returns details of a specified switch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Switch ID | 

### Return type

[**ModelSwitch**](Switch.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

