# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SnapshotsClone**](SnapshotsApi.md#SnapshotsClone) | **Post** /snapshots/{id}/clone/ | Clones snapshot.
[**SnapshotsCreate**](SnapshotsApi.md#SnapshotsCreate) | **Post** /snapshots/ | Create Snapshot
[**SnapshotsDelete**](SnapshotsApi.md#SnapshotsDelete) | **Delete** /snapshots/{id}/ | Delete a Snapshot
[**SnapshotsList**](SnapshotsApi.md#SnapshotsList) | **Get** /snapshots/ | List Snapshots
[**SnapshotsPartialUpdate**](SnapshotsApi.md#SnapshotsPartialUpdate) | **Patch** /snapshots/{id}/ | Modify a Snapshot
[**SnapshotsRead**](SnapshotsApi.md#SnapshotsRead) | **Get** /snapshots/{id}/ | Return Details of a Snapshot

# **SnapshotsClone**
> AsyncGlobalSnapStream SnapshotsClone(ctx, id, optional)
Clones snapshot.

This endpoint clones selected snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***SnapshotsApiSnapshotsCloneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotsApiSnapshotsCloneOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdCloneBody**](IdCloneBody.md)|  | 

### Return type

[**AsyncGlobalSnapStream**](AsyncGlobalSnapStream.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotsCreate**
> Snapshot SnapshotsCreate(ctx, optional)
Create Snapshot

This endpoint creates a snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnapshotsApiSnapshotsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotsApiSnapshotsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SnapshotsBody**](SnapshotsBody.md)|  | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotsDelete**
> SnapshotsDelete(ctx, id)
Delete a Snapshot

This endpoint deletes a snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Specify the ID of the snapshot. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotsList**
> []Snapshot SnapshotsList(ctx, optional)
List Snapshots

This endpoint list snapshots. The list can be filtered.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnapshotsApiSnapshotsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotsApiSnapshotsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **path** | **optional.String**| Filter by snapshot path | 
 **nameContains** | **optional.String**| Filter by part of snapshot name | 
 **expirationTime** | **optional.String**| Filter by expiration time | 
 **protectionPolicyName** | **optional.String**| Filter by snapshot policy name | 
 **protectionPolicyId** | **optional.Int32**| Filter by snapshot policy ID | 
 **state** | **optional.String**| Filter by state | 
 **created** | **optional.String**| Filter by creation time | 
 **locked** | **optional.Bool**| Filter for locked snapshots | 
 **tenantId** | **optional.Int32**| Tenant id to filter by. | 
 **tenantNameIcontains** | **optional.String**| Tenant name to filter by | 

### Return type

[**[]Snapshot**](Snapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotsPartialUpdate**
> SnapshotsPartialUpdate(ctx, id, optional)
Modify a Snapshot

This endpoint modifies a snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Snapshot ID | 
 **optional** | ***SnapshotsApiSnapshotsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotsApiSnapshotsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SnapshotsIdBody**](SnapshotsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnapshotsRead**
> Snapshot SnapshotsRead(ctx, id)
Return Details of a Snapshot

This endpoint returns details of a snapshot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Snapshot ID | 

### Return type

[**Snapshot**](Snapshot.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

