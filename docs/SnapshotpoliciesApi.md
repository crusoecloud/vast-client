# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SnapshotpoliciesCreate**](SnapshotpoliciesApi.md#SnapshotpoliciesCreate) | **Post** /snapshotpolicies/ | Create a Snapshot Policy (deprecated from VAST Cluster 3.4)
[**SnapshotpoliciesDelete**](SnapshotpoliciesApi.md#SnapshotpoliciesDelete) | **Delete** /snapshotpolicies/{id}/ | Delete a Snapshot Policy (deprecated from VAST Cluster 3.4)
[**SnapshotpoliciesList**](SnapshotpoliciesApi.md#SnapshotpoliciesList) | **Get** /snapshotpolicies/ | List Snapshot Policies (deprecated from VAST Cluster 3.4)
[**SnapshotpoliciesPartialUpdate**](SnapshotpoliciesApi.md#SnapshotpoliciesPartialUpdate) | **Patch** /snapshotpolicies/{id}/ | Modify a Snapshot Policy (deprecated from VAST Cluster 3.4)
[**SnapshotpoliciesRead**](SnapshotpoliciesApi.md#SnapshotpoliciesRead) | **Get** /snapshotpolicies/{id}/ | Return Details of a Snapshot Policy (deprecated from VAST Cluster 3.4)

# **SnapshotpoliciesCreate**
> SnapshotPolicy SnapshotpoliciesCreate(ctx, optional)
Create a Snapshot Policy (deprecated from VAST Cluster 3.4)

This endpoint creates a snapshot policy. It is deprecated from VAST Cluster 3.4. Use /protectionpolicies/.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnapshotpoliciesApiSnapshotpoliciesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotpoliciesApiSnapshotpoliciesCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SnapshotpoliciesBody**](SnapshotpoliciesBody.md)|  | 

### Return type

[**SnapshotPolicy**](SnapshotPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotpoliciesDelete**
> SnapshotpoliciesDelete(ctx, id)
Delete a Snapshot Policy (deprecated from VAST Cluster 3.4)

This endpoint deletes a snapshot policy. It is deprecated from VAST Cluster 3.4. Use /protectionpolicies/.

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

# **SnapshotpoliciesList**
> []SnapshotPolicy SnapshotpoliciesList(ctx, optional)
List Snapshot Policies (deprecated from VAST Cluster 3.4)

This endpoint returns the list of snapshot policies. The list can be filtered. It is deprecated from VAST Cluster 3.4. Use /protectionpolicies/.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnapshotpoliciesApiSnapshotpoliciesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotpoliciesApiSnapshotpoliciesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **optional.String**| Filter by snapshot path | 
 **lastOperationState** | **optional.String**| Filter by last operation state | 
 **schedule** | **optional.String**| Filter by schedule | 
 **enabled** | **optional.Bool**| Filter to return only enabled snapshots | 

### Return type

[**[]SnapshotPolicy**](SnapshotPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotpoliciesPartialUpdate**
> SnapshotpoliciesPartialUpdate(ctx, id, optional)
Modify a Snapshot Policy (deprecated from VAST Cluster 3.4)

This endpoint modifies a snapshot policy. It is deprecated from VAST Cluster 3.4. Use /protectionpolicies/.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SnapshotpoliciesApiSnapshotpoliciesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotpoliciesApiSnapshotpoliciesPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SnapshotpoliciesIdBody**](SnapshotpoliciesIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotpoliciesRead**
> SnapshotPolicy SnapshotpoliciesRead(ctx, id)
Return Details of a Snapshot Policy (deprecated from VAST Cluster 3.4)

This endpoint returns details of a snapshot policy. It is deprecated from VAST Cluster 3.4. Use /protectionpolicies/.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**SnapshotPolicy**](SnapshotPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

