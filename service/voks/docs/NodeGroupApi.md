# {{classname}}

All URIs are relative to *https://dev-api.cmp.vtdc.local/csa*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllNodeGroup**](NodeGroupApi.md#AllNodeGroup) | **Get** /api/v1/iac/kubernetes/cluster/{clusterId}/node-group/all | Get all node group
[**CreateNodeGroup**](NodeGroupApi.md#CreateNodeGroup) | **Post** /api/v1/iac/kubernetes/cluster/group/create | Create node group
[**DeleteNodeGroup**](NodeGroupApi.md#DeleteNodeGroup) | **Post** /api/v1/iac/kubernetes/cluster/group/delete | Delete node group
[**DetailNodeGroup**](NodeGroupApi.md#DetailNodeGroup) | **Get** /api/v1/iac/kubernetes/cluster/{clusterId}/node-group/{groupId} | Get detail node group
[**UpdateNodeGroup**](NodeGroupApi.md#UpdateNodeGroup) | **Post** /api/v1/iac/kubernetes/cluster/group/update | Update node group

# **AllNodeGroup**
> DetailNodeGroupResponse AllNodeGroup(ctx, clusterId, optional)
Get all node group

Retrieve list all of the node group within a Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| The ID of a Cluster | 
 **optional** | ***NodeGroupApiAllNodeGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NodeGroupApiAllNodeGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| A query parameter to search for a specific name by exact match. | 

### Return type

[**DetailNodeGroupResponse**](DetailNodeGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNodeGroup**
> DetailNodeGroupResponse CreateNodeGroup(ctx, body)
Create node group

Create a new node group within a Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateNodeGroupRequest**](CreateNodeGroupRequest.md)|  | 

### Return type

[**DetailNodeGroupResponse**](DetailNodeGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNodeGroup**
> DeleteNodeGroup(ctx, body)
Delete node group

Delete a specific node group from a Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteNodeGroupRequest**](DeleteNodeGroupRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetailNodeGroup**
> DetailNodeGroupResponse DetailNodeGroup(ctx, clusterId, groupId)
Get detail node group

Retrieve detailed information about a specific node group within a Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **int32**| The ID of a Cluster | 
  **groupId** | **int32**| The ID of a Cluster | 

### Return type

[**DetailNodeGroupResponse**](DetailNodeGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNodeGroup**
> DetailNodeGroupResponse UpdateNodeGroup(ctx, body)
Update node group

Update the configuration of an existing node group within a Kubernetes cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateNodeGroupRequest**](UpdateNodeGroupRequest.md)|  | 

### Return type

[**DetailNodeGroupResponse**](DetailNodeGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

