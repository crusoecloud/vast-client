# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SupportbundlesCreate**](SupportbundlesApi.md#SupportbundlesCreate) | **Post** /supportbundles/ | Create a Support Bundle
[**SupportbundlesDelete**](SupportbundlesApi.md#SupportbundlesDelete) | **Delete** /supportbundles/{id}/ | Delete a Support Bundle
[**SupportbundlesDownload**](SupportbundlesApi.md#SupportbundlesDownload) | **Get** /supportbundles/{id}/download/ | Download a Support Bundle
[**SupportbundlesList**](SupportbundlesApi.md#SupportbundlesList) | **Get** /supportbundles/ | List Support Bundles
[**SupportbundlesRead**](SupportbundlesApi.md#SupportbundlesRead) | **Get** /supportbundles/{id}/ | Return Details of a Support Bundle
[**SupportbundlesUpload**](SupportbundlesApi.md#SupportbundlesUpload) | **Patch** /supportbundles/{id}/upload/ | Upload Support Bundle to AWS S3.

# **SupportbundlesCreate**
> AsyncSupportBundle SupportbundlesCreate(ctx, optional)
Create a Support Bundle

This endpoint creates a support bundle.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SupportbundlesApiSupportbundlesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SupportbundlesApiSupportbundlesCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SupportbundlesBody**](SupportbundlesBody.md)|  | 

### Return type

[**AsyncSupportBundle**](AsyncSupportBundle.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SupportbundlesDelete**
> SupportbundlesDelete(ctx, id)
Delete a Support Bundle

This endpoint deletes a support bundle.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Support Bundle ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SupportbundlesDownload**
> SupportbundlesDownload(ctx, id, secret)
Download a Support Bundle

This endpoint downloads a support bundle.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Support Bundle ID | 
  **secret** | **string**| Bundle secret key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SupportbundlesList**
> []SupportBundle SupportbundlesList(ctx, optional)
List Support Bundles

This endpoint lists support bundles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SupportbundlesApiSupportbundlesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SupportbundlesApiSupportbundlesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 

### Return type

[**[]SupportBundle**](SupportBundle.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SupportbundlesRead**
> SupportBundle SupportbundlesRead(ctx, id)
Return Details of a Support Bundle

This endpoint returns information about a specific support bundle.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**SupportBundle**](SupportBundle.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SupportbundlesUpload**
> AsyncTaskInResponse SupportbundlesUpload(ctx, id, optional)
Upload Support Bundle to AWS S3.

This endpoint uploads a supportbundle to AWS S3.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Support Bundle ID | 
 **optional** | ***SupportbundlesApiSupportbundlesUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SupportbundlesApiSupportbundlesUploadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdUploadBody**](IdUploadBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

