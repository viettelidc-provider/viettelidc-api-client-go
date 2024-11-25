# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VcloudAutoConfigOIDCForOrg**](OIDCAdminATMApi.md#VcloudAutoConfigOIDCForOrg) | **Post** /api/v1/atm-admin/oidc/vcloud/auto-config | 

# **VcloudAutoConfigOIDCForOrg**
> VcloudOidcConfigurationResponse VcloudAutoConfigOIDCForOrg(ctx, body)


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

