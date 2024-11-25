# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceEndPointByUserId**](ServiceEndpointClientApi.md#GetServiceEndPointByUserId) | **Get** /api/v1/service-end-point/has-permission | 

# **GetServiceEndPointByUserId**
> GetAllEndpointHasPermissionResponse GetServiceEndPointByUserId(ctx, permissionIds)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **permissionIds** | [**[]string**](string.md)|  | 

### Return type

[**GetAllEndpointHasPermissionResponse**](GetAllEndpointHasPermissionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

