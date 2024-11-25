# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceType**](ResourceTypeAdminApi.md#CreateResourceType) | **Post** /api/v1/admin/resource-type | 
[**DeleteResourceType**](ResourceTypeAdminApi.md#DeleteResourceType) | **Delete** /api/v1/admin/resource-type/{resourceTypeId} | 
[**GetAllResourceType**](ResourceTypeAdminApi.md#GetAllResourceType) | **Get** /api/v1/admin/resource-type | 
[**GetAllResourceTypeBelongToVttService**](ResourceTypeAdminApi.md#GetAllResourceTypeBelongToVttService) | **Get** /api/v1/admin/resource-type/belong-to-service/{vttServiceId} | 
[**GetDetailResourceType**](ResourceTypeAdminApi.md#GetDetailResourceType) | **Get** /api/v1/admin/resource-type/{resourceTypeId} | 
[**SetDefaultResourceType**](ResourceTypeAdminApi.md#SetDefaultResourceType) | **Patch** /api/v1/admin/resource-type/set-default/{resourceTypeId} | 
[**UnsetDefaultResourceType**](ResourceTypeAdminApi.md#UnsetDefaultResourceType) | **Patch** /api/v1/admin/resource-type/unset-default/{resourceTypeId} | 
[**UpdateResourceType**](ResourceTypeAdminApi.md#UpdateResourceType) | **Patch** /api/v1/admin/resource-type/{resourceTypeId} | 

# **CreateResourceType**
> CreateResourceTypeResponse CreateResourceType(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateResourceTypeRequest**](CreateResourceTypeRequest.md)|  | 

### Return type

[**CreateResourceTypeResponse**](CreateResourceTypeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteResourceType**
> DeleteResourceTypeResponse DeleteResourceType(ctx, resourceTypeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceTypeId** | **string**|  | 

### Return type

[**DeleteResourceTypeResponse**](DeleteResourceTypeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllResourceType**
> GetAllResourceTypeResponse GetAllResourceType(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ResourceTypeAdminApiGetAllResourceTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ResourceTypeAdminApiGetAllResourceTypeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **value** | **optional.String**|  | 
 **vttServiceName** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllResourceTypeResponse**](GetAllResourceTypeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllResourceTypeBelongToVttService**
> GetAllResourceTypeBelongToVttServiceResponse GetAllResourceTypeBelongToVttService(ctx, vttServiceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vttServiceId** | **string**|  | 

### Return type

[**GetAllResourceTypeBelongToVttServiceResponse**](GetAllResourceTypeBelongToVttServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailResourceType**
> GetDetailResourceTypeResponse GetDetailResourceType(ctx, resourceTypeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceTypeId** | **string**|  | 

### Return type

[**GetDetailResourceTypeResponse**](GetDetailResourceTypeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetDefaultResourceType**
> SetDefaultResourceTypeResponse SetDefaultResourceType(ctx, resourceTypeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceTypeId** | **string**|  | 

### Return type

[**SetDefaultResourceTypeResponse**](SetDefaultResourceTypeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsetDefaultResourceType**
> UnsetDefaultResourceTypeResponse UnsetDefaultResourceType(ctx, resourceTypeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceTypeId** | **string**|  | 

### Return type

[**UnsetDefaultResourceTypeResponse**](UnsetDefaultResourceTypeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateResourceType**
> UpdateResourceTypeResponse UpdateResourceType(ctx, body, resourceTypeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateResourceTypeRequest**](UpdateResourceTypeRequest.md)|  | 
  **resourceTypeId** | **string**|  | 

### Return type

[**UpdateResourceTypeResponse**](UpdateResourceTypeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

