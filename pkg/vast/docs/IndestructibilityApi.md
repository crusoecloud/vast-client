# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IndestructibilityGenerateToken**](IndestructibilityApi.md#IndestructibilityGenerateToken) | **Patch** /indestructibility/{id}/generate_token/ | Generate indestructibility token for validate (valid 1 hour)
[**IndestructibilityList**](IndestructibilityApi.md#IndestructibilityList) | **Get** /indestructibility/ | List Indestructibility
[**IndestructibilityRead**](IndestructibilityApi.md#IndestructibilityRead) | **Get** /indestructibility/{id}/ | Return Details of a Indestructibility.
[**IndestructibilityResetPasswd**](IndestructibilityApi.md#IndestructibilityResetPasswd) | **Patch** /indestructibility/{id}/reset_passwd/ | Reset Indestructibility password
[**IndestructibilityUnlock**](IndestructibilityApi.md#IndestructibilityUnlock) | **Patch** /indestructibility/{id}/unlock/ | Make system destructible for 1 hour
[**IndestructibilityUpdate**](IndestructibilityApi.md#IndestructibilityUpdate) | **Patch** /indestructibility/{id}/ | Modify Indestructibility.

# **IndestructibilityGenerateToken**
> Indestructibility IndestructibilityGenerateToken(ctx, id, optional)
Generate indestructibility token for validate (valid 1 hour)

This endpoint generate Indestructibility token for 1 hour.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***IndestructibilityApiIndestructibilityGenerateTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndestructibilityApiIndestructibilityGenerateTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdGenerateTokenBody**](IdGenerateTokenBody.md)|  | 

### Return type

[**Indestructibility**](Indestructibility.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndestructibilityList**
> []Indestructibility IndestructibilityList(ctx, optional)
List Indestructibility

This endpoint lists the VMS manager users with the option to filter by user name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IndestructibilityApiIndestructibilityListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndestructibilityApiIndestructibilityListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**|  | 

### Return type

[**[]Indestructibility**](Indestructibility.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndestructibilityRead**
> Indestructibility IndestructibilityRead(ctx, id)
Return Details of a Indestructibility.

This endpoint returns details of an Indestructibility.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Indestructibility ID | 

### Return type

[**Indestructibility**](Indestructibility.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndestructibilityResetPasswd**
> Indestructibility IndestructibilityResetPasswd(ctx, id)
Reset Indestructibility password

This endpoint reset Indestructibility password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Indestructibility**](Indestructibility.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndestructibilityUnlock**
> Indestructibility IndestructibilityUnlock(ctx, id, optional)
Make system destructible for 1 hour

This endpoint make system destructible for 1 hour.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***IndestructibilityApiIndestructibilityUnlockOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndestructibilityApiIndestructibilityUnlockOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdUnlockBody**](IdUnlockBody.md)|  | 

### Return type

[**Indestructibility**](Indestructibility.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IndestructibilityUpdate**
> Indestructibility IndestructibilityUpdate(ctx, id, optional)
Modify Indestructibility.

This endpoint updates indestructibility password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***IndestructibilityApiIndestructibilityUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IndestructibilityApiIndestructibilityUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IndestructibilityIdBody**](IndestructibilityIdBody.md)|  | 

### Return type

[**Indestructibility**](Indestructibility.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

