# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllPermissionBelongService**](PermissionClientApi.md#GetAllPermissionBelongService) | **Get** /api/v1/permission/belong-to-service/{serviceId} | 
[**GetAllPermissionBelongToPolicy**](PermissionClientApi.md#GetAllPermissionBelongToPolicy) | **Get** /api/v1/permission/belong-to-policy/{policyId} | 
[**GetAllPermissionBelongToPolicy1**](PermissionClientApi.md#GetAllPermissionBelongToPolicy1) | **Get** /api/v1/permission/belong-to-policy/no-paging/{policyId} | 

# **GetAllPermissionBelongService**
> GetAllPermissionsBelongServiceResponse GetAllPermissionBelongService(ctx, serviceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 

### Return type

[**GetAllPermissionsBelongServiceResponse**](GetAllPermissionsBelongServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPermissionBelongToPolicy**
> GetAllPermissionBelongToPolicyResponse GetAllPermissionBelongToPolicy(ctx, policyId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 
 **optional** | ***PermissionClientApiGetAllPermissionBelongToPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermissionClientApiGetAllPermissionBelongToPolicyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllPermissionBelongToPolicyResponse**](GetAllPermissionBelongToPolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPermissionBelongToPolicy1**
> GetAllPermissionBelongToPolicyNoPagingResponse GetAllPermissionBelongToPolicy1(ctx, policyId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 

### Return type

[**GetAllPermissionBelongToPolicyNoPagingResponse**](GetAllPermissionBelongToPolicyNoPagingResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

