# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckPermissionsTemplates**](ViewsApi.md#CheckPermissionsTemplates) | **Post** /views/{id}/check_permissions_templates/ | Check permissions templates
[**CloseSmbHandle**](ViewsApi.md#CloseSmbHandle) | **Delete** /views/close_smb_handle/ | Close open SMB filehandles
[**ListOpenSmbHandles**](ViewsApi.md#ListOpenSmbHandles) | **Get** /views/list_open_smb_handles/ | Query open SMB filehandles
[**ListSeamlessPeers**](ViewsApi.md#ListSeamlessPeers) | **Get** /views/list_seamless_peers/ | Get a list of available seamless peers
[**StartPermissionsRepair**](ViewsApi.md#StartPermissionsRepair) | **Post** /views/{id}/permissions_repair/ | Start bulk permission update
[**StopPermissionsRepair**](ViewsApi.md#StopPermissionsRepair) | **Delete** /views/{id}/permissions_repair/ | Stop bulk permission update
[**ViewsCreate**](ViewsApi.md#ViewsCreate) | **Post** /views/ | Create View
[**ViewsDelete**](ViewsApi.md#ViewsDelete) | **Delete** /views/{id}/ | Delete View
[**ViewsList**](ViewsApi.md#ViewsList) | **Get** /views/ | List Views
[**ViewsPartialUpdate**](ViewsApi.md#ViewsPartialUpdate) | **Patch** /views/{id}/ | Modify a View
[**ViewsRead**](ViewsApi.md#ViewsRead) | **Get** /views/{id}/ | Return Details of a View
[**ViewsUpdate**](ViewsApi.md#ViewsUpdate) | **Put** /views/{id}/ | Modify View.

# **CheckPermissionsTemplates**
> ViewsCheckPermissionsTemplatesResponse CheckPermissionsTemplates(ctx, id, optional)
Check permissions templates

This endpoint checks provided dir and file template for bulk permission update.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| View ID | 
 **optional** | ***ViewsApiCheckPermissionsTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiCheckPermissionsTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdCheckPermissionsTemplatesBody**](IdCheckPermissionsTemplatesBody.md)|  | 

### Return type

[**ViewsCheckPermissionsTemplatesResponse**](ViewsCheckPermissionsTemplatesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloseSmbHandle**
> ProtocolHandle CloseSmbHandle(ctx, filePath, sessionId, sessionHandleUniqueId, optional)
Close open SMB filehandles

This endpoint closes open smb filehandles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filePath** | **string**| File path | 
  **sessionId** | **string**| Session ID | 
  **sessionHandleUniqueId** | **string**| Handle ID | 
 **optional** | ***ViewsApiCloseSmbHandleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiCloseSmbHandleOpts struct
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

# **ListOpenSmbHandles**
> ProtocolHandle ListOpenSmbHandles(ctx, filePath, optional)
Query open SMB filehandles

This endpoint queries open smb filehandles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filePath** | **string**| File path | 
 **optional** | ***ViewsApiListOpenSmbHandlesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiListOpenSmbHandlesOpts struct
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

# **ListSeamlessPeers**
> []SeamlessPeer ListSeamlessPeers(ctx, filePath, tenantGuid)
Get a list of available seamless peers

This endpoint lists available seamless peers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filePath** | **string**| File path | 
  **tenantGuid** | **string**| Tenant GUID | 

### Return type

[**[]SeamlessPeer**](SeamlessPeer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartPermissionsRepair**
> AsyncTaskInResponse StartPermissionsRepair(ctx, id, optional)
Start bulk permission update

This endpoint starts bulk permission update.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| View ID | 
 **optional** | ***ViewsApiStartPermissionsRepairOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiStartPermissionsRepairOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdPermissionsRepairBody**](IdPermissionsRepairBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopPermissionsRepair**
> StopPermissionsRepair(ctx, id)
Stop bulk permission update

This endpoint stops bulk permission update.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| View ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewsCreate**
> View ViewsCreate(ctx, optional)
Create View

This endpoint creates a view.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ViewsApiViewsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiViewsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ViewsBody**](ViewsBody.md)|  | 

### Return type

[**View**](View.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewsDelete**
> ViewsDelete(ctx, id)
Delete View

This endpoint deletes a view.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| View ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewsList**
> []View ViewsList(ctx, optional)
List Views

This endpoint lists views, with optional filtering

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ViewsApiViewsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiViewsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **name** | **optional.String**| Filter by View name | 
 **path** | **optional.String**| Filter by Element Store path | 
 **alias** | **optional.String**| Filter by NFS export alias | 
 **bucket** | **optional.String**|  | 
 **share** | **optional.String**| Filter by share name | 
 **policyName** | **optional.String**| Filter by view policy name | 
 **policyId** | **optional.String**| Filter by view policy ID | 
 **clusterName** | **optional.String**|  | 
 **clusterId** | **optional.String**|  | 
 **tenantId** | **optional.Int32**| Tenant id to filter by. | 
 **tenantNameIcontains** | **optional.String**| Tenant name to filter by | 

### Return type

[**[]View**](View.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewsPartialUpdate**
> ViewsPartialUpdate(ctx, id, optional)
Modify a View

This endpoint modifies a view.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| View ID | 
 **optional** | ***ViewsApiViewsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiViewsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ViewsIdBody1**](ViewsIdBody1.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewsRead**
> View ViewsRead(ctx, id)
Return Details of a View

This endpoint returns details of a view.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the view to show | 

### Return type

[**View**](View.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewsUpdate**
> ViewsUpdate(ctx, id, optional)
Modify View.

This endpoint modifies a view.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| View ID | 
 **optional** | ***ViewsApiViewsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewsApiViewsUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ViewsIdBody**](ViewsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

