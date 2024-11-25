# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActiveResourcesByGroupId**](ResourcesAdminApi.md#ActiveResourcesByGroupId) | **Put** /api/v1/admin/resources/active/{resourceId} | 
[**GetAllResourcesBelongServiceEndpoint**](ResourcesAdminApi.md#GetAllResourcesBelongServiceEndpoint) | **Get** /api/v1/admin/resources/belong-to/{endpointId} | 
[**GetAllResourcesToCreatePermission**](ResourcesAdminApi.md#GetAllResourcesToCreatePermission) | **Get** /api/v1/admin/resources/dropdown | 
[**GetAllResourcesType**](ResourcesAdminApi.md#GetAllResourcesType) | **Get** /api/v1/admin/resources/types | 
[**GetDetailResourcesResponse**](ResourcesAdminApi.md#GetDetailResourcesResponse) | **Get** /api/v1/admin/resources/{resourceId} | 
[**GetResourcesByUserId**](ResourcesAdminApi.md#GetResourcesByUserId) | **Get** /api/v1/admin/resources | 
[**InactiveResourcesByGroupId**](ResourcesAdminApi.md#InactiveResourcesByGroupId) | **Put** /api/v1/admin/resources/inactive/{resourceId} | 
[**SaveResources**](ResourcesAdminApi.md#SaveResources) | **Post** /api/v1/admin/resources | 
[**UpdateResources**](ResourcesAdminApi.md#UpdateResources) | **Put** /api/v1/admin/resources | 

# **ActiveResourcesByGroupId**
> ActiveResourcesResponse ActiveResourcesByGroupId(ctx, resourceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceId** | **string**|  | 

### Return type

[**ActiveResourcesResponse**](ActiveResourcesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllResourcesBelongServiceEndpoint**
> GetAllResourcesBelongServiceEndpointResponse GetAllResourcesBelongServiceEndpoint(ctx, endpointId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointId** | **string**|  | 
 **optional** | ***ResourcesAdminApiGetAllResourcesBelongServiceEndpointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResourcesAdminApiGetAllResourcesBelongServiceEndpointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllResourcesBelongServiceEndpointResponse**](GetAllResourcesBelongServiceEndpointResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllResourcesToCreatePermission**
> GetAllResourcesToCreatePermissionResponse GetAllResourcesToCreatePermission(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetAllResourcesToCreatePermissionResponse**](GetAllResourcesToCreatePermissionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllResourcesType**
> GetAllResourcesTypeUseCaseResponse GetAllResourcesType(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetAllResourcesTypeUseCaseResponse**](GetAllResourcesTypeUseCaseResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailResourcesResponse**
> GetDetailResourcesResponse GetDetailResourcesResponse(ctx, resourceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceId** | **string**|  | 

### Return type

[**GetDetailResourcesResponse**](GetDetailResourcesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResourcesByUserId**
> GetAllResourcesResponse GetResourcesByUserId(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ResourcesAdminApiGetResourcesByUserIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResourcesAdminApiGetResourcesByUserIdOpts struct
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

[**GetAllResourcesResponse**](GetAllResourcesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InactiveResourcesByGroupId**
> InactiveResourcesResponse InactiveResourcesByGroupId(ctx, resourceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceId** | **string**|  | 

### Return type

[**InactiveResourcesResponse**](InactiveResourcesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveResources**
> SaveResourcesResponse SaveResources(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SaveUpdateResourcesRequest**](SaveUpdateResourcesRequest.md)|  | 

### Return type

[**SaveResourcesResponse**](SaveResourcesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateResources**
> UpdateResourcesResponse UpdateResources(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSaveUpdateResourcesRequest**](UpdateSaveUpdateResourcesRequest.md)|  | 

### Return type

[**UpdateResourcesResponse**](UpdateResourcesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

