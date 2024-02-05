# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CapacityQuery**](CapacityApi.md#CapacityQuery) | **Get** /capacity/ | 

# **CapacityQuery**
> CapacityData CapacityQuery(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CapacityApiCapacityQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CapacityApiCapacityQueryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **tenantId** | **optional.Int32**|  | 

### Return type

[**CapacityData**](CapacityData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

