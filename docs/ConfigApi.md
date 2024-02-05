# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigDelete**](ConfigApi.md#ConfigDelete) | **Delete** /config/{key}/ | Restore dynamic config to its default by key
[**ConfigList**](ConfigApi.md#ConfigList) | **Get** /config/ | List dynamic configs
[**ConfigRead**](ConfigApi.md#ConfigRead) | **Get** /config/{key}/ | Show dynamic configs
[**ConfigReset**](ConfigApi.md#ConfigReset) | **Post** /config/reset/ | Reset all dynamic configs to their default values
[**ConfigUpdate**](ConfigApi.md#ConfigUpdate) | **Patch** /config/{key}/ | Modify dynamic configs

# **ConfigDelete**
> ConfigDelete(ctx, key)
Restore dynamic config to its default by key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Dynamic Config key to filter by | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigList**
> []DynamicConfig ConfigList(ctx, )
List dynamic configs

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]DynamicConfig**](DynamicConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigRead**
> DynamicConfig ConfigRead(ctx, key)
Show dynamic configs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Dynamic Config key to filter by | 

### Return type

[**DynamicConfig**](DynamicConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigReset**
> ConfigReset(ctx, )
Reset all dynamic configs to their default values

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigUpdate**
> ConfigUpdate(ctx, key, optional)
Modify dynamic configs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| Dynamic Config key to filter by | 
 **optional** | ***ConfigApiConfigUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiConfigUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ModifyDynamicKeyParams**](ModifyDynamicKeyParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

