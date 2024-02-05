# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupCreate**](GroupsApi.md#GroupCreate) | **Post** /groups/ | Create Group
[**GroupDelete**](GroupsApi.md#GroupDelete) | **Delete** /groups/{id}/ | Delete Group.
[**GroupPartialUpdate**](GroupsApi.md#GroupPartialUpdate) | **Patch** /groups/{id}/ | Modify Group.
[**GroupRead**](GroupsApi.md#GroupRead) | **Get** /groups/{id}/ | Return Details of a Group.
[**GroupsList**](GroupsApi.md#GroupsList) | **Get** /groups/ | List Groups
[**GroupsNames**](GroupsApi.md#GroupsNames) | **Get** /groups/names/ | Find Group by prefix and domain details
[**GroupsQuery**](GroupsApi.md#GroupsQuery) | **Get** /groups/query/ | Query Group.
[**GroupsQueryModify**](GroupsApi.md#GroupsQueryModify) | **Patch** /groups/query/ | Modify non-Local Group

# **GroupCreate**
> Group GroupCreate(ctx, optional)
Create Group

This endpoint adds a group to the local provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GroupsBody**](GroupsBody.md)|  | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupDelete**
> GroupDelete(ctx, id)
Delete Group.

This endpoint deletes a group from the local provider.

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

# **GroupPartialUpdate**
> Group GroupPartialUpdate(ctx, id, optional)
Modify Group.

This endpoint modifies a group on the local provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Group ID | 
 **optional** | ***GroupsApiGroupPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GroupsIdBody**](GroupsIdBody.md)|  | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupRead**
> Group GroupRead(ctx, id)
Return Details of a Group.

This endpoint returns details of a group defined on the local provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Group ID | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsList**
> []Group GroupsList(ctx, optional)
List Groups

This endpoint lists groups.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**[]Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsNames**
> GroupData GroupsNames(ctx, optional)
Find Group by prefix and domain details

This endpoint queries a group by prefix and domain from ActiveDirectory domains

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsNamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **optional.String**| Prefix to find the group | 
 **domain** | **optional.String**| Domain details to find the group (ALL by default). Format: &lt;BASE_DN&gt;|&lt;FQDN&gt;|&lt;SID&gt; | 
 **tenantId** | **optional.Int32**| Tenant id to filter by. | 

### Return type

[**GroupData**](GroupData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsQuery**
> GroupData GroupsQuery(ctx, optional)
Query Group.

This endpoint queries a specified provider for a group based on a group identifier attribute

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsQueryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gid** | **optional.String**| Group GID | 
 **groupname** | **optional.String**| Groupname | 
 **sid** | **optional.String**| Group SID | 
 **vaid** | **optional.String**| Group VAID (a VAST identifier for groups) | 
 **context** | **optional.String**| The provider to query | 
 **tenantId** | **optional.Int32**| Tenant id to filter by. | 

### Return type

[**GroupData**](GroupData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsQueryModify**
> GroupsQueryModify(ctx, optional)
Modify non-Local Group

This endpoint modifies a non-local group's S3 permissions. Specify the user by either NFS UID or SMB user SID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsApiGroupsQueryModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsApiGroupsQueryModifyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GroupsQueryBody**](GroupsQueryBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

