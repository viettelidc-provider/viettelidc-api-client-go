# {{classname}}

All URIs are relative to *https://dev-api.cmp.vtdc.local/csa*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllCluster**](ClusterApi.md#AllCluster) | **Get** /api/v1/iac/kubernetes/cluster/all | Get all cluster
[**DetailCluster**](ClusterApi.md#DetailCluster) | **Get** /api/v1/iac/kubernetes/cluster/{clusterId} | Get detail cluster
[**KubeConfigCluster**](ClusterApi.md#KubeConfigCluster) | **Post** /api/v1/iac/kubernetes/cluster/kube-config | Get cluster kube-config

# **AllCluster**
> []DetailClusterResponse AllCluster(ctx, optional)
Get all cluster

Retrieve list all of the Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClusterApiAllClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterApiAllClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| A query parameter to search for a specific name by exact match. | 

### Return type

[**[]DetailClusterResponse**](array.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetailCluster**
> DetailClusterResponse DetailCluster(ctx, clusterId)
Get detail cluster

Retrieve detailed information about a specific Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| The ID of a Cluster | 

### Return type

[**DetailClusterResponse**](DetailClusterResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KubeConfigCluster**
> GetDetailClusterKubeConfigResponse KubeConfigCluster(ctx, body)
Get cluster kube-config

Retrieve the kube-config for a specific Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseResourceReq**](BaseResourceReq.md)|  | 

### Return type

[**GetDetailClusterKubeConfigResponse**](GetDetailClusterKubeConfigResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

