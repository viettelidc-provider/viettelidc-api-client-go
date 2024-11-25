# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSamlMetadataFile**](Saml2MetadataServiceProviderClientApi.md#DeleteSamlMetadataFile) | **Delete** /api/v1/saml2/service-provider/{id} | 
[**DownloadFile**](Saml2MetadataServiceProviderClientApi.md#DownloadFile) | **Get** /api/v1/saml2/service-provider/download/{id} | 
[**DownloadFileIdp**](Saml2MetadataServiceProviderClientApi.md#DownloadFileIdp) | **Get** /api/v1/saml2/service-provider/download/idp | 
[**GetDetailSamlMetadata**](Saml2MetadataServiceProviderClientApi.md#GetDetailSamlMetadata) | **Get** /api/v1/saml2/service-provider/{id} | 
[**GetListSamlMetadata**](Saml2MetadataServiceProviderClientApi.md#GetListSamlMetadata) | **Get** /api/v1/saml2/service-provider | 
[**UpdateSamlMedataServiceProvider**](Saml2MetadataServiceProviderClientApi.md#UpdateSamlMedataServiceProvider) | **Put** /api/v1/saml2/service-provider/{id} | 
[**UploadSamlMetadataServiceProvider**](Saml2MetadataServiceProviderClientApi.md#UploadSamlMetadataServiceProvider) | **Post** /api/v1/saml2/service-provider | 

# **DeleteSamlMetadataFile**
> DeleteSamlMetadataResponse DeleteSamlMetadataFile(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**DeleteSamlMetadataResponse**](DeleteSamlMetadataResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadFile**
> *os.File DownloadFile(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadFileIdp**
> *os.File DownloadFileIdp(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailSamlMetadata**
> GetDetailSamlMetadataResponse GetDetailSamlMetadata(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**GetDetailSamlMetadataResponse**](GetDetailSamlMetadataResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListSamlMetadata**
> GetListSamlMetadataResponse GetListSamlMetadata(ctx, request, pageableRequest)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**SamlMetadataInformationCriteria**](.md)|  | 
  **pageableRequest** | [**PageableRequest**](.md)|  | 

### Return type

[**GetListSamlMetadataResponse**](GetListSamlMetadataResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSamlMedataServiceProvider**
> UpdateSamlMetadataFileResponse UpdateSamlMedataServiceProvider(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***Saml2MetadataServiceProviderClientApiUpdateSamlMedataServiceProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Saml2MetadataServiceProviderClientApiUpdateSamlMedataServiceProviderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateSamlMetadataFileRequest**](UpdateSamlMetadataFileRequest.md)|  | 

### Return type

[**UpdateSamlMetadataFileResponse**](UpdateSamlMetadataFileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadSamlMetadataServiceProvider**
> UploadSamlMetadataFileResponse UploadSamlMetadataServiceProvider(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Saml2MetadataServiceProviderClientApiUploadSamlMetadataServiceProviderOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Saml2MetadataServiceProviderClientApiUploadSamlMetadataServiceProviderOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UploadSamlMetadataFileRequest**](UploadSamlMetadataFileRequest.md)|  | 

### Return type

[**UploadSamlMetadataFileResponse**](UploadSamlMetadataFileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

