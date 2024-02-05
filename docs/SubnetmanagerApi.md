# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubnetmanagerCreate**](SubnetmanagerApi.md#SubnetmanagerCreate) | **Post** /subnetmanager/ | Add Subnet manager
[**SubnetmanagerDelete**](SubnetmanagerApi.md#SubnetmanagerDelete) | **Delete** /subnetmanager/{id}/ | Remove a Subnet Manager
[**SubnetmanagerList**](SubnetmanagerApi.md#SubnetmanagerList) | **Get** /subnetmanager/ | List subnet manager
[**SubnetmanagerPartialUpdate**](SubnetmanagerApi.md#SubnetmanagerPartialUpdate) | **Patch** /subnetmanager/{id}/ | Manage Subnet Manager
[**SubnetmanagerRead**](SubnetmanagerApi.md#SubnetmanagerRead) | **Get** /subnetmanager/{id}/ | Return Details of a Subnet manager

# **SubnetmanagerCreate**
> SubnetManager SubnetmanagerCreate(ctx, optional)
Add Subnet manager

This endpoint adds a subnet manager to VMS.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubnetmanagerApiSubnetmanagerCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubnetmanagerApiSubnetmanagerCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SubnetmanagerBody**](SubnetmanagerBody.md)|  | 

### Return type

[**SubnetManager**](SubnetManager.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubnetmanagerDelete**
> SubnetmanagerDelete(ctx, id)
Remove a Subnet Manager

This endpoint removes a subnet manager.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| License ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubnetmanagerList**
> []SubnetManager SubnetmanagerList(ctx, )
List subnet manager

This endpoint lists subnet manager.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SubnetManager**](SubnetManager.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubnetmanagerPartialUpdate**
> SubnetmanagerPartialUpdate(ctx, id, optional)
Manage Subnet Manager

This endpoint modifies some Subnet manager parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Subnet manager ID | 
 **optional** | ***SubnetmanagerApiSubnetmanagerPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubnetmanagerApiSubnetmanagerPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SubnetmanagerIdBody**](SubnetmanagerIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubnetmanagerRead**
> SubnetManager SubnetmanagerRead(ctx, id)
Return Details of a Subnet manager

This endpoint returns details of a specific subnet manager.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| subnet manager ID | 

### Return type

[**SubnetManager**](SubnetManager.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

