# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVmsNetworkSettings**](VmsApi.md#GetVmsNetworkSettings) | **Get** /vms/{id}/network_settings/ | Return the network settings of VMS
[**ModifyVmsNetworkSettings**](VmsApi.md#ModifyVmsNetworkSettings) | **Patch** /vms/{id}/network_settings/ | Modify the network settings of VMS
[**ToggleMaintenanceMode**](VmsApi.md#ToggleMaintenanceMode) | **Patch** /vms/{id}/toggle_maintenance_mode/ | Toggle maintenance mode
[**VmsGetLoginBanner**](VmsApi.md#VmsGetLoginBanner) | **Get** /vms/{id}/login_banner/ | Returns login banner without auth.
[**VmsList**](VmsApi.md#VmsList) | **Get** /vms/ | List VMS
[**VmsNetworkSettingsSummary**](VmsApi.md#VmsNetworkSettingsSummary) | **Post** /vms/{id}/network_settings_summary/ | Preview summary of changes to the network settings of VMS
[**VmsPartialUpdate**](VmsApi.md#VmsPartialUpdate) | **Patch** /vms/{id}/ | Manage VMS
[**VmsRead**](VmsApi.md#VmsRead) | **Get** /vms/{id}/ | Return VMS Details
[**VmsRemoveClientCertificate**](VmsApi.md#VmsRemoveClientCertificate) | **Patch** /vms/{id}/remove_client_certificate/ | Remove VMS SSL Client Certificate
[**VmsResetCertificate**](VmsApi.md#VmsResetCertificate) | **Patch** /vms/{id}/reset_certificate/ | Reset VMS SSL Certificate to Default
[**VmsResetSslCiphers**](VmsApi.md#VmsResetSslCiphers) | **Patch** /vms/{id}/reset_ssl_ciphers/ | Reset VMS SSL Ciphers to Default
[**VmsSetCertificate**](VmsApi.md#VmsSetCertificate) | **Patch** /vms/{id}/set_certificate/ | Install SSL certificate and key on VMS
[**VmsSetClientCertificate**](VmsApi.md#VmsSetClientCertificate) | **Patch** /vms/{id}/set_client_certificate/ | Install SSL certificate and key on VMS
[**VmsSetSslCiphers**](VmsApi.md#VmsSetSslCiphers) | **Patch** /vms/{id}/set_ssl_ciphers/ | Set SSL ciphers use by VMS
[**VmsSetSslPort**](VmsApi.md#VmsSetSslPort) | **Patch** /vms/{id}/set_ssl_port/ | Change VMS SSL Port

# **GetVmsNetworkSettings**
> VmsNetworkSettings GetVmsNetworkSettings(ctx, id)
Return the network settings of VMS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**VmsNetworkSettings**](VMSNetworkSettings.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyVmsNetworkSettings**
> AsyncTaskInResponse ModifyVmsNetworkSettings(ctx, id, optional)
Modify the network settings of VMS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***VmsApiModifyVmsNetworkSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VmsApiModifyVmsNetworkSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of VmsNetworkSettingsInput**](VmsNetworkSettingsInput.md)|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ToggleMaintenanceMode**
> AsyncTaskInResponse ToggleMaintenanceMode(ctx, id)
Toggle maintenance mode

Change the system maintenance mode state

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**AsyncTaskInResponse**](AsyncTaskInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmsGetLoginBanner**
> Vms VmsGetLoginBanner(ctx, id)
Returns login banner without auth.

This endpoint returns login banner without auth.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Vms**](Vms.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmsList**
> []Vms VmsList(ctx, optional)
List VMS

This endpoint lists the Vast Management System (VMS).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VmsApiVmsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VmsApiVmsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.String**|  | 

### Return type

[**[]Vms**](Vms.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmsNetworkSettingsSummary**
> VmsNetworkSettingsSummary VmsNetworkSettingsSummary(ctx, id, optional)
Preview summary of changes to the network settings of VMS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***VmsApiVmsNetworkSettingsSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VmsApiVmsNetworkSettingsSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of VmsNetworkSettingsInput**](VmsNetworkSettingsInput.md)|  | 

### Return type

[**VmsNetworkSettingsSummary**](VMSNetworkSettingsSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmsPartialUpdate**
> VmsPartialUpdate(ctx, id, optional)
Manage VMS

This endpoint modifies some VMS parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| VMS ID | 
 **optional** | ***VmsApiVmsPartialUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VmsApiVmsPartialUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of VmsIdBody**](VmsIdBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmsRead**
> Vms VmsRead(ctx, id)
Return VMS Details

This endpoint details the Vast Management System (VMS).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Vms**](Vms.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmsRemoveClientCertificate**
> VmsRemoveClientCertificate(ctx, id)
Remove VMS SSL Client Certificate

This endpoint removes the SSL client certificate and client private key used by VMS for mTLS.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| VMS ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmsResetCertificate**
> VmsResetCertificate(ctx, id)
Reset VMS SSL Certificate to Default

This endpoint resets the SSL certificate and private key used by VMS to default.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| VMS ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmsResetSslCiphers**
> VmsResetSslCiphers(ctx, id)
Reset VMS SSL Ciphers to Default

This endpoint resets the SSL Ciphers used by VMS to default.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| VMS ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmsSetCertificate**
> VmsSetCertificate(ctx, id, optional)
Install SSL certificate and key on VMS

This endpoint installs a SSL certificate and private key for VMS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| VMS ID | 
 **optional** | ***VmsApiVmsSetCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VmsApiVmsSetCertificateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdSetCertificateBody**](IdSetCertificateBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmsSetClientCertificate**
> VmsSetClientCertificate(ctx, id, optional)
Install SSL certificate and key on VMS

This endpoint installs a SSL client certificate and client private key for mTLS VMS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| VMS ID | 
 **optional** | ***VmsApiVmsSetClientCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VmsApiVmsSetClientCertificateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdSetClientCertificateBody**](IdSetClientCertificateBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmsSetSslCiphers**
> VmsSetSslCiphers(ctx, id, optional)
Set SSL ciphers use by VMS

This endpoint set SSL ciphers used by VMS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| VMS ID | 
 **optional** | ***VmsApiVmsSetSslCiphersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VmsApiVmsSetSslCiphersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdSetSslCiphersBody**](IdSetSslCiphersBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VmsSetSslPort**
> VmsSetSslPort(ctx, id, optional)
Change VMS SSL Port

This endpoint changes the SSL port used for accessing VMS. The default SSL port is 443.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| VMS ID | 
 **optional** | ***VmsApiVmsSetSslPortOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VmsApiVmsSetSslPortOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of IdSetSslPortBody**](IdSetSslPortBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

