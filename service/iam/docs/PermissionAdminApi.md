# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivePermission**](PermissionAdminApi.md#ActivePermission) | **Put** /api/v1/admin/permission/active/{permissionId} | 
[**DeletePermission**](PermissionAdminApi.md#DeletePermission) | **Delete** /api/v1/admin/permission/{permissionId} | 
[**GetAllEndpointPermissionNeeded**](PermissionAdminApi.md#GetAllEndpointPermissionNeeded) | **Get** /api/v1/admin/permission/needed/{endpointId} | 
[**GetAllPermissionBelongService1**](PermissionAdminApi.md#GetAllPermissionBelongService1) | **Get** /api/v1/admin/permission/belong-to-service | 
[**GetAllPermissionBelongToPolicy2**](PermissionAdminApi.md#GetAllPermissionBelongToPolicy2) | **Get** /api/v1/admin/permission/belong-to-policy/{policyId} | 
[**GetAllPermissionBelongToPolicy3**](PermissionAdminApi.md#GetAllPermissionBelongToPolicy3) | **Get** /api/v1/admin/permission/belong-to-policy/no-paging/{policyId} | 
[**GetDetailPermissionResponse**](PermissionAdminApi.md#GetDetailPermissionResponse) | **Get** /api/v1/admin/permission/{permissionId} | 
[**GetPermissionByUserId**](PermissionAdminApi.md#GetPermissionByUserId) | **Get** /api/v1/admin/permission | 
[**GetPermissionByUserId1**](PermissionAdminApi.md#GetPermissionByUserId1) | **Get** /api/v1/admin/permission/belong-to-resources/{resourcesId} | 
[**InactivePermission**](PermissionAdminApi.md#InactivePermission) | **Put** /api/v1/admin/permission/inactive/{permissionId} | 
[**SavePermission**](PermissionAdminApi.md#SavePermission) | **Post** /api/v1/admin/permission | 
[**UpdatePermission**](PermissionAdminApi.md#UpdatePermission) | **Put** /api/v1/admin/permission | 

# **ActivePermission**
> ActivePermissionResponse ActivePermission(ctx, permissionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **permissionId** | **string**|  | 

### Return type

[**ActivePermissionResponse**](ActivePermissionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePermission**
> DeletePermissionResponse DeletePermission(ctx, permissionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **permissionId** | **string**|  | 

### Return type

[**DeletePermissionResponse**](DeletePermissionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllEndpointPermissionNeeded**
> GetAllEndpointPermissionNeededResponse GetAllEndpointPermissionNeeded(ctx, endpointId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointId** | **string**|  | 

### Return type

[**GetAllEndpointPermissionNeededResponse**](GetAllEndpointPermissionNeededResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPermissionBelongService1**
> GetAllPermissionsBelongServiceResponse GetAllPermissionBelongService1(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermissionAdminApiGetAllPermissionBelongService1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermissionAdminApiGetAllPermissionBelongService1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceId** | **optional.String**|  | 
 **scope** | **optional.String**|  | 

### Return type

[**GetAllPermissionsBelongServiceResponse**](GetAllPermissionsBelongServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPermissionBelongToPolicy2**
> GetAllPermissionBelongToPolicyResponse GetAllPermissionBelongToPolicy2(ctx, policyId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 
 **optional** | ***PermissionAdminApiGetAllPermissionBelongToPolicy2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermissionAdminApiGetAllPermissionBelongToPolicy2Opts struct
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

# **GetAllPermissionBelongToPolicy3**
> GetAllPermissionBelongToPolicyNoPagingResponse GetAllPermissionBelongToPolicy3(ctx, policyId)


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

# **GetDetailPermissionResponse**
> GetDetailPermissionResponse GetDetailPermissionResponse(ctx, permissionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **permissionId** | **string**|  | 

### Return type

[**GetDetailPermissionResponse**](GetDetailPermissionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissionByUserId**
> GetAllPermissionResponse GetPermissionByUserId(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PermissionAdminApiGetPermissionByUserIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermissionAdminApiGetPermissionByUserIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nameContains** | **optional.String**|  | 
 **nameDoesNotContain** | **optional.String**|  | 
 **nameEquals** | **optional.String**|  | 
 **nameNotEquals** | **optional.String**|  | 
 **nameSpecified** | **optional.Bool**|  | 
 **nameIn** | [**optional.Interface of []string**](string.md)|  | 
 **nameNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **deletedDateSpecified** | **optional.Bool**|  | 
 **deletedDateIn** | [**optional.Interface of []interface{}**](interface{}.md)|  | 
 **deletedDateNotIn** | [**optional.Interface of []interface{}**](interface{}.md)|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllPermissionResponse**](GetAllPermissionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPermissionByUserId1**
> GetAllPermissionsBelongResourcesResponse GetPermissionByUserId1(ctx, resourcesId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourcesId** | **string**|  | 
 **optional** | ***PermissionAdminApiGetPermissionByUserId1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PermissionAdminApiGetPermissionByUserId1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllPermissionsBelongResourcesResponse**](GetAllPermissionsBelongResourcesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InactivePermission**
> InactivePermissionResponse InactivePermission(ctx, permissionId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **permissionId** | **string**|  | 

### Return type

[**InactivePermissionResponse**](InactivePermissionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SavePermission**
> SavePermissionResponse SavePermission(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SavePermissionRequest**](SavePermissionRequest.md)|  | 

### Return type

[**SavePermissionResponse**](SavePermissionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePermission**
> UpdatePermissionResponse UpdatePermission(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePermissionRequest**](UpdatePermissionRequest.md)|  | 

### Return type

[**UpdatePermissionResponse**](UpdatePermissionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

