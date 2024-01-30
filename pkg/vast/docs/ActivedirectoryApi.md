# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActiveDirectoryCreate**](ActivedirectoryApi.md#ActiveDirectoryCreate) | **Post** /activedirectory/ | Create Active Directory Configuration
[**ActiveDirectoryCurrentGc**](ActivedirectoryApi.md#ActiveDirectoryCurrentGc) | **Get** /activedirectory/{id}/current_gc/ | Get current AD GC
[**ActiveDirectoryDcs**](ActivedirectoryApi.md#ActiveDirectoryDcs) | **Get** /activedirectory/{id}/dcs/ | Get AD DCs
[**ActiveDirectoryDelete**](ActivedirectoryApi.md#ActiveDirectoryDelete) | **Delete** /activedirectory/{id}/ | Delete Active Directory configuration
[**ActiveDirectoryDomains**](ActivedirectoryApi.md#ActiveDirectoryDomains) | **Get** /activedirectory/{id}/domains/ | Get AD domains
[**ActiveDirectoryGcs**](ActivedirectoryApi.md#ActiveDirectoryGcs) | **Get** /activedirectory/{id}/gcs/ | Get AD GCs
[**ActiveDirectoryIsOperationHealthy**](ActivedirectoryApi.md#ActiveDirectoryIsOperationHealthy) | **Post** /activedirectory/{id}/is_operation_healthy | Check whether an operation may be successfully performed
[**ActiveDirectoryList**](ActivedirectoryApi.md#ActiveDirectoryList) | **Get** /activedirectory/ | List Active Directory Configurations
[**ActiveDirectoryPartialUpdate**](ActivedirectoryApi.md#ActiveDirectoryPartialUpdate) | **Patch** /activedirectory/{id}/ | Join or Leave Active Directory
[**ActiveDirectoryRead**](ActivedirectoryApi.md#ActiveDirectoryRead) | **Get** /activedirectory/{id}/ | Return Active Directory Configuration Details
[**ActiveDirectoryRefresh**](ActivedirectoryApi.md#ActiveDirectoryRefresh) | **Patch** /activedirectory/{id}/refresh/ | Refresh AD

# **ActiveDirectoryCreate**
> ActiveDirectory ActiveDirectoryCreate(ctx, optional)
Create Active Directory Configuration

This endpoint creates an Active Directory configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ActivedirectoryApiActiveDirectoryCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActivedirectoryApiActiveDirectoryCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ActivedirectoryBody**](ActivedirectoryBody.md)|  | 

### Return type

[**ActiveDirectory**](ActiveDirectory.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActiveDirectoryCurrentGc**
> interface{} ActiveDirectoryCurrentGc(ctx, id)
Get current AD GC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| AD ID | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActiveDirectoryDcs**
> interface{} ActiveDirectoryDcs(ctx, id, optional)
Get AD DCs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| AD ID | 
 **optional** | ***ActivedirectoryApiActiveDirectoryDcsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActivedirectoryApiActiveDirectoryDcsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **optional.String**|  | 
 **isClosest** | **optional.String**|  | 
 **uri** | **optional.String**|  | 
 **port** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActiveDirectoryDelete**
> ActiveDirectoryDelete(ctx, id, optional)
Delete Active Directory configuration

This endpoint deletes an Active Directory configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Active Directory ID | 
 **optional** | ***ActivedirectoryApiActiveDirectoryDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActivedirectoryApiActiveDirectoryDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ActivedirectoryIdBody**](ActivedirectoryIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActiveDirectoryDomains**
> interface{} ActiveDirectoryDomains(ctx, id, optional)
Get AD domains

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| AD ID | 
 **optional** | ***ActivedirectoryApiActiveDirectoryDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActivedirectoryApiActiveDirectoryDomainsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseDn** | **optional.String**|  | 
 **fqdn** | **optional.String**|  | 
 **sid** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActiveDirectoryGcs**
> interface{} ActiveDirectoryGcs(ctx, id, optional)
Get AD GCs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| AD ID | 
 **optional** | ***ActivedirectoryApiActiveDirectoryGcsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActivedirectoryApiActiveDirectoryGcsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **optional.String**|  | 
 **isClosest** | **optional.String**|  | 
 **uri** | **optional.String**|  | 
 **port** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActiveDirectoryIsOperationHealthy**
> Warnings ActiveDirectoryIsOperationHealthy(ctx, id, optional)
Check whether an operation may be successfully performed

Returns warnings about the future state if operation is performed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| AD ID | 
 **optional** | ***ActivedirectoryApiActiveDirectoryIsOperationHealthyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActivedirectoryApiActiveDirectoryIsOperationHealthyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdIsOperationHealthyBody1**](IdIsOperationHealthyBody1.md)|  | 
 **operation** | **optional.**|  | 

### Return type

[**Warnings**](Warnings.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActiveDirectoryList**
> []ActiveDirectory ActiveDirectoryList(ctx, )
List Active Directory Configurations

This endpoint lists all Active Directory configurations.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ActiveDirectory**](ActiveDirectory.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActiveDirectoryPartialUpdate**
> AsyncTaskInResponse ActiveDirectoryPartialUpdate(ctx, id, optional)
Join or Leave Active Directory

This endpoint joins or leaves an Active Directory server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the Active Directory configuration for the AD server to join | 
 **optional** | ***ActivedirectoryApiActiveDirectoryPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActivedirectoryApiActiveDirectoryPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ActivedirectoryIdBody1**](ActivedirectoryIdBody1.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActiveDirectoryRead**
> ActiveDirectory ActiveDirectoryRead(ctx, id)
Return Active Directory Configuration Details

This endpoint returns details of a specified Active Directory configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Active Directory ID | 

### Return type

[**ActiveDirectory**](ActiveDirectory.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActiveDirectoryRefresh**
> ActiveDirectoryRefresh(ctx, id)
Refresh AD

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| AD ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

