# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VippoolsCreate**](VippoolsApi.md#VippoolsCreate) | **Post** /vippools/ | Create VIP Pool
[**VippoolsDelete**](VippoolsApi.md#VippoolsDelete) | **Delete** /vippools/{id}/ | Delete VIP Pool
[**VippoolsList**](VippoolsApi.md#VippoolsList) | **Get** /vippools/ |  List VIP Pools
[**VippoolsPartialUpdate**](VippoolsApi.md#VippoolsPartialUpdate) | **Patch** /vippools/{id}/ | Modify VIP Pool
[**VippoolsRead**](VippoolsApi.md#VippoolsRead) | **Get** /vippools/{id}/ | Return Details of a VIP Pool

# **VippoolsCreate**
> VipPool VippoolsCreate(ctx, optional)
Create VIP Pool

This endpoint creates a VIP pool for client access or for native replication.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VippoolsApiVippoolsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VippoolsApiVippoolsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VippoolsBody**](VippoolsBody.md)|  | 

### Return type

[**VipPool**](VIPPool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VippoolsDelete**
> VippoolsDelete(ctx, id)
Delete VIP Pool

This endpoint deletes a VIP pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| VIP Pool ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VippoolsList**
> []VipPool VippoolsList(ctx, optional)
 List VIP Pools

This endpoint lists VIP pools.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VippoolsApiVippoolsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VippoolsApiVippoolsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **startIp** | **optional.String**| Filter by start IP of VIP pool range | 
 **endIp** | **optional.String**| Filter by end IP of VIP pool range | 
 **clusterName** | **optional.String**|  | 
 **clusterId** | **optional.Int32**|  | 
 **portMembership** | **optional.Int32**|  | 
 **cnodeIdsIcontains** | [**optional.Interface of []int32**](int32.md)|  | 
 **tenantId** | **optional.Int32**| Tenant id to filter by. | 
 **tenantNameIcontains** | **optional.String**| Tenant name to filter by | 
 **servesTenant** | **optional.String**| Filter by served tenants. Accepts Tenant ID or \&quot;all\&quot;. | 

### Return type

[**[]VipPool**](VIPPool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VippoolsPartialUpdate**
> VippoolsPartialUpdate(ctx, id, optional)
Modify VIP Pool

This endpoint modifies a VIP pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| VIP Pool ID | 
 **optional** | ***VippoolsApiVippoolsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VippoolsApiVippoolsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of VippoolsIdBody**](VippoolsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VippoolsRead**
> VipPool VippoolsRead(ctx, id)
Return Details of a VIP Pool

This endpoint returns details of a VIP pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| VIP Pool ID | 

### Return type

[**VipPool**](VIPPool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

