# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GlobalpoliciesCreate**](GlobalpoliciesApi.md#GlobalpoliciesCreate) | **Post** /globalpolicies/ | Create Exposed Path Policy
[**GlobalpoliciesDelete**](GlobalpoliciesApi.md#GlobalpoliciesDelete) | **Delete** /globalpolicies/{id}/ | Delete Exposed Path Policy
[**GlobalpoliciesList**](GlobalpoliciesApi.md#GlobalpoliciesList) | **Get** /globalpolicies/ | Return Exposed Path Policies
[**GlobalpoliciesPartialUpdate**](GlobalpoliciesApi.md#GlobalpoliciesPartialUpdate) | **Patch** /globalpolicies/{id}/ | Modify an Exposed Path Policy
[**GlobalpoliciesRead**](GlobalpoliciesApi.md#GlobalpoliciesRead) | **Get** /globalpolicies/{id}/ | Return Details of an Exposed Path Policy

# **GlobalpoliciesCreate**
> GlobalPolicy GlobalpoliciesCreate(ctx, optional)
Create Exposed Path Policy

This endpoint creates an exposed path policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GlobalpoliciesApiGlobalpoliciesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalpoliciesApiGlobalpoliciesCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GlobalpoliciesBody**](GlobalpoliciesBody.md)|  | 

### Return type

[**GlobalPolicy**](GlobalPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalpoliciesDelete**
> GlobalpoliciesDelete(ctx, id)
Delete Exposed Path Policy

This endpoint deletes an exposed path policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Exposed Path Policy ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalpoliciesList**
> []GlobalPolicy GlobalpoliciesList(ctx, optional)
Return Exposed Path Policies

This endpoint lists Exposed Path Policies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GlobalpoliciesApiGlobalpoliciesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalpoliciesApiGlobalpoliciesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **name** | **optional.String**| Filter by exposed path policy name | 

### Return type

[**[]GlobalPolicy**](GlobalPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalpoliciesPartialUpdate**
> GlobalpoliciesPartialUpdate(ctx, id, optional)
Modify an Exposed Path Policy

This endpoint modifies an exposed path policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Exposed Path Policy ID | 
 **optional** | ***GlobalpoliciesApiGlobalpoliciesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GlobalpoliciesApiGlobalpoliciesPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GlobalpoliciesIdBody**](GlobalpoliciesIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GlobalpoliciesRead**
> GlobalPolicy GlobalpoliciesRead(ctx, id)
Return Details of an Exposed Path Policy

This endpoint returns details of a specified exposed path policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Exposed Path Policy ID | 

### Return type

[**GlobalPolicy**](GlobalPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

