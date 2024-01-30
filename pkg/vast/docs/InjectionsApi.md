# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InjectionsCreate**](InjectionsApi.md#InjectionsCreate) | **Post** /injections/ | 
[**InjectionsDelete**](InjectionsApi.md#InjectionsDelete) | **Delete** /injections/ | Delete Injections
[**InjectionsGet**](InjectionsApi.md#InjectionsGet) | **Get** /injections/ | List injections

# **InjectionsCreate**
> InjectionsCreate(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InjectionsApiInjectionsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InjectionsApiInjectionsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InjectionsBody**](InjectionsBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InjectionsDelete**
> InjectionsDelete(ctx, )
Delete Injections

This endpoint deletes all injections.

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

# **InjectionsGet**
> []Injection InjectionsGet(ctx, optional)
List injections

This endpoint list injections with the option to filter by type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InjectionsApiInjectionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InjectionsApiInjectionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **injectionType** | **optional.String**| Injection type | 

### Return type

[**[]Injection**](Injection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

