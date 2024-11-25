# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActiveServiceEndPointByGroupId**](ServiceEndpointAdminApi.md#ActiveServiceEndPointByGroupId) | **Put** /api/v1/admin/service-end-point/active/{serviceEndpointId} | 
[**AttachEndPointToVttService**](ServiceEndpointAdminApi.md#AttachEndPointToVttService) | **Put** /api/v1/admin/service-end-point/attach-to/{serviceId} | 
[**GetAllEndpointToCreateResources**](ServiceEndpointAdminApi.md#GetAllEndpointToCreateResources) | **Get** /api/v1/admin/service-end-point/dropdown | 
[**GetAllEndpointsBelongServiceNoPaging**](ServiceEndpointAdminApi.md#GetAllEndpointsBelongServiceNoPaging) | **Get** /api/v1/admin/service-end-point/belong-to/{vttServiceId}/all | 
[**GetDetailServiceEndPointResponse**](ServiceEndpointAdminApi.md#GetDetailServiceEndPointResponse) | **Get** /api/v1/admin/service-end-point/{serviceEndpointId} | 
[**GetEndpointsBelongToService**](ServiceEndpointAdminApi.md#GetEndpointsBelongToService) | **Get** /api/v1/admin/service-end-point/belong-to/{vttServiceId} | 
[**GetServiceEndPointByUserId1**](ServiceEndpointAdminApi.md#GetServiceEndPointByUserId1) | **Get** /api/v1/admin/service-end-point | 
[**InactiveServiceEndPointByGroupId**](ServiceEndpointAdminApi.md#InactiveServiceEndPointByGroupId) | **Put** /api/v1/admin/service-end-point/inactive/{serviceEndpointId} | 
[**SaveServiceEndPoint**](ServiceEndpointAdminApi.md#SaveServiceEndPoint) | **Post** /api/v1/admin/service-end-point | 
[**UpdateServiceEndPoint**](ServiceEndpointAdminApi.md#UpdateServiceEndPoint) | **Put** /api/v1/admin/service-end-point | 

# **ActiveServiceEndPointByGroupId**
> ActiveServiceEndPointResponse ActiveServiceEndPointByGroupId(ctx, serviceEndpointId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceEndpointId** | **string**|  | 

### Return type

[**ActiveServiceEndPointResponse**](ActiveServiceEndPointResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttachEndPointToVttService**
> AttachEndpointToVttServiceResponse AttachEndPointToVttService(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AttachEndpointToVttServiceRequest**](AttachEndpointToVttServiceRequest.md)|  | 

### Return type

[**AttachEndpointToVttServiceResponse**](AttachEndpointToVttServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllEndpointToCreateResources**
> GetAllEndpointToCreateResourcesResponse GetAllEndpointToCreateResources(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetAllEndpointToCreateResourcesResponse**](GetAllEndpointToCreateResourcesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllEndpointsBelongServiceNoPaging**
> GetAllEndpointsBelongServiceNoPagingResponse GetAllEndpointsBelongServiceNoPaging(ctx, vttServiceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vttServiceId** | **string**|  | 

### Return type

[**GetAllEndpointsBelongServiceNoPagingResponse**](GetAllEndpointsBelongServiceNoPagingResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailServiceEndPointResponse**
> GetDetailServiceEndPointResponse GetDetailServiceEndPointResponse(ctx, serviceEndpointId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceEndpointId** | **string**|  | 

### Return type

[**GetDetailServiceEndPointResponse**](GetDetailServiceEndPointResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEndpointsBelongToService**
> GetAllEndpointsBelongServiceResponse GetEndpointsBelongToService(ctx, vttServiceId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vttServiceId** | **string**|  | 
 **optional** | ***ServiceEndpointAdminApiGetEndpointsBelongToServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceEndpointAdminApiGetEndpointsBelongToServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllEndpointsBelongServiceResponse**](GetAllEndpointsBelongServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceEndPointByUserId1**
> GetAllServiceEndPointResponse GetServiceEndPointByUserId1(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceEndpointAdminApiGetServiceEndPointByUserId1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceEndpointAdminApiGetServiceEndPointByUserId1Opts struct
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

[**GetAllServiceEndPointResponse**](GetAllServiceEndPointResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InactiveServiceEndPointByGroupId**
> InactiveServiceEndPointResponse InactiveServiceEndPointByGroupId(ctx, serviceEndpointId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceEndpointId** | **string**|  | 

### Return type

[**InactiveServiceEndPointResponse**](InactiveServiceEndPointResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveServiceEndPoint**
> SaveServiceEndPointResponse SaveServiceEndPoint(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SaveServiceEndPointRequest**](SaveServiceEndPointRequest.md)|  | 

### Return type

[**SaveServiceEndPointResponse**](SaveServiceEndPointResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceEndPoint**
> UpdateServiceEndPointResponse UpdateServiceEndPoint(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateServiceEndPointRequest**](UpdateServiceEndPointRequest.md)|  | 

### Return type

[**UpdateServiceEndPointResponse**](UpdateServiceEndPointResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

