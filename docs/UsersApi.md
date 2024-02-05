# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersAdd**](UsersApi.md#UsersAdd) | **Post** /users/ | Add Local User
[**UsersDelete**](UsersApi.md#UsersDelete) | **Delete** /users/{id}/ | Delete User
[**UsersGenerateKey**](UsersApi.md#UsersGenerateKey) | **Post** /users/{id}/access_keys/ | Generate S3 access key pair for a User
[**UsersList**](UsersApi.md#UsersList) | **Get** /users/ | List Users
[**UsersModifyKey**](UsersApi.md#UsersModifyKey) | **Patch** /users/{id}/access_keys/ | Enable/Disable S3 Access Key Pair
[**UsersNames**](UsersApi.md#UsersNames) | **Get** /users/names/ | Find User by prefix and domain details
[**UsersNonLocalGenerateKey**](UsersApi.md#UsersNonLocalGenerateKey) | **Post** /users/non_local_keys/ | Generate S3 Access Key Pair for Non-Local User
[**UsersNonLocalModifyKey**](UsersApi.md#UsersNonLocalModifyKey) | **Patch** /users/non_local_keys/ | Enable or Disable S3 Access Key Pair (non-local user)
[**UsersNonLocalRemoveKey**](UsersApi.md#UsersNonLocalRemoveKey) | **Delete** /users/non_local_keys/ | Remove S3 Access Key Pair from Non-Local User
[**UsersPartialUpdate**](UsersApi.md#UsersPartialUpdate) | **Patch** /users/{id}/ | Modify Local User
[**UsersQuery**](UsersApi.md#UsersQuery) | **Get** /users/query/ | Query User
[**UsersQueryModify**](UsersApi.md#UsersQueryModify) | **Patch** /users/query/ | Modify non-Local User
[**UsersRead**](UsersApi.md#UsersRead) | **Get** /users/{id}/ | Return Local User Details
[**UsersRefresh**](UsersApi.md#UsersRefresh) | **Patch** /users/refresh/ | Refresh User
[**UsersRemoveKey**](UsersApi.md#UsersRemoveKey) | **Delete** /users/{id}/access_keys/ | Remove S3 Access Key Pair

# **UsersAdd**
> User UsersAdd(ctx, optional)
Add Local User

This endpoint adds a user to the local provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersAddOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersBody**](UsersBody.md)|  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersDelete**
> UsersDelete(ctx, id)
Delete User

This endpoint deletes a specified user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| User ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGenerateKey**
> UserKeyPair UsersGenerateKey(ctx, id)
Generate S3 access key pair for a User

This endpoint generates an S3 access key pair for a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| user ID | 

### Return type

[**UserKeyPair**](UserKeyPair.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersList**
> []User UsersList(ctx, optional)
List Users

This endpoint lists local users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **pageSize** | **optional.String**|  | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersModifyKey**
> UsersModifyKey(ctx, id, optional)
Enable/Disable S3 Access Key Pair

This endpoint enables/disables a specified S3 access key pair for a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| user ID | 
 **optional** | ***UsersApiUsersModifyKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersModifyKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdAccessKeysBody1**](IdAccessKeysBody1.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNames**
> UserData UsersNames(ctx, optional)
Find User by prefix and domain details

This endpoint queries a user by prefix and domain from ActiveDirectory domains

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersNamesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **optional.String**| Prefix to find the user | 
 **domain** | **optional.String**| Domain details to find the user (ALL by default). Format: &lt;BASE_DN&gt;|&lt;FQDN&gt;|&lt;SID&gt; | 
 **tenantId** | **optional.Int32**| Tenant id to filter by. | 

### Return type

[**UserData**](UserData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNonLocalGenerateKey**
> UserKeyPair UsersNonLocalGenerateKey(ctx, optional)
Generate S3 Access Key Pair for Non-Local User

This endpoint generates an S3 access key pair for a non local user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersNonLocalGenerateKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersNonLocalGenerateKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersNonLocalKeysBody**](UsersNonLocalKeysBody.md)|  | 

### Return type

[**UserKeyPair**](UserKeyPair.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNonLocalModifyKey**
> UsersNonLocalModifyKey(ctx, optional)
Enable or Disable S3 Access Key Pair (non-local user)

This endpoint enables or disables an S3 access key pair for a non local user whose attributes were already retrieved from an external provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersNonLocalModifyKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersNonLocalModifyKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersNonLocalKeysBody2**](UsersNonLocalKeysBody2.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersNonLocalRemoveKey**
> UsersNonLocalRemoveKey(ctx, optional)
Remove S3 Access Key Pair from Non-Local User

This endpoint removes an S3 access key pair for a non local user whose attributes were already retrieved from an external provider. Specify the user by either NFS UID or SMB user SID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersNonLocalRemoveKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersNonLocalRemoveKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersNonLocalKeysBody1**](UsersNonLocalKeysBody1.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPartialUpdate**
> UsersPartialUpdate(ctx, id, optional)
Modify Local User

This endpoint modifies a local user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| User ID | 
 **optional** | ***UsersApiUsersPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UsersIdBody**](UsersIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersQuery**
> UserQueryData UsersQuery(ctx, optional)
Query User

This endpoint queries a user by an identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersQueryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **optional.String**| NFS UID | 
 **username** | **optional.String**| username | 
 **sid** | **optional.String**| User SID | 
 **vid** | **optional.String**| Vast user ID | 
 **context** | **optional.String**|  | 
 **loginName** | **optional.String**| User login name | 
 **tenantId** | **optional.Int32**| Tenant id to filter by. | 

### Return type

[**UserQueryData**](UserQueryData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersQueryModify**
> UsersQueryModify(ctx, optional)
Modify non-Local User

This endpoint modifies a non-local user's S3 permissions. Specify the user by either NFS UID or SMB user SID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersQueryModifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersQueryModifyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersQueryBody**](UsersQueryBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersRead**
> User UsersRead(ctx, id)
Return Local User Details

This endpoint returns details of a specified user on the local provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| user ID | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersRefresh**
> UsersRefresh(ctx, optional)
Refresh User

This endpoint triggers a fresh query of the user from external providers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersRefreshOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersRefreshOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UsersRefreshBody**](UsersRefreshBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersRemoveKey**
> UsersRemoveKey(ctx, id, optional)
Remove S3 Access Key Pair

This endpoint removes a specified S3 access key pair from a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| User ID | 
 **optional** | ***UsersApiUsersRemoveKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersRemoveKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdAccessKeysBody**](IdAccessKeysBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

