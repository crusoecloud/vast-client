# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NisCreate**](NisApi.md#NisCreate) | **Post** /nis/ | Create a NIS Configuration
[**NisDelete**](NisApi.md#NisDelete) | **Delete** /nis/{id}/ | Delete a NIS Configuration
[**NisList**](NisApi.md#NisList) | **Get** /nis/ | List NIS Configurations
[**NisPartialUpdate**](NisApi.md#NisPartialUpdate) | **Patch** /nis/{id}/ | Modify NIS
[**NisRead**](NisApi.md#NisRead) | **Get** /nis/{id}/ | Return NIS Details.
[**NisSetPosixPrimary**](NisApi.md#NisSetPosixPrimary) | **Patch** /nis/{id}/set_posix_primary/ | Set NIS as Primary POSIX Provider
[**Refresh**](NisApi.md#Refresh) | **Patch** /nis/refresh/ | Refresh NIS Cache

# **NisCreate**
> Nis NisCreate(ctx, optional)
Create a NIS Configuration

This endpoint creates a NIS configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NisApiNisCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NisApiNisCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NisBody**](NisBody.md)|  | 

### Return type

[**Nis**](NIS.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NisDelete**
> NisDelete(ctx, id)
Delete a NIS Configuration

This endpoint deletes a NIS configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NisList**
> []Nis NisList(ctx, )
List NIS Configurations

This endpoint lists NIS configurations.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Nis**](NIS.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NisPartialUpdate**
> Nis NisPartialUpdate(ctx, id, optional)
Modify NIS

This endpoint modifies a NIS configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| NIS ID | 
 **optional** | ***NisApiNisPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NisApiNisPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NisIdBody**](NisIdBody.md)|  | 

### Return type

[**Nis**](NIS.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NisRead**
> Nis NisRead(ctx, id)
Return NIS Details.

This endpoint returns details of a NIS configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| NIS ID | 

### Return type

[**Nis**](NIS.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NisSetPosixPrimary**
> NisSetPosixPrimary(ctx, id)
Set NIS as Primary POSIX Provider

This endpoint sets NIS as the primary POSIX provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| NIS ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Refresh**
> Nis Refresh(ctx, )
Refresh NIS Cache

This endpoint refreshes the NIS cache by looking up user and netgroup maps on the NIS server.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Nis**](NIS.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

