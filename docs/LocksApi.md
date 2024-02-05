# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterLocksDelete**](LocksApi.md#ClusterLocksDelete) | **Delete** /clusters/{id}/locks/ | Deletes NLM Locks
[**ClusterLocksList**](LocksApi.md#ClusterLocksList) | **Post** /clusters/{id}/locks/ | List NLM Locks

# **ClusterLocksDelete**
> ClusterLocksDelete(ctx, id, optional)
Deletes NLM Locks

This endpoint deletes NLM locks on a file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***LocksApiClusterLocksDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocksApiClusterLocksDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ClusterLocksParams**](ClusterLocksParams.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterLocksList**
> []Lock ClusterLocksList(ctx, id, optional)
List NLM Locks

This endpoint returns all NLM locks on a file at a specified path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***LocksApiClusterLocksListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocksApiClusterLocksListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ClusterLocksParams**](ClusterLocksParams.md)|  | 

### Return type

[**[]Lock**](Lock.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

