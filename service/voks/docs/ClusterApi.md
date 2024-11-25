# {{classname}}

All URIs are relative to *https://dev-api.cmp.vtdc.local/csa*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DetailCluster**](ClusterApi.md#DetailCluster) | **Post** /api/v1/iac/kubernetes/cluster/detail | Get detail cluster
[**KubeConfigCluster**](ClusterApi.md#KubeConfigCluster) | **Post** /api/v1/iac/kubernetes/cluster/kube-config | Get cluster kube-config

# **DetailCluster**
> DetailClusterResponse DetailCluster(ctx, body)
Get detail cluster

Retrieve detailed information about a specific Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseResourceReq**](BaseResourceReq.md)|  | 

### Return type

[**DetailClusterResponse**](DetailClusterResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
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

