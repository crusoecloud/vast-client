# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CallhomeconfigsList**](CallhomeconfigsApi.md#CallhomeconfigsList) | **Get** /callhomeconfigs/ | List Call Home Configurations
[**CallhomeconfigsPartialUpdate**](CallhomeconfigsApi.md#CallhomeconfigsPartialUpdate) | **Patch** /callhomeconfigs/{id}/ | Modify the Call Home Configuration
[**CallhomeconfigsRead**](CallhomeconfigsApi.md#CallhomeconfigsRead) | **Get** /callhomeconfigs/{id}/ | Return Call Home Configuration Details
[**CallhomeconfigsRegisterCluster**](CallhomeconfigsApi.md#CallhomeconfigsRegisterCluster) | **Patch** /callhomeconfigs/{id}/register-cluster/ | Register Cluster with VAST Cloud Services
[**CallhomeconfigsSend**](CallhomeconfigsApi.md#CallhomeconfigsSend) | **Patch** /callhomeconfigs/{id}/send/ | Send Call Home Configuration

# **CallhomeconfigsList**
> []CallhomeConfig CallhomeconfigsList(ctx, optional)
List Call Home Configurations

This endpoint lists Call Home configurations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CallhomeconfigsApiCallhomeconfigsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CallhomeconfigsApiCallhomeconfigsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 

### Return type

[**[]CallhomeConfig**](CallhomeConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallhomeconfigsPartialUpdate**
> CallhomeconfigsPartialUpdate(ctx, id, optional)
Modify the Call Home Configuration

This endpoint modifies the Call Home configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Call Home Configuration ID | 
 **optional** | ***CallhomeconfigsApiCallhomeconfigsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CallhomeconfigsApiCallhomeconfigsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CallhomeconfigsIdBody**](CallhomeconfigsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallhomeconfigsRead**
> CallhomeConfig CallhomeconfigsRead(ctx, id)
Return Call Home Configuration Details

This endpoint returns Call Home Configuration details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Callhomeconfig ID | 

### Return type

[**CallhomeConfig**](CallhomeConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallhomeconfigsRegisterCluster**
> AsyncTaskInResponse CallhomeconfigsRegisterCluster(ctx, id, optional)
Register Cluster with VAST Cloud Services

This endpoint registers a cluster with VAST Cloud Services.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***CallhomeconfigsApiCallhomeconfigsRegisterClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CallhomeconfigsApiCallhomeconfigsRegisterClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdRegisterclusterBody**](IdRegisterclusterBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallhomeconfigsSend**
> AsyncTaskInResponse CallhomeconfigsSend(ctx, id)
Send Call Home Configuration

This endpoint sends a callhomeconfig.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

