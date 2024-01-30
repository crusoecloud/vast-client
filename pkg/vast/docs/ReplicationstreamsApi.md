# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReplicationstreamsCreate**](ReplicationstreamsApi.md#ReplicationstreamsCreate) | **Post** /replicationstreams/ | Create Replication Stream (deprecated from VAST Cluster 3.4)
[**ReplicationstreamsDelete**](ReplicationstreamsApi.md#ReplicationstreamsDelete) | **Delete** /replicationstreams/{id}/ | Delete a Replication Stream (deprecated from VAST Cluster 3.4)
[**ReplicationstreamsList**](ReplicationstreamsApi.md#ReplicationstreamsList) | **Get** /replicationstreams/ | List Replication Streams (deprecated from VAST Cluster 3.4)
[**ReplicationstreamsPartialUpdate**](ReplicationstreamsApi.md#ReplicationstreamsPartialUpdate) | **Patch** /replicationstreams/{id}/ | Modify a Replication Stream (deprecated from VAST Cluster 3.4)
[**ReplicationstreamsRead**](ReplicationstreamsApi.md#ReplicationstreamsRead) | **Get** /replicationstreams/{id}/ | Return Details of a Replication Stream (deprecated from VAST Cluster 3.4)

# **ReplicationstreamsCreate**
> ReplicationStream ReplicationstreamsCreate(ctx, optional)
Create Replication Stream (deprecated from VAST Cluster 3.4)

This endpoint creates a replication stream. It is deprecated from VAST Cluster 3.4. Use /protectedpaths/.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationstreamsApiReplicationstreamsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationstreamsApiReplicationstreamsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ReplicationstreamsBody**](ReplicationstreamsBody.md)|  | 

### Return type

[**ReplicationStream**](ReplicationStream.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationstreamsDelete**
> ReplicationstreamsDelete(ctx, id)
Delete a Replication Stream (deprecated from VAST Cluster 3.4)

This endpoint deletes a replication stream. It is deprecated from VAST Cluster 3.4. Use /protectedpaths/.

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

# **ReplicationstreamsList**
> []ReplicationStream ReplicationstreamsList(ctx, optional)
List Replication Streams (deprecated from VAST Cluster 3.4)

This endpoint presents the replication streams. It is deprecated from VAST Cluster 3.4. Use /protectedpaths/.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationstreamsApiReplicationstreamsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationstreamsApiReplicationstreamsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **state** | **optional.String**|  | 
 **sourceDir** | **optional.String**|  | 
 **replicationPolicyName** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **enabled** | **optional.String**| start/pause replication stream | 
 **ordering** | **optional.String**| orders by some field | 

### Return type

[**[]ReplicationStream**](ReplicationStream.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationstreamsPartialUpdate**
> ReplicationstreamsPartialUpdate(ctx, id, optional)
Modify a Replication Stream (deprecated from VAST Cluster 3.4)

This endpoint modifies a replication stream. It is deprecated from VAST Cluster 3.4. Use /protectedpaths/.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ReplicationstreamsApiReplicationstreamsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationstreamsApiReplicationstreamsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ReplicationstreamsIdBody**](ReplicationstreamsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationstreamsRead**
> ReplicationStream ReplicationstreamsRead(ctx, id)
Return Details of a Replication Stream (deprecated from VAST Cluster 3.4)

This endpoint returns details of a replication stream. It is deprecated from VAST Cluster 3.4. Use /protectedpaths/.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| index of replication stream | 

### Return type

[**ReplicationStream**](ReplicationStream.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

