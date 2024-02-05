# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**S3lifecyclerulesCreate**](S3lifecyclerulesApi.md#S3lifecyclerulesCreate) | **Post** /s3lifecyclerules/ | 
[**S3lifecyclerulesDelete**](S3lifecyclerulesApi.md#S3lifecyclerulesDelete) | **Delete** /s3lifecyclerules/{id}/ | 
[**S3lifecyclerulesGetObjectExpiration**](S3lifecyclerulesApi.md#S3lifecyclerulesGetObjectExpiration) | **Get** /s3lifecyclerules/get_object_expiration/ | 
[**S3lifecyclerulesList**](S3lifecyclerulesApi.md#S3lifecyclerulesList) | **Get** /s3lifecyclerules/ | List S3 life cycle rules
[**S3lifecyclerulesPartialUpdate**](S3lifecyclerulesApi.md#S3lifecyclerulesPartialUpdate) | **Patch** /s3lifecyclerules/{id}/ | 
[**S3lifecyclerulesRead**](S3lifecyclerulesApi.md#S3lifecyclerulesRead) | **Get** /s3lifecyclerules/{id}/ | 

# **S3lifecyclerulesCreate**
> S3LifeCycleRule S3lifecyclerulesCreate(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***S3lifecyclerulesApiS3lifecyclerulesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a S3lifecyclerulesApiS3lifecyclerulesCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of S3lifecyclerulesBody**](S3lifecyclerulesBody.md)|  | 

### Return type

[**S3LifeCycleRule**](S3LifeCycleRule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **S3lifecyclerulesDelete**
> S3lifecyclerulesDelete(ctx, id)


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

# **S3lifecyclerulesGetObjectExpiration**
> S3BucketObjectData S3lifecyclerulesGetObjectExpiration(ctx, objectName, bucketName)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectName** | **string**|  | 
  **bucketName** | **string**|  | 

### Return type

[**S3BucketObjectData**](S3BucketObjectData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **S3lifecyclerulesList**
> []S3LifeCycleRule S3lifecyclerulesList(ctx, optional)
List S3 life cycle rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***S3lifecyclerulesApiS3lifecyclerulesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a S3lifecyclerulesApiS3lifecyclerulesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **pageSize** | **optional.String**|  | 
 **viewId** | **optional.String**|  | 

### Return type

[**[]S3LifeCycleRule**](S3LifeCycleRule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **S3lifecyclerulesPartialUpdate**
> S3lifecyclerulesPartialUpdate(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***S3lifecyclerulesApiS3lifecyclerulesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a S3lifecyclerulesApiS3lifecyclerulesPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of S3lifecyclerulesIdBody**](S3lifecyclerulesIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **S3lifecyclerulesRead**
> S3LifeCycleRule S3lifecyclerulesRead(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**S3LifeCycleRule**](S3LifeCycleRule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

