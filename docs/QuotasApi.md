# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuotaRefreshUserQuotas**](QuotasApi.md#QuotaRefreshUserQuotas) | **Patch** /quotas/{id}/refresh_user_quotas/ | Refresh user quotas of a quota
[**QuotasCreate**](QuotasApi.md#QuotasCreate) | **Post** /quotas/ | Create a Quota
[**QuotasDelete**](QuotasApi.md#QuotasDelete) | **Delete** /quotas/{id}/ | Delete Quota
[**QuotasList**](QuotasApi.md#QuotasList) | **Get** /quotas/ | List Quotas
[**QuotasPartialUpdate**](QuotasApi.md#QuotasPartialUpdate) | **Patch** /quotas/{id}/ | Modify Quota
[**QuotasRead**](QuotasApi.md#QuotasRead) | **Get** /quotas/{id}/ | Return details of a quota.
[**QuotasRecalc**](QuotasApi.md#QuotasRecalc) | **Patch** /quotas/recalc/ | Start recalculation flow all quotas
[**QuotasRecalcStop**](QuotasApi.md#QuotasRecalcStop) | **Patch** /quotas/recalc_stop/ | Stop recalculation flow for all quotas
[**QuotasResetGracePeriod**](QuotasApi.md#QuotasResetGracePeriod) | **Patch** /quotas/{id}/reset_grace_period/ | Rest the grace period countdown after soft limit is exceeded

# **QuotaRefreshUserQuotas**
> QuotaRefreshUserQuotas(ctx, id)
Refresh user quotas of a quota

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Quota ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuotasCreate**
> Quota QuotasCreate(ctx, optional)
Create a Quota

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QuotasApiQuotasCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuotasApiQuotasCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of QuotasBody**](QuotasBody.md)|  | 

### Return type

[**Quota**](Quota.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuotasDelete**
> QuotasDelete(ctx, id)
Delete Quota

This endpoint deletes a quota.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Quota ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuotasList**
> []Quota QuotasList(ctx, optional)
List Quotas

This endpoint lists quotas. The parameters enable you to filter the list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QuotasApiQuotasListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuotasApiQuotasListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **pageSize** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **hardLimit** | **optional.String**|  | 
 **softLimit** | **optional.String**|  | 
 **softLimitInodes** | **optional.String**|  | 
 **hardLimitInodes** | **optional.String**|  | 
 **systemId** | **optional.String**|  | 
 **showUserRules** | **optional.Bool**|  | 
 **tenantId** | **optional.Int32**| Tenant id to filter by. | 
 **tenantNameIcontains** | **optional.String**| Tenant name to filter by | 

### Return type

[**[]Quota**](Quota.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuotasPartialUpdate**
> QuotasPartialUpdate(ctx, id, optional)
Modify Quota

This endpoint modifies a quota.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Quota ID | 
 **optional** | ***QuotasApiQuotasPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuotasApiQuotasPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of QuotasIdBody**](QuotasIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuotasRead**
> Quota QuotasRead(ctx, id)
Return details of a quota.

This endpoint returns details of a specified quota.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Quota ID | 

### Return type

[**Quota**](Quota.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuotasRecalc**
> QuotasRecalc(ctx, optional)
Start recalculation flow all quotas

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QuotasApiQuotasRecalcOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuotasApiQuotasRecalcOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| Quota ID - if not given, will be applied to all quotas | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuotasRecalcStop**
> QuotasRecalcStop(ctx, )
Stop recalculation flow for all quotas

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuotasResetGracePeriod**
> QuotasResetGracePeriod(ctx, id)
Rest the grace period countdown after soft limit is exceeded

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Quota ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

