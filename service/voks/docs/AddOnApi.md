# {{classname}}

All URIs are relative to *https://dev-api.cmp.vtdc.local/csa*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDetailAddon**](AddOnApi.md#GetDetailAddon) | **Post** /api/v1/iac/kubernetes/cluster/add-on/detail | Get detail Addon
[**InstallAddOn**](AddOnApi.md#InstallAddOn) | **Post** /api/v1/iac/kubernetes/cluster/add-on/install | Install Add On
[**UninstallAddOn**](AddOnApi.md#UninstallAddOn) | **Post** /api/v1/iac/kubernetes/cluster/add-on/uninstall | Uninstall Add On

# **GetDetailAddon**
> GetDetailAddOnResponse GetDetailAddon(ctx, body)
Get detail Addon

Retrieve detailed information about a specific add-on within a Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetDetailAddonRequest**](GetDetailAddonRequest.md)|  | 

### Return type

[**GetDetailAddOnResponse**](GetDetailAddOnResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
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

