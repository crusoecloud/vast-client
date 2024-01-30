# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserQuotasCreate**](UserquotasApi.md#UserQuotasCreate) | **Post** /userquotas/ | Create a User Quota
[**UserQuotasDelete**](UserquotasApi.md#UserQuotasDelete) | **Delete** /userquotas/{id}/ | Delete User Quota
[**UserQuotasList**](UserquotasApi.md#UserQuotasList) | **Get** /userquotas/ | List User Quotas
[**UserQuotasPartialUpdate**](UserquotasApi.md#UserQuotasPartialUpdate) | **Patch** /userquotas/{id}/ | Modify User Quota
[**UserQuotasRead**](UserquotasApi.md#UserQuotasRead) | **Get** /userquotas/{id}/ | Return details of a user quota.

# **UserQuotasCreate**
> UserQuota UserQuotasCreate(ctx, optional)
Create a User Quota

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserquotasApiUserQuotasCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserquotasApiUserQuotasCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UserquotasBody**](UserquotasBody.md)|  | 

### Return type

[**UserQuota**](UserQuota.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserQuotasDelete**
> UserQuotasDelete(ctx, id)
Delete User Quota

This endpoint deletes a user quota.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| User Quota ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserQuotasList**
> []UserQuota UserQuotasList(ctx, optional)
List User Quotas

This endpoint lists user quotas. The parameters enable you to filter the list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserquotasApiUserQuotasListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserquotasApiUserQuotasListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **pageSize** | **optional.String**|  | 
 **isAccountable** | **optional.Bool**|  | 
 **quotaSystemId** | **optional.Int32**|  | 
 **quotaId** | **optional.Int32**|  | 
 **entityIsGroup** | **optional.Bool**|  | 
 **refreshData** | **optional.Bool**|  | 

### Return type

[**[]UserQuota**](UserQuota.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserQuotasPartialUpdate**
> UserQuotasPartialUpdate(ctx, id, optional)
Modify User Quota

This endpoint modifies a user quota.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| User Quota ID | 
 **optional** | ***UserquotasApiUserQuotasPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserquotasApiUserQuotasPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UserquotasIdBody**](UserquotasIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserQuotasRead**
> UserQuota UserQuotasRead(ctx, id)
Return details of a user quota.

This endpoint returns details of a specified user quota.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| User Quota ID | 

### Return type

[**UserQuota**](UserQuota.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

