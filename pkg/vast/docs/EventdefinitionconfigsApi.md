# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventdefinitionconfigsList**](EventdefinitionconfigsApi.md#EventdefinitionconfigsList) | **Get** /eventdefinitionconfigs/ | List Event Definition Configurations
[**EventdefinitionconfigsPartialUpdate**](EventdefinitionconfigsApi.md#EventdefinitionconfigsPartialUpdate) | **Patch** /eventdefinitionconfigs/{id}/ | Modify Event Definition Configuration
[**EventdefinitionconfigsRead**](EventdefinitionconfigsApi.md#EventdefinitionconfigsRead) | **Get** /eventdefinitionconfigs/{id}/ | Return Details of Event Definition Configuration

# **EventdefinitionconfigsList**
> []EventDefinitionConfig EventdefinitionconfigsList(ctx, )
List Event Definition Configurations

This endpoint lists Event Definition configurations

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EventDefinitionConfig**](EventDefinitionConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventdefinitionconfigsPartialUpdate**
> EventdefinitionconfigsPartialUpdate(ctx, id, optional)
Modify Event Definition Configuration

This endpoint modifies the global event definition configuration which includes the default actions triggered by alarms.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Event Definition Configuration ID | 
 **optional** | ***EventdefinitionconfigsApiEventdefinitionconfigsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventdefinitionconfigsApiEventdefinitionconfigsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EventdefinitionconfigsIdBody**](EventdefinitionconfigsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventdefinitionconfigsRead**
> EventDefinitionConfig EventdefinitionconfigsRead(ctx, id)
Return Details of Event Definition Configuration

This endpoint returns the details of an Event Definition configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Event Definition ID | 

### Return type

[**EventDefinitionConfig**](EventDefinitionConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

