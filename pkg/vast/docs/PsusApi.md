# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PsusList**](PsusApi.md#PsusList) | **Get** /psus/ | List PSUs
[**PsusRead**](PsusApi.md#PsusRead) | **Get** /psus/{id}/ | Return Details of One PSU

# **PsusList**
> []Psu PsusList(ctx, optional)
List PSUs

This endpoint lists PSUs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PsusApiPsusListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PsusApiPsusListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **state** | **optional.String**| Filter by PSU state | 
 **sn** | **optional.String**| Filter by PSU serial number | 
 **model** | **optional.String**| Filter by PSU model | 
 **boxName** | **optional.String**| Filter by parent Box name | 
 **boxId** | **optional.String**| Filter by parent Box ID | 

### Return type

[**[]Psu**](PSU.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PsusRead**
> Psu PsusRead(ctx, id)
Return Details of One PSU

This endpoint returns details of a specified PSU.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| PSU ID | 

### Return type

[**Psu**](PSU.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

