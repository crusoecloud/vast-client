# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ObjectpermissionsList**](ObjectpermissionsApi.md#ObjectpermissionsList) | **Get** /permissions/objects/ | List Object Permissions

# **ObjectpermissionsList**
> []ObjectPermission ObjectpermissionsList(ctx, optional)
List Object Permissions

This endpoint lists object permissions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ObjectpermissionsApiObjectpermissionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ObjectpermissionsApiObjectpermissionsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectType** | **optional.String**|  | 
 **objectId** | **optional.String**|  | 
 **realm** | **optional.String**|  | 
 **permissions** | **optional.String**|  | 
 **managerId** | **optional.String**|  | 
 **roleId** | **optional.String**|  | 

### Return type

[**[]ObjectPermission**](ObjectPermission.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

