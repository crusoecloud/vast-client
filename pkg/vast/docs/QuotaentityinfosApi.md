# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuotaEntityInfosList**](QuotaentityinfosApi.md#QuotaEntityInfosList) | **Get** /quotaentityinfos/ | List Quota Entities

# **QuotaEntityInfosList**
> []QuotaEntityInfo QuotaEntityInfosList(ctx, optional)
List Quota Entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QuotaentityinfosApiQuotaEntityInfosListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QuotaentityinfosApiQuotaEntityInfosListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 
 **pageSize** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **vastId** | **optional.Int32**|  | 

### Return type

[**[]QuotaEntityInfo**](QuotaEntityInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

