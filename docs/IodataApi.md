# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IodataQuery**](IodataApi.md#IodataQuery) | **Get** /iodata/ | List IO Data (not supported)

# **IodataQuery**
> IoData IodataQuery(ctx, optional)
List IO Data (not supported)

This endpoint is not supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IodataApiIodataQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IodataApiIodataQueryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graph** | **optional.Bool**|  | 
 **protocolFilter** | **optional.String**|  | 
 **timeFrame** | **optional.String**|  | 
 **startTime** | **optional.String**|  | 
 **endTime** | **optional.String**|  | 

### Return type

[**IoData**](IOData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

