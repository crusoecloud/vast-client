# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRemoteMapping**](ViewpoliciesApi.md#AddRemoteMapping) | **Post** /viewpolicies/{id}/remote_mapping/ | 
[**RefreshNetgroups**](ViewpoliciesApi.md#RefreshNetgroups) | **Patch** /viewpolicies/{id}/refresh_netgroups/ | refresh View policy netgroups
[**RemoveRemoteMapping**](ViewpoliciesApi.md#RemoveRemoteMapping) | **Delete** /viewpolicies/{id}/remote_mapping/ | 
[**ViewpoliciesCreate**](ViewpoliciesApi.md#ViewpoliciesCreate) | **Post** /viewpolicies/ | Create View Policy
[**ViewpoliciesDelete**](ViewpoliciesApi.md#ViewpoliciesDelete) | **Delete** /viewpolicies/{id}/ | Delete View Policy.
[**ViewpoliciesList**](ViewpoliciesApi.md#ViewpoliciesList) | **Get** /viewpolicies/ | List View Policies
[**ViewpoliciesPartialUpdate**](ViewpoliciesApi.md#ViewpoliciesPartialUpdate) | **Patch** /viewpolicies/{id}/ | Modify a View Policy
[**ViewpoliciesRead**](ViewpoliciesApi.md#ViewpoliciesRead) | **Get** /viewpolicies/{id}/ | Return Details of a View Policy

# **AddRemoteMapping**
> AddRemoteMapping(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ViewpoliciesApiAddRemoteMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewpoliciesApiAddRemoteMappingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdRemoteMappingBody**](IdRemoteMappingBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshNetgroups**
> RefreshNetgroups(ctx, id)
refresh View policy netgroups

refresh View policy netgroups

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

# **RemoveRemoteMapping**
> RemoveRemoteMapping(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ViewpoliciesApiRemoveRemoteMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewpoliciesApiRemoveRemoteMappingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdRemoteMappingBody1**](IdRemoteMappingBody1.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewpoliciesCreate**
> ViewPolicy ViewpoliciesCreate(ctx, optional)
Create View Policy

This endpoint creates a view policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ViewpoliciesApiViewpoliciesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewpoliciesApiViewpoliciesCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ViewpoliciesBody**](ViewpoliciesBody.md)|  | 

### Return type

[**ViewPolicy**](ViewPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewpoliciesDelete**
> ViewpoliciesDelete(ctx, id)
Delete View Policy.

This endpoint deletes a view policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the view policy to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewpoliciesList**
> []ViewPolicy ViewpoliciesList(ctx, optional)
List View Policies

This endpoint lists view policies.<p>By default, all view policies are returned. The list can be filtered using optional parameters. A view policy is a reusable set of configurations that apply to views. Every view has a view policy. A view policy may be used by multiple views.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ViewpoliciesApiViewpoliciesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewpoliciesApiViewpoliciesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **atimeFrequency** | **optional.String**| Filter by atime frequency. | 
 **page** | **optional.String**|  | 
 **name** | **optional.String**| Filter by name. | 
 **clusterName** | **optional.String**|  | 
 **clusterId** | **optional.String**|  | 
 **nfsReturnOpenPermissions** | **optional.Bool**| Filter by enabled nfs-return-open-permissions flag | 
 **smbFileMode** | **optional.Int32**| Filter by smb_file_mode. smb_file_mode is the default unix permission bits applied to files created by SMB clients. It is relevant only to views that are exposed to both SMB and NFS access protocols and have NFS security flavor. | 
 **smbDirectoryMode** | **optional.Int32**| Filter by smb_directory_mode. smb_directory_mode is the default unix permission bits applied to directories created by SMB clients. It is relevant only to views that are exposed to both SMB and NFS access protocols and have NFS security flavor. | 
 **appleSid** | **optional.Bool**| apple sid | 
 **tenantId** | **optional.Int32**| Tenant id to filter by. | 
 **tenantNameIcontains** | **optional.String**| Tenant name to filter by | 
 **servesTenant** | **optional.String**| Filter by served tenants. Accepts Tenant ID or \&quot;all\&quot;. | 

### Return type

[**[]ViewPolicy**](ViewPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewpoliciesPartialUpdate**
> ViewpoliciesPartialUpdate(ctx, id, optional)
Modify a View Policy

Modify a view policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***ViewpoliciesApiViewpoliciesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ViewpoliciesApiViewpoliciesPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ViewpoliciesIdBody**](ViewpoliciesIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ViewpoliciesRead**
> ViewPolicy ViewpoliciesRead(ctx, id)
Return Details of a View Policy

This endpoint returns a specified view policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the view policy | 

### Return type

[**ViewPolicy**](ViewPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

