# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaskRetry**](VtasksApi.md#TaskRetry) | **Patch** /vtasks/{id}/retry/ | Retry/Rerun VTask
[**TasksList**](VtasksApi.md#TasksList) | **Get** /vtasks/ | List vtasks
[**TasksPartialUpdate**](VtasksApi.md#TasksPartialUpdate) | **Patch** /vtasks/{id}/ | Modify Vtask
[**TasksRead**](VtasksApi.md#TasksRead) | **Get** /vtasks/{id}/ | Return Task Details

# **TaskRetry**
> TaskRetry(ctx, id, optional)
Retry/Rerun VTask

This endpoint reruns a vtask.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| VTask ID | 
 **optional** | ***VtasksApiTaskRetryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VtasksApiTaskRetryOpts struct
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

# **TasksList**
> []VTask TasksList(ctx, optional)
List vtasks

Get the list of async tasks.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VtasksApiTasksListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VtasksApiTasksListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.String**|  | 
 **page** | **optional.String**|  | 
 **name** | **optional.String**| Task name | 
 **startTimeGte** | **optional.String**|  | 
 **endTimeLte** | **optional.String**|  | 
 **state** | **optional.String**| Task state | 

### Return type

[**[]VTask**](VTask.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TasksPartialUpdate**
> TasksPartialUpdate(ctx, id, optional)
Modify Vtask

This endpoint partially modifies a vtask.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Vtask ID | 
 **optional** | ***VtasksApiTasksPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VtasksApiTasksPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of VtasksIdBody**](VtasksIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TasksRead**
> VTask TasksRead(ctx, id)
Return Task Details

This endpoint returns details of a vtask.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**VTask**](VTask.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

