# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authorize**](AuthorizationControllerApi.md#Authorize) | **Post** /api/v1/authorization | 
[**ChangePassword**](AuthorizationControllerApi.md#ChangePassword) | **Post** /api/v1/authorization/login/change-password | 
[**ChangePasswordByToken**](AuthorizationControllerApi.md#ChangePasswordByToken) | **Post** /api/v1/authorization/login/change-password-by-token | 
[**GetAllowResourceIds**](AuthorizationControllerApi.md#GetAllowResourceIds) | **Get** /api/v1/authorization | 
[**LoginViaATM**](AuthorizationControllerApi.md#LoginViaATM) | **Post** /api/v1/authorization/via-atm-login | 
[**LoginViaLoginPage**](AuthorizationControllerApi.md#LoginViaLoginPage) | **Post** /api/v1/authorization/login | 
[**VerifyMfaTokenCode**](AuthorizationControllerApi.md#VerifyMfaTokenCode) | **Post** /api/v1/authorization/login/verify-mfa-token | 

# **Authorize**
> CheckUserHasPermissionToExecuteResponse Authorize(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CheckUserHasPermissionToExecuteRequest**](CheckUserHasPermissionToExecuteRequest.md)|  | 

### Return type

[**CheckUserHasPermissionToExecuteResponse**](CheckUserHasPermissionToExecuteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangePassword**
> ChangePassword(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthorizationControllerApiChangePasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthorizationControllerApiChangePasswordOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RequiredChangePasswordRequest**](RequiredChangePasswordRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangePasswordByToken**
> RequiredChangePasswordByTokenResponse ChangePasswordByToken(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequiredChangePasswordByTokenRequest**](RequiredChangePasswordByTokenRequest.md)|  | 

### Return type

[**RequiredChangePasswordByTokenResponse**](RequiredChangePasswordByTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllowResourceIds**
> GetAllowResourcesIdsResponse GetAllowResourceIds(ctx, url)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **url** | **string**|  | 

### Return type

[**GetAllowResourcesIdsResponse**](GetAllowResourcesIdsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoginViaATM**
> LoginViaAtmResponse LoginViaATM(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LoginViaAtmRequest**](LoginViaAtmRequest.md)|  | 

### Return type

[**LoginViaAtmResponse**](LoginViaATMResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoginViaLoginPage**
> LoginViaLoginPageResponse LoginViaLoginPage(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LoginViaLoginPageRequest**](LoginViaLoginPageRequest.md)|  | 

### Return type

[**LoginViaLoginPageResponse**](LoginViaLoginPageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyMfaTokenCode**
> LoginViaPageWithMfaResponse VerifyMfaTokenCode(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LoginViaPageWithMfaCodeRequest**](LoginViaPageWithMfaCodeRequest.md)|  | 

### Return type

[**LoginViaPageWithMfaResponse**](LoginViaPageWithMfaResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

