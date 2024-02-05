# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRemoteMapping**](NativereplicationremotetargetsApi.md#GetRemoteMapping) | **Get** /nativereplicationremotetargets/get_remote_mapping/ | Get the mapping between targets and remote view poicies
[**NativereplicationremotetargetsCreate**](NativereplicationremotetargetsApi.md#NativereplicationremotetargetsCreate) | **Post** /nativereplicationremotetargets/ | Create a Native Replication Peer
[**NativereplicationremotetargetsDelete**](NativereplicationremotetargetsApi.md#NativereplicationremotetargetsDelete) | **Delete** /nativereplicationremotetargets/{id}/ | Delete a Native Replication Peer
[**NativereplicationremotetargetsList**](NativereplicationremotetargetsApi.md#NativereplicationremotetargetsList) | **Get** /nativereplicationremotetargets/ | List Native Replication Peers
[**NativereplicationremotetargetsPartialUpdate**](NativereplicationremotetargetsApi.md#NativereplicationremotetargetsPartialUpdate) | **Patch** /nativereplicationremotetargets/{id}/ | Modify a Native Replication Peer (aka Native Replication Remote Target)
[**NativereplicationremotetargetsRead**](NativereplicationremotetargetsApi.md#NativereplicationremotetargetsRead) | **Get** /nativereplicationremotetargets/{id}/ | Return Details of a Native Replication Peer

# **GetRemoteMapping**
> interface{} GetRemoteMapping(ctx, )
Get the mapping between targets and remote view poicies

This endpoint presents the remote mapping

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NativereplicationremotetargetsCreate**
> NativeReplicationRemoteTarget NativereplicationremotetargetsCreate(ctx, optional)
Create a Native Replication Peer

This endpoint creates a native replication peer configuration between the local cluster and a remote cluster. Prerequisite: a VIP pool on each cluster with role=replication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NativereplicationremotetargetsApiNativereplicationremotetargetsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NativereplicationremotetargetsApiNativereplicationremotetargetsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NativereplicationremotetargetsBody**](NativereplicationremotetargetsBody.md)|  | 

### Return type

[**NativeReplicationRemoteTarget**](NativeReplicationRemoteTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NativereplicationremotetargetsDelete**
> NativereplicationremotetargetsDelete(ctx, id)
Delete a Native Replication Peer

This endpoint deletes a native replication peer (aka NativeReplicationRemoteTarget).

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

# **NativereplicationremotetargetsList**
> []NativeReplicationRemoteTarget NativereplicationremotetargetsList(ctx, optional)
List Native Replication Peers

This endpoint presents the native replication remote targets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NativereplicationremotetargetsApiNativereplicationremotetargetsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NativereplicationremotetargetsApiNativereplicationremotetargetsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **name** | **optional.String**| Filter by name | 
 **leadingVip** | **optional.String**| Filter by leading remote VIP, a VIP that is specified when creating the peer. It is one of the VIPs in the remote peer&#x27;s replication VIP Pool. | 
 **pool** | **optional.String**| Filter by the name of the local cluster&#x27;s replication VIP pool | 
 **version** | **optional.String**| Filter by local cluster&#x27;s software version | 
 **remoteVersion** | **optional.String**| Filter by remote peer&#x27;s software version | 
 **lastHeartBeat** | **optional.String**| Filter by last heartbeat, the time of the last successful message sent, arrived and acknowledged by the peer. | 
 **spaceLeft** | **optional.String**| Filter by logical capacity remaining available on the remote peer.  | 
 **remoteVips** | **optional.String**| remote vips | 
 **state** | **optional.String**| Filter by state | 
 **secureMode** | **optional.String**| Filter by secure_mode | 

### Return type

[**[]NativeReplicationRemoteTarget**](NativeReplicationRemoteTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NativereplicationremotetargetsPartialUpdate**
> NativereplicationremotetargetsPartialUpdate(ctx, id, optional)
Modify a Native Replication Peer (aka Native Replication Remote Target)

This endpoint modifies a native replication peer (aka NativeReplicationRemoteTarget).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the Native Replication Remote Target to modify | 
 **optional** | ***NativereplicationremotetargetsApiNativereplicationremotetargetsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NativereplicationremotetargetsApiNativereplicationremotetargetsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NativereplicationremotetargetsIdBody**](NativereplicationremotetargetsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NativereplicationremotetargetsRead**
> NativeReplicationRemoteTarget NativereplicationremotetargetsRead(ctx, id)
Return Details of a Native Replication Peer

This endpoint returns details of a native replication peer, also known as a replication target.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| index of replication target | 

### Return type

[**NativeReplicationRemoteTarget**](NativeReplicationRemoteTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

