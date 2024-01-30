# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProtectedpathsAddStream**](ProtectedpathsApi.md#ProtectedpathsAddStream) | **Patch** /protectedpaths/{id}/add_stream/ | Add a Replication Stream to a group connected to this Protected Path
[**ProtectedpathsCommit**](ProtectedpathsApi.md#ProtectedpathsCommit) | **Patch** /protectedpaths/{id}/commit | Commits snapshot restore result to protected path.
[**ProtectedpathsCreate**](ProtectedpathsApi.md#ProtectedpathsCreate) | **Post** /protectedpaths/ | Create a Protected Path
[**ProtectedpathsDelete**](ProtectedpathsApi.md#ProtectedpathsDelete) | **Delete** /protectedpaths/{id}/ | Delete a Protected Path
[**ProtectedpathsForceFailover**](ProtectedpathsApi.md#ProtectedpathsForceFailover) | **Patch** /protectedpaths/{id}/force_failover | Force failover protected path
[**ProtectedpathsList**](ProtectedpathsApi.md#ProtectedpathsList) | **Get** /protectedpaths/ | List Protected Paths
[**ProtectedpathsPartialUpdate**](ProtectedpathsApi.md#ProtectedpathsPartialUpdate) | **Patch** /protectedpaths/{id}/ | Modify a Protected Path
[**ProtectedpathsPause**](ProtectedpathsApi.md#ProtectedpathsPause) | **Patch** /protectedpaths/{id}/pause | Pause snapshot restore on protected path.
[**ProtectedpathsRead**](ProtectedpathsApi.md#ProtectedpathsRead) | **Get** /protectedpaths/{id}/ | Return Details of a Protected Path
[**ProtectedpathsReattachStream**](ProtectedpathsApi.md#ProtectedpathsReattachStream) | **Patch** /protectedpaths/{id}/reattach_stream/ | Reattach a Stream to this Protected Path
[**ProtectedpathsRemoveStream**](ProtectedpathsApi.md#ProtectedpathsRemoveStream) | **Patch** /protectedpaths/{id}/remove_stream/ | Remove a Stream from the group connected to this Protected Path
[**ProtectedpathsRestore**](ProtectedpathsApi.md#ProtectedpathsRestore) | **Post** /protectedpaths/{id}/restore | Restores selected snapshot on protected path.
[**ProtectedpathsResume**](ProtectedpathsApi.md#ProtectedpathsResume) | **Patch** /protectedpaths/{id}/resume | Resume snapshot restore on protected path.
[**ProtectedpathsStop**](ProtectedpathsApi.md#ProtectedpathsStop) | **Patch** /protectedpaths/{id}/stop | Stops snapshot restore on protected path.
[**ProtectedpathsValidate**](ProtectedpathsApi.md#ProtectedpathsValidate) | **Get** /protectedpaths/{id}/validate | Returns validations results for Protected Path.

# **ProtectedpathsAddStream**
> ProtectedpathsAddStream(ctx, id, optional)
Add a Replication Stream to a group connected to this Protected Path

This endpoint adds Replication Stream to a group connected to this Protected Path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProtectedpathsApiProtectedpathsAddStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectedpathsApiProtectedpathsAddStreamOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdAddStreamBody**](IdAddStreamBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectedpathsCommit**
> ProtectedpathsCommit(ctx, id, optional)
Commits snapshot restore result to protected path.

This endpoint commits snapshot restore result to protected path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProtectedpathsApiProtectedpathsCommitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectedpathsApiProtectedpathsCommitOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectedpathsCreate**
> ProtectedPath ProtectedpathsCreate(ctx, optional)
Create a Protected Path

This endpoint creates a protected path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectedpathsApiProtectedpathsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectedpathsApiProtectedpathsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ProtectedpathsBody**](ProtectedpathsBody.md)|  | 

### Return type

[**ProtectedPath**](ProtectedPath.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectedpathsDelete**
> AsyncTaskInResponse ProtectedpathsDelete(ctx, id)
Delete a Protected Path

This endpoint deletes a protected path.

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

# **ProtectedpathsForceFailover**
> ProtectedpathsForceFailover(ctx, id)
Force failover protected path

This endpoint performs force fail-over on protected path.

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

# **ProtectedpathsList**
> []ProtectedPath ProtectedpathsList(ctx, optional)
List Protected Paths

This endpoint lists protected paths.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectedpathsApiProtectedpathsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectedpathsApiProtectedpathsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **state** | **optional.String**|  | 
 **role** | **optional.String**|  | 
 **sourceDir** | **optional.String**|  | 
 **replicationPolicyName** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **enabled** | **optional.String**| start/pause protected path | 

### Return type

[**[]ProtectedPath**](ProtectedPath.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectedpathsPartialUpdate**
> ProtectedpathsPartialUpdate(ctx, id, optional)
Modify a Protected Path

This endpoint modifies a protected path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProtectedpathsApiProtectedpathsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectedpathsApiProtectedpathsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProtectedpathsIdBody**](ProtectedpathsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectedpathsPause**
> ProtectedpathsPause(ctx, id, optional)
Pause snapshot restore on protected path.

This endpoint pauses snapshot restore on protected path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProtectedpathsApiProtectedpathsPauseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectedpathsApiProtectedpathsPauseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdPauseBody**](IdPauseBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectedpathsRead**
> ProtectedPath ProtectedpathsRead(ctx, id)
Return Details of a Protected Path

This endpoint returns details of a protected path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| index of protected path | 

### Return type

[**ProtectedPath**](ProtectedPath.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectedpathsReattachStream**
> ProtectedpathsReattachStream(ctx, id, optional)
Reattach a Stream to this Protected Path

This endpoint re-attaches a replication stream to a current protected path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProtectedpathsApiProtectedpathsReattachStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectedpathsApiProtectedpathsReattachStreamOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdReattachStreamBody**](IdReattachStreamBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectedpathsRemoveStream**
> ProtectedpathsRemoveStream(ctx, id, optional)
Remove a Stream from the group connected to this Protected Path

This endpoint removes a replication stream from current protected path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProtectedpathsApiProtectedpathsRemoveStreamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectedpathsApiProtectedpathsRemoveStreamOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdRemoveStreamBody**](IdRemoveStreamBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectedpathsRestore**
> AsyncGlobalSnapStream ProtectedpathsRestore(ctx, id, optional)
Restores selected snapshot on protected path.

This endpoint restores selected snapshot on protected path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProtectedpathsApiProtectedpathsRestoreOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectedpathsApiProtectedpathsRestoreOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdRestoreBody**](IdRestoreBody.md)|  | 

### Return type

[**AsyncGlobalSnapStream**](AsyncGlobalSnapStream.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectedpathsResume**
> ProtectedpathsResume(ctx, id, optional)
Resume snapshot restore on protected path.

This endpoint resumes snapshot restore on protected path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ProtectedpathsApiProtectedpathsResumeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectedpathsApiProtectedpathsResumeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdResumeBody**](IdResumeBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectedpathsStop**
> ProtectedpathsStop(ctx, id)
Stops snapshot restore on protected path.

This endpoint stops snapshot restore on protected path.

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

# **ProtectedpathsValidate**
> InlineResponse2002 ProtectedpathsValidate(ctx, id)
Returns validations results for Protected Path.

This endpoint returns validations results for Protected Path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

