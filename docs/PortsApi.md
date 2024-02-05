# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortsList**](PortsApi.md#PortsList) | **Get** /ports/ | List Switch Ports
[**PortsRead**](PortsApi.md#PortsRead) | **Get** /ports/{id}/ | Return Details of One Switch Port

# **PortsList**
> []Port PortsList(ctx, optional)
List Switch Ports

This endpoint lists switch ports.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PortsApiPortsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PortsApiPortsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **optional.String**| Filter by port state | 
 **model** | **optional.String**| Filter by port model | 

### Return type

[**[]Port**](Port.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortsRead**
> Port PortsRead(ctx, id)
Return Details of One Switch Port

This endpoint returns details of a specific switch port.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Port ID | 

### Return type

[**Port**](Port.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

