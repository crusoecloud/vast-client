# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CnodesAdd**](CnodesApi.md#CnodesAdd) | **Post** /cnodes/ | Add CNode
[**CnodesList**](CnodesApi.md#CnodesList) | **Get** /cnodes/ | List Cnodes
[**CnodesPartialUpdate**](CnodesApi.md#CnodesPartialUpdate) | **Patch** /cnodes/{id}/ | Activate/Deactivate/Replace/Power On/Off CNode
[**CnodesRead**](CnodesApi.md#CnodesRead) | **Get** /cnodes/{id}/ | Return Details of a CNode
[**CnodesRemove**](CnodesApi.md#CnodesRemove) | **Delete** /cnodes/{id}/ | Remove CNode
[**ControlLed**](CnodesApi.md#ControlLed) | **Patch** /cnodes/{id}/control_led/ | Control CNode LED
[**Highlight**](CnodesApi.md#Highlight) | **Patch** /cnodes/{id}/highlight/ | HighLight CNode
[**Rename**](CnodesApi.md#Rename) | **Patch** /cnodes/{id}/rename/ | Rename CNode

# **CnodesAdd**
> AsyncCNode CnodesAdd(ctx, optional)
Add CNode

This endpoint adds a CNode to the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CnodesApiCnodesAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CnodesApiCnodesAddOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CnodesBody**](CnodesBody.md)|  | 

### Return type

[**AsyncCNode**](AsyncCNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CnodesList**
> []CNode CnodesList(ctx, optional)
List Cnodes

Get the list of cnodes already assigned to cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CnodesApiCnodesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CnodesApiCnodesListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **ip** | **optional.String**| Filter by CNode IP | 
 **state** | **optional.String**| Filter by state | 
 **clusterName** | **optional.String**|  | 
 **clusterId** | **optional.Int32**|  | 
 **name** | **optional.String**| Filter by CNode name | 
 **enabled** | **optional.Bool**| Return only enabled CNodes | 
 **vippoolId** | **optional.Int32**|  | 

### Return type

[**[]CNode**](CNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CnodesPartialUpdate**
> AsyncTaskInResponse CnodesPartialUpdate(ctx, id, optional)
Activate/Deactivate/Replace/Power On/Off CNode

This endpoint activates, deactivates, replaces, powers off and powers on a CNode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| CNode ID | 
 **optional** | ***CnodesApiCnodesPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CnodesApiCnodesPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CnodesIdBody**](CnodesIdBody.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CnodesRead**
> CNode CnodesRead(ctx, id)
Return Details of a CNode

This endpoint returns details of a CNode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| CNode ID | 

### Return type

[**CNode**](CNode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CnodesRemove**
> AsyncTaskInResponse CnodesRemove(ctx, id, optional)
Remove CNode

This endpoint removes a CNode from the cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| CNode ID | 
 **optional** | ***CnodesApiCnodesRemoveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CnodesApiCnodesRemoveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force CNode removal | [default to false]

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ControlLed**
> ControlLed(ctx, id, optional)
Control CNode LED

This endpoint controls the CNode LED

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| CNode ID | 
 **optional** | ***CnodesApiControlLedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CnodesApiControlLedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdControlLedBody**](IdControlLedBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Highlight**
> AsyncTaskInResponse Highlight(ctx, id)
HighLight CNode

This endpoint highlights a CNode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| CNode ID | 

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
Rename CNode

This endpoint renames a CNode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| CNode ID | 
 **optional** | ***CnodesApiRenameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CnodesApiRenameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdRenameBody**](IdRenameBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

