# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportTSVDataByServiceId**](ImportApiControllerApi.md#ImportTSVDataByServiceId) | **Post** /api/v1/admin/import/local-data/upload-file | 
[**ImportTsvDataFromResourceData**](ImportApiControllerApi.md#ImportTsvDataFromResourceData) | **Post** /api/v1/admin/import/local-data | 

# **ImportTSVDataByServiceId**
> string ImportTSVDataByServiceId(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImportApiControllerApiImportTSVDataByServiceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportApiControllerApiImportTSVDataByServiceIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LocaldataUploadfileBody**](LocaldataUploadfileBody.md)|  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportTsvDataFromResourceData**
> string ImportTsvDataFromResourceData(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

