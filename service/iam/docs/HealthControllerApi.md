# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckHealth**](HealthControllerApi.md#CheckHealth) | **Post** /api/v1/iam/check-health | 
[**CheckHealth1**](HealthControllerApi.md#CheckHealth1) | **Get** /api/v1/iam/check-health | 

# **CheckHealth**
> BaseApiResponse CheckHealth(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**map[string]string**](map.md)|  | 

### Return type

[**BaseApiResponse**](BaseApiResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckHealth1**
> BaseApiResponse CheckHealth1(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**BaseApiResponse**](BaseApiResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

