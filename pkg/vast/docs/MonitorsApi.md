# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MonitorsAdHocQuery**](MonitorsApi.md#MonitorsAdHocQuery) | **Get** /monitors/ad_hoc_query/ | Query Analytics with Ad Hoc Query Parameters
[**MonitorsCreate**](MonitorsApi.md#MonitorsCreate) | **Post** /monitors/ | Create Monitor
[**MonitorsDelete**](MonitorsApi.md#MonitorsDelete) | **Delete** /monitors/{id}/ | Delete Monitor (aka Analytics Report)
[**MonitorsList**](MonitorsApi.md#MonitorsList) | **Get** /monitors/ | List Monitors
[**MonitorsPartialUpdate**](MonitorsApi.md#MonitorsPartialUpdate) | **Patch** /monitors/{id}/ | Modify Analytics Report
[**MonitorsQuery**](MonitorsApi.md#MonitorsQuery) | **Get** /monitors/{id}/query/ | Query Analytics using Defined Analytics Report (aka Monitor)
[**MonitorsRead**](MonitorsApi.md#MonitorsRead) | **Get** /monitors/{id}/ | Return Details of an Analytics Report
[**MonitorsTopn**](MonitorsApi.md#MonitorsTopn) | **Get** /monitors/topn/ | Query Top Performer Data

# **MonitorsAdHocQuery**
> MonitorData MonitorsAdHocQuery(ctx, objectType, optional)
Query Analytics with Ad Hoc Query Parameters

This endpoint queries VMS for analytics.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectType** | **string**| The type of object to monitor. | 
 **optional** | ***MonitorsApiMonitorsAdHocQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorsApiMonitorsAdHocQueryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromTime** | **optional.String**| Start time of report. e.g. 2021-01-15T13:41:52Z | 
 **toTime** | **optional.String**| End time of report. e.g. 2021-01-15T13:41:52Z | 
 **timeFrame** | **optional.String**| Time frame to report. Specify as an integer followed by &#x27;m&#x27; (minutes) &#x27;h&#x27; (hours), &#x27;D&#x27; (days) ,&#x27;W&#x27; (weeks) or &#x27;M&#x27; (months) &#x27;Y&#x27; (years). e.g. 10m | 
 **objectIds** | [**optional.Interface of []int32**](int32.md)| Specify a list of IDs of objects of the specified object_type to filter the report on those specific objects. | 
 **propList** | [**optional.Interface of []string**](string.md)| A list of metrics to query. To get the full list of metrics, use GET /metrics/. | 
 **granularity** | **optional.String**| Data granularity: seconds (raw), minutes (five minute aggregated samples), hours (hourly aggregated samples), or days (daily aggregated samples) | 
 **aggregation** | **optional.String**| If granularity is minutes, hours or days, the data is aggregated. This parameter selects which aggregation function to use. | 
 **formatData** | **optional.Bool**|  | [default to false]

### Return type

[**MonitorData**](MonitorData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsCreate**
> Monitor MonitorsCreate(ctx, optional)
Create Monitor

This endpoint defines a custom analytics report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MonitorsApiMonitorsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorsApiMonitorsCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MonitorsBody**](MonitorsBody.md)|  | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsDelete**
> MonitorsDelete(ctx, id)
Delete Monitor (aka Analytics Report)

This endpoint deletes an analytics report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Monitor ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsList**
> []Monitor MonitorsList(ctx, )
List Monitors

This endpoint lists properties of predefined analytics reports.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Monitor**](Monitor.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsPartialUpdate**
> MonitorsPartialUpdate(ctx, id, optional)
Modify Analytics Report

This endpoint partially modifies the specified monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Monitor ID | 
 **optional** | ***MonitorsApiMonitorsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorsApiMonitorsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MonitorsIdBody**](MonitorsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsQuery**
> MonitorData MonitorsQuery(ctx, id, optional)
Query Analytics using Defined Analytics Report (aka Monitor)

This endpoint queries for analytics, based on a defined monitor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Monitor ID, to specify the analytics you want to query. | 
 **optional** | ***MonitorsApiMonitorsQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorsApiMonitorsQueryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromTime** | **optional.String**| Start time for the report. e.g. 2021-01-15T13:41:52Z | 
 **toTime** | **optional.String**| End time for the report. e.g. 2021-01-15T13:41:52Z | 
 **timeFrame** | **optional.String**| Time frame to report. Specify as an integer followed by &#x27;m&#x27; (minutes) &#x27;h&#x27; (hours), &#x27;D&#x27; (days) ,&#x27;W&#x27; (weeks) or &#x27;M&#x27; (months) &#x27;Y&#x27; (years). e.g. 10m | 
 **granularity** | **optional.String**| Data granularity: seconds (raw), minutes (five minute aggregated samples), hours (hourly aggregated samples), or days (daily aggregated samples) | 
 **aggregation** | **optional.String**| If granularity is minutes, hours or days, the data is aggregated. This parameter selects which aggregation function to use. | 
 **formatData** | **optional.Bool**|  | [default to false]

### Return type

[**MonitorData**](MonitorData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsRead**
> Monitor MonitorsRead(ctx, id)
Return Details of an Analytics Report

This endpoint returns the properties of a specified analytics report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Monitor ID | 

### Return type

[**Monitor**](Monitor.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MonitorsTopn**
> TopNData MonitorsTopn(ctx, key, type_, optional)
Query Top Performer Data

This endpoint queries VMS for top performer data. It can return the most active client users, NFS exports and client machines.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**|  | 
  **type_** | **string**|  | 
 **optional** | ***MonitorsApiMonitorsTopnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonitorsApiMonitorsTopnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **count** | **optional.Int32**|  | 
 **timestamp** | **optional.String**|  | 
 **userTitle** | **optional.String**|  | 
 **exportTitle** | **optional.String**|  | 
 **clientTitle** | **optional.String**|  | 

### Return type

[**TopNData**](TopNData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

