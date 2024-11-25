# {{classname}}

All URIs are relative to *https://dev-api.cmp.vtdc.local/csa*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DetailNfsStorage**](NFSApi.md#DetailNfsStorage) | **Post** /api/v1/iac/kubernetes/cluster/nfs/detail | Detail NFS Storage
[**ExtendNfsStorage**](NFSApi.md#ExtendNfsStorage) | **Post** /api/v1/iac/kubernetes/cluster/nfs/add-ons | Extend NFS Storage

# **DetailNfsStorage**
> GetDetailNfsResponse DetailNfsStorage(ctx, body)
Detail NFS Storage

Retrieve detailed information about a specific NFS storage within a Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseResourceReq**](BaseResourceReq.md)|  | 

### Return type

[**GetDetailNfsResponse**](GetDetailNFSResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtendNfsStorage**
> ExtendNfsStorage(ctx, body)
Extend NFS Storage

Extend the storage capacity of a specific NFS storage within a Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddonNfsRequest**](AddonNfsRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

