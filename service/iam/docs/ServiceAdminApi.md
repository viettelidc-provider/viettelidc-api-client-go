# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActiveVttServiceByGroupId**](ServiceAdminApi.md#ActiveVttServiceByGroupId) | **Put** /api/v1/admin/vtt-service/active/{vttServiceId} | 
[**GetAllVttServiceToCreateEndpoint1**](ServiceAdminApi.md#GetAllVttServiceToCreateEndpoint1) | **Get** /api/v1/admin/vtt-service/dropdown | 
[**GetDetailVttServiceResponse**](ServiceAdminApi.md#GetDetailVttServiceResponse) | **Get** /api/v1/admin/vtt-service/{vttServiceId} | 
[**GetVttService**](ServiceAdminApi.md#GetVttService) | **Get** /api/v1/admin/vtt-service | 
[**InactiveVttServiceByGroupId**](ServiceAdminApi.md#InactiveVttServiceByGroupId) | **Put** /api/v1/admin/vtt-service/inactive/{vttServiceId} | 
[**SaveVttService**](ServiceAdminApi.md#SaveVttService) | **Post** /api/v1/admin/vtt-service | 
[**UpdateVttService**](ServiceAdminApi.md#UpdateVttService) | **Put** /api/v1/admin/vtt-service | 

# **ActiveVttServiceByGroupId**
> ActiveServiceResponse ActiveVttServiceByGroupId(ctx, vttServiceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vttServiceId** | **string**|  | 

### Return type

[**ActiveServiceResponse**](ActiveServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllVttServiceToCreateEndpoint1**
> GetAllVttServiceToCreateEndpointResponse GetAllVttServiceToCreateEndpoint1(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetAllVttServiceToCreateEndpointResponse**](GetAllVttServiceToCreateEndpointResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailVttServiceResponse**
> GetDetailVttServiceResponse GetDetailVttServiceResponse(ctx, vttServiceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vttServiceId** | **string**|  | 

### Return type

[**GetDetailVttServiceResponse**](GetDetailVttServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVttService**
> GetAllVttServiceResponse GetVttService(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceAdminApiGetVttServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceAdminApiGetVttServiceOpts struct
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

[**GetAllVttServiceResponse**](GetAllVttServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InactiveVttServiceByGroupId**
> InactiveServiceResponse InactiveVttServiceByGroupId(ctx, vttServiceId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vttServiceId** | **string**|  | 

### Return type

[**InactiveServiceResponse**](InactiveServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveVttService**
> SaveVttServiceResponse SaveVttService(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SaveVttServiceRequest**](SaveVttServiceRequest.md)|  | 

### Return type

[**SaveVttServiceResponse**](SaveVttServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVttService**
> UpdateVttServiceResponse UpdateVttService(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVttServiceRequest**](UpdateVttServiceRequest.md)|  | 

### Return type

[**UpdateVttServiceResponse**](UpdateVttServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

