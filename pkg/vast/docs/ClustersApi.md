# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBoxes**](ClustersApi.md#AddBoxes) | **Patch** /clusters/add_boxes/ | Start Add Boxes procedure
[**CloseProtocolHandle**](ClustersApi.md#CloseProtocolHandle) | **Delete** /clusters/close_protocol_handle/ | Close open protocol filehandles
[**ClusterAdvancedMetrics**](ClustersApi.md#ClusterAdvancedMetrics) | **Get** /clusters/{id}/advanced/ | Aggregated information about the cluster
[**ClusterAuditingModify**](ClustersApi.md#ClusterAuditingModify) | **Patch** /clusters/{id}/auditing/ | Modify Cluster Audit Settings
[**ClusterAuditingShow**](ClustersApi.md#ClusterAuditingShow) | **Get** /clusters/{id}/auditing/ | Show Cluster Audit Settings
[**ClusterBlockProviders**](ClustersApi.md#ClusterBlockProviders) | **Patch** /clusters/block_providers/ | This endpoint temporarily blocks all authentication providers
[**ClusterLocksDelete**](ClustersApi.md#ClusterLocksDelete) | **Delete** /clusters/{id}/locks/ | Deletes NLM Locks
[**ClusterLocksList**](ClustersApi.md#ClusterLocksList) | **Post** /clusters/{id}/locks/ | List NLM Locks
[**ClustersCreate**](ClustersApi.md#ClustersCreate) | **Post** /clusters/ | Create Cluster
[**ClustersDelete**](ClustersApi.md#ClustersDelete) | **Delete** /clusters/{id}/ | Delete Cluster
[**ClustersDeleteFolder**](ClustersApi.md#ClustersDeleteFolder) | **Delete** /clusters/{id}/delete_folder/ | Delete Cluster Folder
[**ClustersExpand**](ClustersApi.md#ClustersExpand) | **Post** /clusters/{id}/expand/ | Expand Cluster
[**ClustersList**](ClustersApi.md#ClustersList) | **Get** /clusters/ | List Clusters
[**ClustersPartialUpdate**](ClustersApi.md#ClustersPartialUpdate) | **Patch** /clusters/{id}/ | Manage Cluster
[**ClustersRead**](ClustersApi.md#ClustersRead) | **Get** /clusters/{id}/ | Detail Cluster Properties
[**ClustersRpc**](ClustersApi.md#ClustersRpc) | **Patch** /clusters/{id}/rpc/ | This endpoint runs cluster rpc.
[**ClustersSetPassword**](ClustersApi.md#ClustersSetPassword) | **Patch** /clusters/{id}/set_password/ | Change Cluster Passwords
[**ClustersSystemSettings**](ClustersApi.md#ClustersSystemSettings) | **Patch** /clusters/{id}/system_settings/ | Set Cluster System Settings
[**ClustersUpgrade**](ClustersApi.md#ClustersUpgrade) | **Patch** /clusters/{id}/upgrade/ | Upgrade Cluster
[**ClustersWipe**](ClustersApi.md#ClustersWipe) | **Post** /clusters/wipe/ | Stop Cluster s/w on All Nodes
[**DiscoveredHosts**](ClustersApi.md#DiscoveredHosts) | **Get** /hosts/discovered_hosts/ | Return new discovered hosts
[**GetShardExpansionStatus**](ClustersApi.md#GetShardExpansionStatus) | **Get** /clusters/get_shard_expansion_status/ | System shard expansion status
[**GetSnapshotedPaths**](ClustersApi.md#GetSnapshotedPaths) | **Get** /clusters/get_snapshoted_paths/ | Returns a list of snapshoted paths
[**ListCloneSnapshotedPathsRemote**](ClustersApi.md#ListCloneSnapshotedPathsRemote) | **Get** /clusters/list_clone_snapshoted_paths_remote/ | List snapshots on remote target
[**ListOpenProtocolHandles**](ClustersApi.md#ListOpenProtocolHandles) | **Get** /clusters/list_open_protocol_handles/ | Query open protocol filehandles
[**ListSmbClientConnections**](ClustersApi.md#ListSmbClientConnections) | **Get** /clusters/list_smb_client_connections/ | Query SMB client connections
[**ListSnapshotedPathsRemote**](ClustersApi.md#ListSnapshotedPathsRemote) | **Get** /clusters/list_snapshoted_paths_remote/ | List snapshoted paths on remote target
[**ListTenantsRemote**](ClustersApi.md#ListTenantsRemote) | **Get** /clusters/list_tenants_remote/ | List tenants on remote target
[**NotifyNewVersion**](ClustersApi.md#NotifyNewVersion) | **Post** /clusters/{id}/notify_new_version/ | Notify of New Version Available for Download
[**PreUpgradeValidationExceptions**](ClustersApi.md#PreUpgradeValidationExceptions) | **Get** /clusters/{id}/pre_upgrade_validation_exceptions/ | Run Pre-Upgrade Validation
[**RotateMasterEncryptionGroupKey**](ClustersApi.md#RotateMasterEncryptionGroupKey) | **Post** /clusters/rotate_master_encryption_group_key/ | Rotate Master encryption group key.
[**RunHardwareCheck**](ClustersApi.md#RunHardwareCheck) | **Patch** /clusters/run_hardware_check/ | Run Harwdware Validations
[**ShardExpand**](ClustersApi.md#ShardExpand) | **Post** /clusters/shard_expand/ | Shard expansion
[**ShouldUpload**](ClustersApi.md#ShouldUpload) | **Post** /clusters/{id}/should_upload/ | Detect Need for Bundle Upload
[**StopUpgrade**](ClustersApi.md#StopUpgrade) | **Post** /clusters/{id}/stop_upgrade/ | Stop Running Upgrade
[**UpgradeOptane**](ClustersApi.md#UpgradeOptane) | **Post** /clusters/{id}/upgrade_optane/ | Upgrade Optane NVRAM
[**UpgradeSsd**](ClustersApi.md#UpgradeSsd) | **Post** /clusters/{id}/upgrade_ssd/ | Upgrade SSD
[**UpgradeWithoutFile**](ClustersApi.md#UpgradeWithoutFile) | **Post** /clusters/{id}/upgrade_without_file/ | Upgrade Cluster from Pre-Uploaded Bundle
[**UploadBundle**](ClustersApi.md#UploadBundle) | **Post** /clusters/{id}/upload_bundle/ | Upload Upgrade Bundle to Cluster
[**UploadFromS3**](ClustersApi.md#UploadFromS3) | **Post** /clusters/{id}/upload_from_s3/ | Upload Upgrade Bundle from S3 URL

# **AddBoxes**
> AsyncTaskInResponse AddBoxes(ctx, optional)
Start Add Boxes procedure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClustersApiAddBoxesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiAddBoxesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AddBoxesParams**](AddBoxesParams.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloseProtocolHandle**
> ProtocolHandle CloseProtocolHandle(ctx, filePath, sessionId, sessionHandleUniqueId, optional)
Close open protocol filehandles

This endpoint closes open protocol filehandles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filePath** | **string**| File path | 
  **sessionId** | **string**| Session ID | 
  **sessionHandleUniqueId** | **string**| Handle ID | 
 **optional** | ***ClustersApiCloseProtocolHandleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiCloseProtocolHandleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tenantGuid** | **optional.String**| Tenant GUID | 

### Return type

[**ProtocolHandle**](ProtocolHandle.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterAdvancedMetrics**
> ClusterAdvanced ClusterAdvancedMetrics(ctx, id)
Aggregated information about the cluster

This endpoint returns aggregated information about the cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ClusterAdvanced**](ClusterAdvanced.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterAuditingModify**
> ClusterAuditingReadOnlyWithAsyncTask ClusterAuditingModify(ctx, id, optional)
Modify Cluster Audit Settings

This endpoint updates cluster audit settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***ClustersApiClusterAuditingModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClusterAuditingModifyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ClusterAuditing**](ClusterAuditing.md)|  | 

### Return type

[**ClusterAuditingReadOnlyWithAsyncTask**](ClusterAuditingReadOnlyWithAsyncTask.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterAuditingShow**
> ClusterAuditingReadOnly ClusterAuditingShow(ctx, id)
Show Cluster Audit Settings

This endpoint returns cluster audit settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ClusterAuditingReadOnly**](ClusterAuditingReadOnly.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterBlockProviders**
> ClusterBlockProviders(ctx, optional)
This endpoint temporarily blocks all authentication providers

This endpoint temporarily blocks all authentication providers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClustersApiClusterBlockProvidersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClusterBlockProvidersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ClustersBlockProvidersBody**](ClustersBlockProvidersBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterLocksDelete**
> ClusterLocksDelete(ctx, id, optional)
Deletes NLM Locks

This endpoint deletes NLM locks on a file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***ClustersApiClusterLocksDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClusterLocksDeleteOpts struct
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
 **optional** | ***ClustersApiClusterLocksListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClusterLocksListOpts struct
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

# **ClustersCreate**
> AsyncCluster ClustersCreate(ctx, optional)
Create Cluster

This endpoint creates the cluster.\\ Ideally, VMS will auto-discover the nodes and its logic will be able to decide on\\ exact C/D nodes assignment for the cluster.\\ Therefore, in ideal case this command gets no node-related arguments.\\ Optionally specific C/D nodes list or C/D nodes number can be provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClustersApiClustersCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClustersCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ClustersBody**](ClustersBody.md)|  | 

### Return type

[**AsyncCluster**](AsyncCluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersDelete**
> ClustersDelete(ctx, id)
Delete Cluster

This endpoint deletes the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Cluster ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersDeleteFolder**
> ClustersDeleteFolder(ctx, id, optional)
Delete Cluster Folder

This endpoint deletes a folder from the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ClustersApiClustersDeleteFolderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClustersDeleteFolderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdDeleteFolderBody**](IdDeleteFolderBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersExpand**
> AsyncTaskInResponse ClustersExpand(ctx, id, optional)
Expand Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Cluster ID | 
 **optional** | ***ClustersApiClustersExpandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClustersExpandOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdExpandBody**](IdExpandBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersList**
> []Cluster ClustersList(ctx, optional)
List Clusters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClustersApiClustersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClustersListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 

### Return type

[**[]Cluster**](Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersPartialUpdate**
> AsyncTaskInResponse ClustersPartialUpdate(ctx, id, optional)
Manage Cluster

This endpoint stops/starts the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ClustersApiClustersPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClustersPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ClustersIdBody**](ClustersIdBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersRead**
> Cluster ClustersRead(ctx, id)
Detail Cluster Properties

This endpoint returns details of the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Cluster ID | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersRpc**
> Generic ClustersRpc(ctx, id, optional)
This endpoint runs cluster rpc.

This endpoint runs cluster rpc.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ClustersApiClustersRpcOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClustersRpcOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdRpcBody**](IdRpcBody.md)|  | 

### Return type

[**Generic**](Generic.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersSetPassword**
> AsyncTaskInResponse ClustersSetPassword(ctx, id, optional)
Change Cluster Passwords

This endpoint changes the cluster passwords.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Cluster ID | 
 **optional** | ***ClustersApiClustersSetPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClustersSetPasswordOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdSetPasswordBody**](IdSetPasswordBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersSystemSettings**
> ClustersSystemSettings(ctx, id, optional)
Set Cluster System Settings

This endpoint sets the cluster system settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Cluster ID | 
 **optional** | ***ClustersApiClustersSystemSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClustersSystemSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdSystemSettingsBody**](IdSystemSettingsBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersUpgrade**
> AsyncTaskInResponse ClustersUpgrade(ctx, id, optional)
Upgrade Cluster

This endpoint upgrades the cluster s/w on the nodes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Cluster ID | 
 **optional** | ***ClustersApiClustersUpgradeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClustersUpgradeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdUpgradeBody**](IdUpgradeBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClustersWipe**
> AsyncTaskInResponse ClustersWipe(ctx, optional)
Stop Cluster s/w on All Nodes

This endpoint stops the cluster s/w on all nodes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClustersApiClustersWipeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiClustersWipeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ClustersWipeBody**](ClustersWipeBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoveredHosts**
> []Host DiscoveredHosts(ctx, optional)
Return new discovered hosts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClustersApiDiscoveredHostsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiDiscoveredHostsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostType** | **optional.String**|  | 

### Return type

[**[]Host**](Host.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShardExpansionStatus**
> SystemShardExpandResponse GetShardExpansionStatus(ctx, )
System shard expansion status

This endpoint returns the status of system shard expansion.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemShardExpandResponse**](SystemShardExpandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshotedPaths**
> interface{} GetSnapshotedPaths(ctx, optional)
Returns a list of snapshoted paths

This endpoint returns a list of a snapshoted pathes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClustersApiGetSnapshotedPathsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiGetSnapshotedPathsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **optional.String**| Tenant ID | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCloneSnapshotedPathsRemote**
> interface{} ListCloneSnapshotedPathsRemote(ctx, remoteTargetGuid, handle, optional)
List snapshots on remote target

This endpoint returns the list of snapshots by path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **remoteTargetGuid** | **string**| remote target GUID | 
  **handle** | **string**| Path handle | 
 **optional** | ***ClustersApiListCloneSnapshotedPathsRemoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiListCloneSnapshotedPathsRemoteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startSnapshotId** | **optional.String**| Start snapshot ID | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOpenProtocolHandles**
> ProtocolHandle ListOpenProtocolHandles(ctx, filePath, optional)
Query open protocol filehandles

This endpoint queries open protocol filehandles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filePath** | **string**| File path | 
 **optional** | ***ClustersApiListOpenProtocolHandlesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiListOpenProtocolHandlesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantGuid** | **optional.String**| Tenant GUID | 

### Return type

[**ProtocolHandle**](ProtocolHandle.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSmbClientConnections**
> SmbClientConnection ListSmbClientConnections(ctx, clientIp, optional)
Query SMB client connections

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientIp** | **string**| Client IP | 
 **optional** | ***ClustersApiListSmbClientConnectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiListSmbClientConnectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **optional.Int32**| Tenant id to filter by. | 

### Return type

[**SmbClientConnection**](SmbClientConnection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSnapshotedPathsRemote**
> interface{} ListSnapshotedPathsRemote(ctx, remoteTargetGuid, remoteTenantGuid)
List snapshoted paths on remote target

This endpoint returns the list of snapshoted paths.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **remoteTargetGuid** | **string**| remote target GUID | 
  **remoteTenantGuid** | **string**| remote tenant GUID | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTenantsRemote**
> interface{} ListTenantsRemote(ctx, remoteTargetGuid)
List tenants on remote target

This endpoint returns the list of tenants.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **remoteTargetGuid** | **string**| remote target GUID | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotifyNewVersion**
> NotifyNewVersion(ctx, id, optional)
Notify of New Version Available for Download

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***ClustersApiNotifyNewVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiNotifyNewVersionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdNotifyNewVersionBody**](IdNotifyNewVersionBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PreUpgradeValidationExceptions**
> PreUpgradeValidationExceptions PreUpgradeValidationExceptions(ctx, id)
Run Pre-Upgrade Validation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Cluster ID | 

### Return type

[**PreUpgradeValidationExceptions**](PreUpgradeValidationExceptions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RotateMasterEncryptionGroupKey**
> RotateMasterEncryptionGroupKey(ctx, )
Rotate Master encryption group key.

Rotate Master encryption group key.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunHardwareCheck**
> AsyncTaskInResponse RunHardwareCheck(ctx, optional)
Run Harwdware Validations

This endpoint runs the hardware upgrade validations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClustersApiRunHardwareCheckOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiRunHardwareCheckOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ClustersRunHardwareCheckBody**](ClustersRunHardwareCheckBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShardExpand**
> SystemShardExpandResponse ShardExpand(ctx, optional)
Shard expansion

This endpoint runs system shard expansion.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClustersApiShardExpandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiShardExpandOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ClustersShardExpandBody**](ClustersShardExpandBody.md)|  | 

### Return type

[**SystemShardExpandResponse**](SystemShardExpandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShouldUpload**
> ShouldUpgrade ShouldUpload(ctx, id, optional)
Detect Need for Bundle Upload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Cluster ID | 
 **optional** | ***ClustersApiShouldUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiShouldUploadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdShouldUploadBody**](IdShouldUploadBody.md)|  | 

### Return type

[**ShouldUpgrade**](ShouldUpgrade.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopUpgrade**
> AsyncTaskInResponse StopUpgrade(ctx, id)
Stop Running Upgrade

This endpoint stops a running upgrade flow.

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

# **UpgradeOptane**
> AsyncTaskInResponse UpgradeOptane(ctx, id, optional)
Upgrade Optane NVRAM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Cluster ID | 
 **optional** | ***ClustersApiUpgradeOptaneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiUpgradeOptaneOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdUpgradeOptaneBody**](IdUpgradeOptaneBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeSsd**
> AsyncTaskInResponse UpgradeSsd(ctx, id, optional)
Upgrade SSD

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Cluster ID | 
 **optional** | ***ClustersApiUpgradeSsdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiUpgradeSsdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdUpgradeSsdBody**](IdUpgradeSsdBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeWithoutFile**
> AsyncTaskInResponse UpgradeWithoutFile(ctx, id, optional)
Upgrade Cluster from Pre-Uploaded Bundle

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Cluster ID | 
 **optional** | ***ClustersApiUpgradeWithoutFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiUpgradeWithoutFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdUpgradeWithoutFileBody**](IdUpgradeWithoutFileBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadBundle**
> AsyncTaskInResponse UploadBundle(ctx, buildPackage, skipPrepare, id)
Upload Upgrade Bundle to Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildPackage** | ***os.File*****os.File**|  | 
  **skipPrepare** | **bool**|  | 
  **id** | **int32**| Cluster ID | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadFromS3**
> AsyncTaskInResponse UploadFromS3(ctx, id, optional)
Upload Upgrade Bundle from S3 URL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Cluster ID | 
 **optional** | ***ClustersApiUploadFromS3Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClustersApiUploadFromS3Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdUploadFromS3Body**](IdUploadFromS3Body.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

