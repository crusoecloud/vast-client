# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRemoteTenants**](TenantsApi.md#GetRemoteTenants) | **Get** /tenants/remote_objects/ | Return Details of Remote Tenants
[**GetTenantSameEncryptionGroupTenants**](TenantsApi.md#GetTenantSameEncryptionGroupTenants) | **Get** /tenants/{id}/same_encryption_group_tenants/ | Get tenants with the same encryption group
[**GetTenantVippoolIpRanges**](TenantsApi.md#GetTenantVippoolIpRanges) | **Get** /tenants/{id}/vippool_ip_ranges/ | Get tenant VIP pools IP ranges
[**ReinstateEncryptionGroup**](TenantsApi.md#ReinstateEncryptionGroup) | **Post** /tenants/{id}/reinstate_encryption_group/ | Reinstate tenant&#x27;s encryption group.
[**RevokeEncryptionGroup**](TenantsApi.md#RevokeEncryptionGroup) | **Post** /tenants/{id}/revoke_encryption_group/ | Revoke tenant&#x27;s encryption group.
[**RotateEncryptionGroupKey**](TenantsApi.md#RotateEncryptionGroupKey) | **Post** /tenants/{id}/rotate_encryption_group_key/ | Rotate tenant&#x27;s encryption group key.
[**TenantsCreate**](TenantsApi.md#TenantsCreate) | **Post** /tenants/ | Create Tenant
[**TenantsDelete**](TenantsApi.md#TenantsDelete) | **Delete** /tenants/{id}/ | Delete Tenant
[**TenantsIsOperationHealthy**](TenantsApi.md#TenantsIsOperationHealthy) | **Post** /tenants/{id}/is_operation_healthy | Check whether an operation may be successfully performed
[**TenantsList**](TenantsApi.md#TenantsList) | **Get** /tenants/ | List Tenants
[**TenantsPartialUpdate**](TenantsApi.md#TenantsPartialUpdate) | **Patch** /tenants/{id}/ | Modify Tenant
[**TenantsRead**](TenantsApi.md#TenantsRead) | **Get** /tenants/{id}/ | Return Details of Tenant

# **GetRemoteTenants**
> []RemoteTenant GetRemoteTenants(ctx, optional)
Return Details of Remote Tenants

Returns details of remote tenants.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TenantsApiGetRemoteTenantsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenantsApiGetRemoteTenantsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Filter remote tenants by name | 
 **targetId** | **optional.String**| Filter remote tenants by target | 

### Return type

[**[]RemoteTenant**](RemoteTenant.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTenantSameEncryptionGroupTenants**
> []string GetTenantSameEncryptionGroupTenants(ctx, id)
Get tenants with the same encryption group

Returns names of tenants with the same encryption group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTenantVippoolIpRanges**
> VipPoolIpRanges GetTenantVippoolIpRanges(ctx, id)
Get tenant VIP pools IP ranges

Returns IP ranges of VIP Pools that can be used with the tenant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**VipPoolIpRanges**](VIPPoolIPRanges.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReinstateEncryptionGroup**
> ReinstateEncryptionGroup(ctx, id)
Reinstate tenant's encryption group.

Reinstate tenant's encryption group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeEncryptionGroup**
> RevokeEncryptionGroup(ctx, id)
Revoke tenant's encryption group.

Revoke tenant's encryption group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RotateEncryptionGroupKey**
> RotateEncryptionGroupKey(ctx, id)
Rotate tenant's encryption group key.

Rotate tenant's encryption group key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenantsCreate**
> Tenant TenantsCreate(ctx, optional)
Create Tenant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TenantsApiTenantsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenantsApiTenantsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TenantsBody**](TenantsBody.md)|  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenantsDelete**
> TenantsDelete(ctx, id, optional)
Delete Tenant

Deletes a specified tenant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***TenantsApiTenantsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenantsApiTenantsDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceRemove** | **optional.Bool**| Forcefully remove - ignore existing files. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenantsIsOperationHealthy**
> Warnings TenantsIsOperationHealthy(ctx, id, optional)
Check whether an operation may be successfully performed

Returns warnings about the future state if operation is performed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***TenantsApiTenantsIsOperationHealthyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenantsApiTenantsIsOperationHealthyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdIsOperationHealthyBody**](IdIsOperationHealthyBody.md)|  | 
 **operation** | **optional.**|  | 

### Return type

[**Warnings**](Warnings.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenantsList**
> []Tenant TenantsList(ctx, optional)
List Tenants

List tenants.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TenantsApiTenantsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenantsApiTenantsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **nameIcontains** | **optional.String**| Name to filter by | 

### Return type

[**[]Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenantsPartialUpdate**
> Tenant TenantsPartialUpdate(ctx, id, optional)
Modify Tenant

Modifies a specified tenant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***TenantsApiTenantsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenantsApiTenantsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TenantsIdBody**](TenantsIdBody.md)|  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenantsRead**
> Tenant TenantsRead(ctx, id)
Return Details of Tenant

Returns details of a specified tenant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

