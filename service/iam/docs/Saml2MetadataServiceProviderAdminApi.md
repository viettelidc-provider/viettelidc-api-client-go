# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSamlMetadataFile1**](Saml2MetadataServiceProviderAdminApi.md#DeleteSamlMetadataFile1) | **Delete** /api/v1/admin/saml2/service-provider/{id} | 
[**UpdateSamlMedataServiceProvider1**](Saml2MetadataServiceProviderAdminApi.md#UpdateSamlMedataServiceProvider1) | **Put** /api/v1/admin/saml2/service-provider/{id} | 
[**UploadSamlMetadataServiceProvider1**](Saml2MetadataServiceProviderAdminApi.md#UploadSamlMetadataServiceProvider1) | **Post** /api/v1/admin/saml2/service-provider | 

# **DeleteSamlMetadataFile1**
> DeleteSamlMetadataResponse DeleteSamlMetadataFile1(ctx, id)


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

# **UpdateSamlMedataServiceProvider1**
> UpdateSamlMetadataFileResponse UpdateSamlMedataServiceProvider1(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***Saml2MetadataServiceProviderAdminApiUpdateSamlMedataServiceProvider1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Saml2MetadataServiceProviderAdminApiUpdateSamlMedataServiceProvider1Opts struct
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

# **UploadSamlMetadataServiceProvider1**
> UploadSamlMetadataFileResponse UploadSamlMetadataServiceProvider1(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Saml2MetadataServiceProviderAdminApiUploadSamlMetadataServiceProvider1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Saml2MetadataServiceProviderAdminApiUploadSamlMetadataServiceProvider1Opts struct
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

