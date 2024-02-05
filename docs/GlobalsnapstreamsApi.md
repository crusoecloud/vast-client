# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalsnapstreamsCreate**](GlobalsnapstreamsApi.md#GlobalsnapstreamsCreate) | **Post** /globalsnapstreams/ | Create Global Snap Stream
[**GlobalsnapstreamsDelete**](GlobalsnapstreamsApi.md#GlobalsnapstreamsDelete) | **Delete** /globalsnapstreams/{id}/ | Delete Global Snap Stream
[**GlobalsnapstreamsList**](GlobalsnapstreamsApi.md#GlobalsnapstreamsList) | **Get** /globalsnapstreams/ | Return Global Snap Streams
[**GlobalsnapstreamsPartialUpdate**](GlobalsnapstreamsApi.md#GlobalsnapstreamsPartialUpdate) | **Patch** /globalsnapstreams/{id}/ | Modify a Global Snap Stream
[**GlobalsnapstreamsPause**](GlobalsnapstreamsApi.md#GlobalsnapstreamsPause) | **Patch** /globalsnapstreams/{id}/pause | Pause a Global Snap Stream
[**GlobalsnapstreamsRead**](GlobalsnapstreamsApi.md#GlobalsnapstreamsRead) | **Get** /globalsnapstreams/{id}/ | Return Details of a Global Snap Stream
[**GlobalsnapstreamsResume**](GlobalsnapstreamsApi.md#GlobalsnapstreamsResume) | **Patch** /globalsnapstreams/{id}/resume | REsume a Global Snap Stream
[**GlobalsnapstreamsStop**](GlobalsnapstreamsApi.md#GlobalsnapstreamsStop) | **Patch** /globalsnapstreams/{id}/stop | Stop a Global Snap Stream

# **GlobalsnapstreamsCreate**
> AsyncGlobalSnapStream GlobalsnapstreamsCreate(ctx, optional)
Create Global Snap Stream

This endpoint creates a Global Snap Stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GlobalsnapstreamsApiGlobalsnapstreamsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalsnapstreamsApiGlobalsnapstreamsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GlobalsnapstreamsBody**](GlobalsnapstreamsBody.md)|  | 

### Return type

[**AsyncGlobalSnapStream**](AsyncGlobalSnapStream.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalsnapstreamsDelete**
> AsyncTaskInResponse GlobalsnapstreamsDelete(ctx, id, optional)
Delete Global Snap Stream

This endpoint deletes a Global Snap Stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Global Snap Stream ID | 
 **optional** | ***GlobalsnapstreamsApiGlobalsnapstreamsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalsnapstreamsApiGlobalsnapstreamsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GlobalsnapstreamsIdBody**](GlobalsnapstreamsIdBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalsnapstreamsList**
> []GlobalSnapStream GlobalsnapstreamsList(ctx, optional)
Return Global Snap Streams

This endpoint lists Global Snap Streams.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GlobalsnapstreamsApiGlobalsnapstreamsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalsnapstreamsApiGlobalsnapstreamsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **name** | **optional.String**| Filter by Global Snap Stream name | 

### Return type

[**[]GlobalSnapStream**](GlobalSnapStream.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalsnapstreamsPartialUpdate**
> GlobalsnapstreamsPartialUpdate(ctx, id, optional)
Modify a Global Snap Stream

This endpoint modifies a Global Snap Stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Global Snap Stream ID | 
 **optional** | ***GlobalsnapstreamsApiGlobalsnapstreamsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalsnapstreamsApiGlobalsnapstreamsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GlobalsnapstreamsIdBody1**](GlobalsnapstreamsIdBody1.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalsnapstreamsPause**
> GlobalsnapstreamsPause(ctx, id, optional)
Pause a Global Snap Stream

This endpoint pauses a Global Snap Stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Global Snap Stream ID | 
 **optional** | ***GlobalsnapstreamsApiGlobalsnapstreamsPauseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalsnapstreamsApiGlobalsnapstreamsPauseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdPauseBody1**](IdPauseBody1.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalsnapstreamsRead**
> GlobalSnapStream GlobalsnapstreamsRead(ctx, id)
Return Details of a Global Snap Stream

This endpoint returns details of a specified Global Snap Stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Global Snap Stream ID | 

### Return type

[**GlobalSnapStream**](GlobalSnapStream.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalsnapstreamsResume**
> GlobalsnapstreamsResume(ctx, id, optional)
REsume a Global Snap Stream

This endpoint resumes a Global Snap Stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Global Snap Stream ID | 
 **optional** | ***GlobalsnapstreamsApiGlobalsnapstreamsResumeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalsnapstreamsApiGlobalsnapstreamsResumeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdResumeBody1**](IdResumeBody1.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalsnapstreamsStop**
> AsyncTaskInResponse GlobalsnapstreamsStop(ctx, id)
Stop a Global Snap Stream

This endpoint stops a Global Snap Stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Global Snap Stream ID | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

