# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LdapsCreate**](LdapsApi.md#LdapsCreate) | **Post** /ldaps/ | Create LDAP Configuration
[**LdapsDelete**](LdapsApi.md#LdapsDelete) | **Delete** /ldaps/{id}/ | Delete LDAP Configuration
[**LdapsList**](LdapsApi.md#LdapsList) | **Get** /ldaps/ | List LDAP Configurations
[**LdapsPartialUpdate**](LdapsApi.md#LdapsPartialUpdate) | **Patch** /ldaps/{id}/ | Modify LDAP configuration
[**LdapsRead**](LdapsApi.md#LdapsRead) | **Get** /ldaps/{id}/ | Return Details of an LDAP configuration.
[**LdapsSetPosixPrimary**](LdapsApi.md#LdapsSetPosixPrimary) | **Patch** /ldaps/{id}/set_posix_primary/ | Set LDAP as POSIX Primary provider.

# **LdapsCreate**
> Ldap LdapsCreate(ctx, optional)
Create LDAP Configuration

This endpoint creates an LDAP configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapsApiLdapsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapsApiLdapsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LdapsBody**](LdapsBody.md)|  | 

### Return type

[**Ldap**](Ldap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapsDelete**
> LdapsDelete(ctx, id)
Delete LDAP Configuration

This endpoint deletes an LDAP configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| LDAP ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapsList**
> []Ldap LdapsList(ctx, optional)
List LDAP Configurations

This endpoint lists LDAP configurations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapsApiLdapsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapsApiLdapsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 

### Return type

[**[]Ldap**](Ldap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapsPartialUpdate**
> LdapsPartialUpdate(ctx, id, optional)
Modify LDAP configuration

This endpoint modifies an LDAP configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| LDAP ID | 
 **optional** | ***LdapsApiLdapsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapsApiLdapsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of LdapsIdBody**](LdapsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapsRead**
> Ldap LdapsRead(ctx, id)
Return Details of an LDAP configuration.

This endpoint returns details of an LDAP configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| LDAP ID | 

### Return type

[**Ldap**](Ldap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapsSetPosixPrimary**
> LdapsSetPosixPrimary(ctx, id)
Set LDAP as POSIX Primary provider.

This endpoint sets the LDAP server as the primary POSIX provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| LDAP ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

