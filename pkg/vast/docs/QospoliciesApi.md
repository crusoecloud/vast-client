# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QospoliciesCreate**](QospoliciesApi.md#QospoliciesCreate) | **Post** /qospolicies/ | Create QoS Policy
[**QospoliciesDelete**](QospoliciesApi.md#QospoliciesDelete) | **Delete** /qospolicies/{id}/ | Delete QoS Policy
[**QospoliciesList**](QospoliciesApi.md#QospoliciesList) | **Get** /qospolicies/ | List QoS Policies
[**QospoliciesPartialUpdate**](QospoliciesApi.md#QospoliciesPartialUpdate) | **Patch** /qospolicies/{id}/ | Modify QoS Policy
[**QospoliciesRead**](QospoliciesApi.md#QospoliciesRead) | **Get** /qospolicies/{id}/ | Return Details of QoS Policy

# **QospoliciesCreate**
> QosPolicy QospoliciesCreate(ctx, optional)
Create QoS Policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QospoliciesApiQospoliciesCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QospoliciesApiQospoliciesCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of QospoliciesBody**](QospoliciesBody.md)|  | 

### Return type

[**QosPolicy**](QOSPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QospoliciesDelete**
> QospoliciesDelete(ctx, id)
Delete QoS Policy

Deletes a specified QoS Policy.

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
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QospoliciesList**
> []QosPolicy QospoliciesList(ctx, optional)
List QoS Policies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QospoliciesApiQospoliciesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QospoliciesApiQospoliciesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**[]QosPolicy**](QOSPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QospoliciesPartialUpdate**
> QosPolicy QospoliciesPartialUpdate(ctx, id, optional)
Modify QoS Policy

Modifies a specified QoS Policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***QospoliciesApiQospoliciesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QospoliciesApiQospoliciesPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of QospoliciesIdBody**](QospoliciesIdBody.md)|  | 

### Return type

[**QosPolicy**](QOSPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QospoliciesRead**
> QosPolicy QospoliciesRead(ctx, id)
Return Details of QoS Policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**QosPolicy**](QOSPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

