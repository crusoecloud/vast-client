# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventdefinitionsList**](EventdefinitionsApi.md#EventdefinitionsList) | **Get** /eventdefinitions/ | List Event Definitions
[**EventdefinitionsPartialUpdate**](EventdefinitionsApi.md#EventdefinitionsPartialUpdate) | **Patch** /eventdefinitions/{id}/ | Modify Event Definition
[**EventdefinitionsRead**](EventdefinitionsApi.md#EventdefinitionsRead) | **Get** /eventdefinitions/{id}/ | Return Details of an Event Definition
[**EventdefinitionsTest**](EventdefinitionsApi.md#EventdefinitionsTest) | **Patch** /eventdefinitions/{id}/test/ | Test Event Definition

# **EventdefinitionsList**
> []EventDefinition EventdefinitionsList(ctx, optional)
List Event Definitions

This endpoint lists event definitions. The parameters filter the list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventdefinitionsApiEventdefinitionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventdefinitionsApiEventdefinitionsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectType** | **optional.String**|  | 
 **eventType** | **optional.String**|  | 
 **metadataProperty** | **optional.String**|  | 

### Return type

[**[]EventDefinition**](EventDefinition.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventdefinitionsPartialUpdate**
> EventdefinitionsPartialUpdate(ctx, id, optional)
Modify Event Definition

This endpoint modifies an event definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Event Definition ID | 
 **optional** | ***EventdefinitionsApiEventdefinitionsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventdefinitionsApiEventdefinitionsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EventdefinitionsIdBody**](EventdefinitionsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventdefinitionsRead**
> EventDefinition EventdefinitionsRead(ctx, id)
Return Details of an Event Definition

This endpoint returns details of a specified event definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Event Definition ID | 

### Return type

[**EventDefinition**](EventDefinition.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventdefinitionsTest**
> AsyncTaskInResponse EventdefinitionsTest(ctx, id)
Test Event Definition

This endpoint tests an event definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Event Definition ID | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

