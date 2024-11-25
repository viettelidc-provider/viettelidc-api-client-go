# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVcloudOIDCRegisterClient**](OIDCClientApi.md#CreateVcloudOIDCRegisterClient) | **Post** /api/v1/oidc/vcloud | 
[**DeleteVcloudOIDCRegisterClient**](OIDCClientApi.md#DeleteVcloudOIDCRegisterClient) | **Delete** /api/v1/oidc/vcloud | 
[**GetDetailVcloudOIDCRegisterClient**](OIDCClientApi.md#GetDetailVcloudOIDCRegisterClient) | **Get** /api/v1/oidc/vcloud | 
[**RegenerateVcloudClientSecret**](OIDCClientApi.md#RegenerateVcloudClientSecret) | **Put** /api/v1/oidc/vcloud/regenerate | 
[**UpdateVcloudOIDCRegisterClient**](OIDCClientApi.md#UpdateVcloudOIDCRegisterClient) | **Put** /api/v1/oidc/vcloud | 

# **CreateVcloudOIDCRegisterClient**
> CreateVcloudOidcRegisterClientResponse CreateVcloudOIDCRegisterClient(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVcloudOidcRegisterClientRequest**](CreateVcloudOidcRegisterClientRequest.md)|  | 

### Return type

[**CreateVcloudOidcRegisterClientResponse**](CreateVcloudOIDCRegisterClientResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVcloudOIDCRegisterClient**
> DeleteVcloudOidcRegisterClientResponse DeleteVcloudOIDCRegisterClient(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DeleteVcloudOidcRegisterClientResponse**](DeleteVcloudOIDCRegisterClientResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailVcloudOIDCRegisterClient**
> GetDetailVcloudOidcRegisterClientResponse GetDetailVcloudOIDCRegisterClient(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetDetailVcloudOidcRegisterClientResponse**](GetDetailVcloudOIDCRegisterClientResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateVcloudClientSecret**
> RegenerateVcloudClientSecretClientResponse RegenerateVcloudClientSecret(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RegenerateVcloudClientSecretClientResponse**](RegenerateVcloudClientSecretClientResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVcloudOIDCRegisterClient**
> UpdateVcloudOidcRegisterClientResponse UpdateVcloudOIDCRegisterClient(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVcloudOidcRegisterClientRequest**](UpdateVcloudOidcRegisterClientRequest.md)|  | 

### Return type

[**UpdateVcloudOidcRegisterClientResponse**](UpdateVcloudOIDCRegisterClientResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

