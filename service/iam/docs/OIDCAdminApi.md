# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVcloudOIDCRegisterClient1**](OIDCAdminApi.md#CreateVcloudOIDCRegisterClient1) | **Post** /api/v1/admin/oidc/vcloud/config-oidc | 
[**VcloudAutoConfigOIDCForOrg1**](OIDCAdminApi.md#VcloudAutoConfigOIDCForOrg1) | **Post** /api/v1/admin/oidc/vcloud/auto-config | 

# **CreateVcloudOIDCRegisterClient1**
> AutoConfigOidcvCloudClientResponse CreateVcloudOIDCRegisterClient1(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AutoConfigOidcvCloudClientRequest**](AutoConfigOidcvCloudClientRequest.md)|  | 

### Return type

[**AutoConfigOidcvCloudClientResponse**](AutoConfigOIDCVCloudClientResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VcloudAutoConfigOIDCForOrg1**
> VcloudOidcConfigurationResponse VcloudAutoConfigOIDCForOrg1(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VcloudOidcConfigurationRequest**](VcloudOidcConfigurationRequest.md)|  | 

### Return type

[**VcloudOidcConfigurationResponse**](VcloudOIDCConfigurationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

