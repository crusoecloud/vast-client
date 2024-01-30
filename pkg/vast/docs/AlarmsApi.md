# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlarmsClear**](AlarmsApi.md#AlarmsClear) | **Patch** /alarms/clear/ | Clear Alarms
[**AlarmsDelete**](AlarmsApi.md#AlarmsDelete) | **Delete** /alarms/{id}/ | Delete Alarm
[**AlarmsList**](AlarmsApi.md#AlarmsList) | **Get** /alarms/ | List Alarms
[**AlarmsPartialUpdate**](AlarmsApi.md#AlarmsPartialUpdate) | **Patch** /alarms/{id}/ | Acknowledge Alarm
[**AlarmsRead**](AlarmsApi.md#AlarmsRead) | **Get** /alarms/{id}/ | Return Details of an Alarm

# **AlarmsClear**
> AlarmsClear(ctx, )
Clear Alarms

This endpoint clears all alarms.

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

# **AlarmsDelete**
> AlarmsDelete(ctx, id)
Delete Alarm

This endpoint deletes an alarm.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Alarm ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlarmsList**
> []Alarm AlarmsList(ctx, optional)
List Alarms

This endpoint lists alarms. The parameters filter the list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlarmsApiAlarmsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlarmsApiAlarmsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **pageSize** | **optional.String**|  | 
 **timestampGte** | **optional.String**|  | 
 **timestampLte** | **optional.String**|  | 
 **severity** | **optional.String**|  | 
 **alarmMessageContains** | **optional.String**|  | 
 **objectGuid** | **optional.String**|  | 
 **objectType** | **optional.String**|  | 
 **objectId** | **optional.String**|  | 
 **onlyDiscoveredHosts** | **optional.String**| get only discovery related hosts | 

### Return type

[**[]Alarm**](Alarm.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlarmsPartialUpdate**
> AlarmsPartialUpdate(ctx, id, optional)
Acknowledge Alarm

This endpoint acknowledges an alarm.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Alarm ID | 
 **optional** | ***AlarmsApiAlarmsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlarmsApiAlarmsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AlarmsIdBody**](AlarmsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AlarmsRead**
> Alarm AlarmsRead(ctx, id)
Return Details of an Alarm

This endpoint returns details of a specified alarm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Alarm ID | 

### Return type

[**Alarm**](Alarm.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

