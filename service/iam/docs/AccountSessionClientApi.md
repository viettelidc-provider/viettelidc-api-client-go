# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeactivateAccountSession**](AccountSessionClientApi.md#DeactivateAccountSession) | **Put** /api/v1/account-session/deactivate | 
[**DeleteAccountSession**](AccountSessionClientApi.md#DeleteAccountSession) | **Delete** /api/v1/account-session | 
[**GetDetailAccountSession**](AccountSessionClientApi.md#GetDetailAccountSession) | **Get** /api/v1/account-session/{id} | 
[**GetListAccountSession**](AccountSessionClientApi.md#GetListAccountSession) | **Get** /api/v1/account-session | 
[**GetListAccountSessionDevice**](AccountSessionClientApi.md#GetListAccountSessionDevice) | **Get** /api/v1/account-session/device | 

# **DeactivateAccountSession**
> DeactivateAccountSessionResponse DeactivateAccountSession(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeactivateAccountSessionRequest**](DeactivateAccountSessionRequest.md)|  | 

### Return type

[**DeactivateAccountSessionResponse**](DeactivateAccountSessionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountSession**
> DeleteAccountSessionResponse DeleteAccountSession(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteAccountSessionRequest**](DeleteAccountSessionRequest.md)|  | 

### Return type

[**DeleteAccountSessionResponse**](DeleteAccountSessionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailAccountSession**
> GetDetailAccountSessionResponse GetDetailAccountSession(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**GetDetailAccountSessionResponse**](GetDetailAccountSessionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListAccountSession**
> GetListAccountSessionResponse GetListAccountSession(ctx, request, pageableRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**AccountSessionCriteria**](.md)|  | 
  **pageableRequest** | [**PageableRequest**](.md)|  | 

### Return type

[**GetListAccountSessionResponse**](GetListAccountSessionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListAccountSessionDevice**
> AccountSessionGroupByDeviceResponse GetListAccountSessionDevice(ctx, request)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**AccountSessionCriteria**](.md)|  | 

### Return type

[**AccountSessionGroupByDeviceResponse**](AccountSessionGroupByDeviceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

