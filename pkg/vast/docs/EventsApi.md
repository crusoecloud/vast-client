# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsCreate**](EventsApi.md#EventsCreate) | **Post** /events/ | Create Event
[**EventsList**](EventsApi.md#EventsList) | **Get** /events/ | List Events
[**EventsRead**](EventsApi.md#EventsRead) | **Get** /events/{id}/ | Return Details of an Event

# **EventsCreate**
> Event EventsCreate(ctx, optional)
Create Event

This endpoint creates an event.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventsApiEventsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiEventsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EventsBody**](EventsBody.md)|  | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventsList**
> []Event EventsList(ctx, optional)
List Events

This endpoint lists events. The parameters filter the list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventsApiEventsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiEventsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **pageSize** | **optional.String**|  | 
 **timestampGte** | **optional.String**|  | 
 **timestampLte** | **optional.String**|  | 
 **eventType** | **optional.String**|  | 
 **eventTypeIn** | **optional.String**|  | 
 **eventOrigin** | **optional.String**|  | 
 **eventOriginIn** | **optional.String**|  | 
 **eventMessageContains** | **optional.String**|  | 
 **objectGuid** | **optional.String**|  | 
 **objectType** | **optional.String**|  | 
 **objectTypeIn** | **optional.String**|  | 
 **objectId** | **optional.String**|  | 
 **severity** | **optional.String**|  | 
 **severityIn** | **optional.String**|  | 

### Return type

[**[]Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventsRead**
> Event EventsRead(ctx, id)
Return Details of an Event

This endpoint returns details of a specified event.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

