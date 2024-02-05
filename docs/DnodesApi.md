# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ControlLed**](DnodesApi.md#ControlLed) | **Patch** /dnodes/{id}/control_led/ | Control DNode LED
[**DnodesList**](DnodesApi.md#DnodesList) | **Get** /dnodes/ | List DNodes
[**DnodesPartialUpdate**](DnodesApi.md#DnodesPartialUpdate) | **Patch** /dnodes/{id}/ | Activate/Deactivate/Replace/Power On/Off DNode
[**DnodesRead**](DnodesApi.md#DnodesRead) | **Get** /dnodes/{id}/ | Return Details of a DNode
[**Highlight**](DnodesApi.md#Highlight) | **Patch** /dnodes/{id}/highlight/ | Highlight DNode
[**Rename**](DnodesApi.md#Rename) | **Patch** /dnodes/{id}/rename/ | Rename DNode

# **ControlLed**
> ControlLed(ctx, id, optional)
Control DNode LED

This endpoint controls a DNode LED

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DNode ID | 
 **optional** | ***DnodesApiControlLedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DnodesApiControlLedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdControlLedBody1**](IdControlLedBody1.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DnodesList**
> []DNode DnodesList(ctx, optional)
List DNodes

This endpoint returns a list of Dnodes in the cluster, with optional filtering.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DnodesApiDnodesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DnodesApiDnodesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **ip** | **optional.String**| Filter by DNode IP | 
 **state** | **optional.String**| Filter by DNode state | 
 **clusterName** | **optional.String**|  | 
 **clusterId** | **optional.Int32**|  | 
 **name** | **optional.String**| Filter by DNode name | 
 **enabled** | **optional.Bool**| List only enabled DNodes | 

### Return type

[**[]DNode**](DNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DnodesPartialUpdate**
> AsyncTaskInResponse DnodesPartialUpdate(ctx, id, optional)
Activate/Deactivate/Replace/Power On/Off DNode

This endpoint activates, deactivates, replaces, powers off and powers on a DNode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DNode ID | 
 **optional** | ***DnodesApiDnodesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DnodesApiDnodesPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DnodesIdBody**](DnodesIdBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DnodesRead**
> DNode DnodesRead(ctx, id)
Return Details of a DNode

This endpoint returns details of a DNode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DNode ID | 

### Return type

[**DNode**](DNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Highlight**
> AsyncTaskInResponse Highlight(ctx, id)
Highlight DNode

This endpoint highlights a DNode

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DNode ID | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Rename**
> Rename(ctx, id, optional)
Rename DNode

This endpoint renames a DNode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| DNode ID | 
 **optional** | ***DnodesApiRenameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DnodesApiRenameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdRenameBody1**](IdRenameBody1.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

