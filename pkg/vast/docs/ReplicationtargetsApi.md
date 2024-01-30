# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReplicationtargetsCreate**](ReplicationtargetsApi.md#ReplicationtargetsCreate) | **Post** /replicationtargets/ | Create an S3 Replication Peer
[**ReplicationtargetsDelete**](ReplicationtargetsApi.md#ReplicationtargetsDelete) | **Delete** /replicationtargets/{id}/ | Delete a Replication Target
[**ReplicationtargetsList**](ReplicationtargetsApi.md#ReplicationtargetsList) | **Get** /replicationtargets/ | List S3 Replication Peers
[**ReplicationtargetsPartialUpdate**](ReplicationtargetsApi.md#ReplicationtargetsPartialUpdate) | **Patch** /replicationtargets/{id}/ | Modify an S3 Replication Peer
[**ReplicationtargetsRead**](ReplicationtargetsApi.md#ReplicationtargetsRead) | **Get** /replicationtargets/{id}/ | Return Details of a Single S3 Replication Peer.

# **ReplicationtargetsCreate**
> ReplicationTarget ReplicationtargetsCreate(ctx, optional)
Create an S3 Replication Peer

This endpoint creates an S3 replication peer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationtargetsApiReplicationtargetsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationtargetsApiReplicationtargetsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ReplicationtargetsBody**](ReplicationtargetsBody.md)|  | 

### Return type

[**ReplicationTarget**](ReplicationTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationtargetsDelete**
> ReplicationtargetsDelete(ctx, id)
Delete a Replication Target

This endpoint deletes an S3 replication peer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Replication Target ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationtargetsList**
> []ReplicationTarget ReplicationtargetsList(ctx, optional)
List S3 Replication Peers

This endpoint lists S3 replication peers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationtargetsApiReplicationtargetsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationtargetsApiReplicationtargetsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **name** | **optional.String**| Filter by name | 
 **bucketName** | **optional.String**| Filter by target bucket name | 
 **httpProtocol** | **optional.String**| Filter by protocol used to connect to the bucket (http or https) | 
 **customBucketUrl** | **optional.String**| Filter by the URL of a custom target bucket | 

### Return type

[**[]ReplicationTarget**](ReplicationTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationtargetsPartialUpdate**
> ReplicationtargetsPartialUpdate(ctx, id, optional)
Modify an S3 Replication Peer

This endpoint modifies an S3 replication peer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Replication Target ID | 
 **optional** | ***ReplicationtargetsApiReplicationtargetsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationtargetsApiReplicationtargetsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ReplicationtargetsIdBody**](ReplicationtargetsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationtargetsRead**
> ReplicationTarget ReplicationtargetsRead(ctx, id)
Return Details of a Single S3 Replication Peer.

This endpoint returns details of a specified S3 replication peer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Replication Target ID | 

### Return type

[**ReplicationTarget**](ReplicationTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

