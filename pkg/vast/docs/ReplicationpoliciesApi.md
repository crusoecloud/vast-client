# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReplicationpoliciesCreate**](ReplicationpoliciesApi.md#ReplicationpoliciesCreate) | **Post** /replicationpolicies/ | Create Replication Policy (deprecated from VAST Cluster 3.4)
[**ReplicationpoliciesDelete**](ReplicationpoliciesApi.md#ReplicationpoliciesDelete) | **Delete** /replicationpolicies/{id}/ | Delete a Replication Policy (deprecated from VAST Cluster 3.4)
[**ReplicationpoliciesList**](ReplicationpoliciesApi.md#ReplicationpoliciesList) | **Get** /replicationpolicies/ | List Replication Policies (deprecated from VAST Cluster 3.4)
[**ReplicationpoliciesPartialUpdate**](ReplicationpoliciesApi.md#ReplicationpoliciesPartialUpdate) | **Patch** /replicationpolicies/{id}/ | Modify a Replication Policy (deprecated from VAST Cluster 3.4)
[**ReplicationpoliciesRead**](ReplicationpoliciesApi.md#ReplicationpoliciesRead) | **Get** /replicationpolicies/{id}/ | Return Details of a Replication Policy (deprecated from VAST Cluster 3.4)

# **ReplicationpoliciesCreate**
> ReplicationPolicy ReplicationpoliciesCreate(ctx, optional)
Create Replication Policy (deprecated from VAST Cluster 3.4)

This endpoint creates a replication policy. It is deprecated from VAST Cluster 3.4. Use /protectionpolicies/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationpoliciesApiReplicationpoliciesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationpoliciesApiReplicationpoliciesCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ReplicationpoliciesBody**](ReplicationpoliciesBody.md)|  | 

### Return type

[**ReplicationPolicy**](ReplicationPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationpoliciesDelete**
> ReplicationpoliciesDelete(ctx, id)
Delete a Replication Policy (deprecated from VAST Cluster 3.4)

This endpoint deletes a replication policy. It is deprecated from VAST Cluster 3.4. Use /protectionpolicies/.

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

# **ReplicationpoliciesList**
> []ReplicationPolicy ReplicationpoliciesList(ctx, optional)
List Replication Policies (deprecated from VAST Cluster 3.4)

This endpoint presents the replication policies. It is deprecated from VAST Cluster 3.4. Use /protectionpolicies/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationpoliciesApiReplicationpoliciesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationpoliciesApiReplicationpoliciesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **replicationTargetName** | **optional.String**| Filter by replication peer | 
 **name** | **optional.String**| Filter by replication policy name | 

### Return type

[**[]ReplicationPolicy**](ReplicationPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationpoliciesPartialUpdate**
> ReplicationpoliciesPartialUpdate(ctx, id, optional)
Modify a Replication Policy (deprecated from VAST Cluster 3.4)

This endpoint modifies a replication policy. It is deprecated from VAST Cluster 3.4. Use /protectionpolicies/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ReplicationpoliciesApiReplicationpoliciesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationpoliciesApiReplicationpoliciesPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ReplicationpoliciesIdBody**](ReplicationpoliciesIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationpoliciesRead**
> ReplicationPolicy ReplicationpoliciesRead(ctx, id)
Return Details of a Replication Policy (deprecated from VAST Cluster 3.4)

This endpoint returns details of a replication policy. It is deprecated from VAST Cluster 3.4. Use /protectionpolicies/

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| index of replication policy | 

### Return type

[**ReplicationPolicy**](ReplicationPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

