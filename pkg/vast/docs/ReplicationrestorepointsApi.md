# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReplicationrestorepointsList**](ReplicationrestorepointsApi.md#ReplicationrestorepointsList) | **Get** /replicationrestorepoints/ | List Replication Restore Points
[**ReplicationrestorepointsRead**](ReplicationrestorepointsApi.md#ReplicationrestorepointsRead) | **Get** /replicationrestorepoints/{id}/ | Return Details of a Replication Restore Point

# **ReplicationrestorepointsList**
> []ReplicationRestorePoint ReplicationrestorepointsList(ctx, optional)
List Replication Restore Points

This endpoint lists replication restore points.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationrestorepointsApiReplicationrestorepointsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationrestorepointsApiReplicationrestorepointsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **streamId** | **optional.String**|  | 
 **streamName** | **optional.String**|  | 
 **streamNameContains** | **optional.String**|  | 

### Return type

[**[]ReplicationRestorePoint**](ReplicationRestorePoint.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationrestorepointsRead**
> ReplicationRestorePoint ReplicationrestorepointsRead(ctx, id)
Return Details of a Replication Restore Point

This endpoint returns details of a replication restore point.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| index of replication restore point | 

### Return type

[**ReplicationRestorePoint**](ReplicationRestorePoint.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

