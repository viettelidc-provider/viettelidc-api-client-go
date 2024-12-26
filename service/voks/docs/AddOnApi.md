# {{classname}}

All URIs are relative to *https://dev-api.cmp.vtdc.local/csa*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllAddOn**](AddOnApi.md#GetAllAddOn) | **Get** /api/v1/iac/kubernetes/cluster/add-on/all | Get all Add On
[**GetAllAddonVersion**](AddOnApi.md#GetAllAddonVersion) | **Get** /api/v1/iac/kubernetes/cluster/add-on/versions | Get all Add On Version (order by lasted version descending)
[**GetDetailAddon**](AddOnApi.md#GetDetailAddon) | **Get** /api/v1/iac/kubernetes/cluster/{clusterId}/add-on | Get detail Addon
[**InstallAddOn**](AddOnApi.md#InstallAddOn) | **Post** /api/v1/iac/kubernetes/cluster/add-on/install | Install Add On
[**UninstallAddOn**](AddOnApi.md#UninstallAddOn) | **Post** /api/v1/iac/kubernetes/cluster/add-on/uninstall | Uninstall Add On

# **GetAllAddOn**
> []Addon GetAllAddOn(ctx, kubernetesVersion, optional)
Get all Add On

By passing in the appropriate options, you can search for available inventory in the system 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **kubernetesVersion** | **string**| Kubernetes version | 
 **optional** | ***AddOnApiGetAllAddOnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddOnApiGetAllAddOnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| A query parameter to search for a specific resource by exact match. | 

### Return type

[**[]Addon**](array.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAddonVersion**
> []AddonVersion GetAllAddonVersion(ctx, addonName, kubernetesVersion, optional)
Get all Add On Version (order by lasted version descending)

By passing in the appropriate options, you can search for available inventory in the system 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addonName** | **string**| The name of the addon | 
  **kubernetesVersion** | **string**| Kubernetes version | 
 **optional** | ***AddOnApiGetAllAddonVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AddOnApiGetAllAddonVersionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.String**| A query parameter to search for a specific version by exact match. | 

### Return type

[**[]AddonVersion**](array.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailAddon**
> GetDetailAddOnResponse GetDetailAddon(ctx, clusterId, addonName)
Get detail Addon

Retrieve detailed information about a specific add-on within a Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| The id of the cluster | 
  **addonName** | **string**| The name of the addon | 

### Return type

[**GetDetailAddOnResponse**](GetDetailAddOnResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstallAddOn**
> InstallAddOn(ctx, body)
Install Add On

Install a specific add-on within a Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddonInstallRequest**](AddonInstallRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UninstallAddOn**
> UninstallAddOn(ctx, body)
Uninstall Add On

Uninstall a specific add-on within a Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddonUninstallRequest**](AddonUninstallRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

