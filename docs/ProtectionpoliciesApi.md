# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProtectionpoliciesCreate**](ProtectionpoliciesApi.md#ProtectionpoliciesCreate) | **Post** /protectionpolicies/ | Create Protection Policy
[**ProtectionpoliciesDelete**](ProtectionpoliciesApi.md#ProtectionpoliciesDelete) | **Delete** /protectionpolicies/{id}/ | Delete Protection Policy
[**ProtectionpoliciesList**](ProtectionpoliciesApi.md#ProtectionpoliciesList) | **Get** /protectionpolicies/ | Return Protection Policies
[**ProtectionpoliciesPartialUpdate**](ProtectionpoliciesApi.md#ProtectionpoliciesPartialUpdate) | **Patch** /protectionpolicies/{id}/ | Modify a Protection Policy
[**ProtectionpoliciesRead**](ProtectionpoliciesApi.md#ProtectionpoliciesRead) | **Get** /protectionpolicies/{id}/ | Return Details of a Protection Policy

# **ProtectionpoliciesCreate**
> ProtectionPolicy ProtectionpoliciesCreate(ctx, optional)
Create Protection Policy

This endpoint creates a protection policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionpoliciesApiProtectionpoliciesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionpoliciesApiProtectionpoliciesCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ProtectionpoliciesBody**](ProtectionpoliciesBody.md)|  | 

### Return type

[**ProtectionPolicy**](ProtectionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectionpoliciesDelete**
> ProtectionpoliciesDelete(ctx, id)
Delete Protection Policy

This endpoint deletes a protection policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Protection Policy ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectionpoliciesList**
> []ProtectionPolicy ProtectionpoliciesList(ctx, optional)
Return Protection Policies

This endpoint lists protection policies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProtectionpoliciesApiProtectionpoliciesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionpoliciesApiProtectionpoliciesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **targetName** | **optional.String**| Filter by name of replication peer | 
 **name** | **optional.String**| Filter by protection policy name | 

### Return type

[**[]ProtectionPolicy**](ProtectionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectionpoliciesPartialUpdate**
> ProtectionpoliciesPartialUpdate(ctx, id, optional)
Modify a Protection Policy

This endpoint modifies a protection policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Protection Policy ID | 
 **optional** | ***ProtectionpoliciesApiProtectionpoliciesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProtectionpoliciesApiProtectionpoliciesPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProtectionpoliciesIdBody**](ProtectionpoliciesIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProtectionpoliciesRead**
> ProtectionPolicy ProtectionpoliciesRead(ctx, id)
Return Details of a Protection Policy

This endpoint returns details of a specified protection policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Protection Policy ID | 

### Return type

[**ProtectionPolicy**](ProtectionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

