# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterVsettings**](VsettingsApi.md#ClusterVsettings) | **Get** /clusters/{id}/vsettings/ | Show or list cluster vsettings
[**ClusterVsettingsDelete**](VsettingsApi.md#ClusterVsettingsDelete) | **Delete** /clusters/{id}/vsettings/ | Delete Cluster VSetting by key
[**ClusterVsettingsPartialUpdate**](VsettingsApi.md#ClusterVsettingsPartialUpdate) | **Patch** /clusters/{id}/vsettings/ | Modify Cluster VSettings

# **ClusterVsettings**
> map[string]string ClusterVsettings(ctx, id, optional)
Show or list cluster vsettings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***VsettingsApiClusterVsettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VsettingsApiClusterVsettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **optional.String**| VSetting key | 

### Return type

[**map[string]string**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterVsettingsDelete**
> ClusterVsettingsDelete(ctx, id, optional)
Delete Cluster VSetting by key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***VsettingsApiClusterVsettingsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VsettingsApiClusterVsettingsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DeleteVsettingParams**](DeleteVsettingParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterVsettingsPartialUpdate**
> ClusterVsettingsPartialUpdate(ctx, id, optional)
Modify Cluster VSettings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***VsettingsApiClusterVsettingsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VsettingsApiClusterVsettingsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ModifyVsettingsParams**](ModifyVsettingsParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

