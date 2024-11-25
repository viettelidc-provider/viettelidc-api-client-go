# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDetailVcloudOIDCResource**](OidcAdminQueryControllerApi.md#GetDetailVcloudOIDCResource) | **Get** /api/v1/admin/oidc/vcloud/resources/{accountId} | 
[**GetListVcloudOIDCResource**](OidcAdminQueryControllerApi.md#GetListVcloudOIDCResource) | **Get** /api/v1/admin/oidc/vcloud/resources | 

# **GetDetailVcloudOIDCResource**
> GetDetailVcloudOidcResourceResponse GetDetailVcloudOIDCResource(ctx, accountId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**GetDetailVcloudOidcResourceResponse**](GetDetailVcloudOIDCResourceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListVcloudOIDCResource**
> GetListVcloudOidcResourceResponse GetListVcloudOIDCResource(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OidcAdminQueryControllerApiGetListVcloudOIDCResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OidcAdminQueryControllerApiGetListVcloudOIDCResourceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetListVcloudOidcResourceResponse**](GetListVcloudOIDCResourceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

