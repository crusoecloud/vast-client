# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeltaConfigPartialUpdate**](DeltaApi.md#DeltaConfigPartialUpdate) | **Patch** /delta/config | Update Delta config
[**DeltaConfigRead**](DeltaApi.md#DeltaConfigRead) | **Get** /delta/config | Returns Delta Storage config.
[**DeltaStorageRead**](DeltaApi.md#DeltaStorageRead) | **Get** /delta/ | Returns Delta Protocol records.

# **DeltaConfigPartialUpdate**
> DeltaConfig DeltaConfigPartialUpdate(ctx, optional)
Update Delta config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeltaApiDeltaConfigPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeltaApiDeltaConfigPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DeltaConfigBody**](DeltaConfigBody.md)|  | 

### Return type

[**DeltaConfig**](DeltaConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeltaConfigRead**
> DeltaConfig DeltaConfigRead(ctx, )
Returns Delta Storage config.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DeltaConfig**](DeltaConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeltaStorageRead**
> DeltaStorage DeltaStorageRead(ctx, sequence, generation, optional)
Returns Delta Protocol records.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sequence** | **string**| Runtime sequence id for delta. Reset if HA happened. example: 8f021a67-cfd6-4896-ad67-3e32a376c6f4 | 
  **generation** | **int32**|  | 
 **optional** | ***DeltaApiDeltaStorageReadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeltaApiDeltaStorageReadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **opType** | **optional.String**| Operation type | 
 **objectType** | **optional.String**| example: cluster | 
 **objectGuid** | **optional.String**|  | 
 **objectId** | **optional.Int32**|  | 

### Return type

[**DeltaStorage**](DeltaStorage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

