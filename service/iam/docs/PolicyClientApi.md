# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePolicyByPolicyId**](PolicyClientApi.md#DeletePolicyByPolicyId) | **Delete** /api/v1/policy/{policyId} | 
[**GetAllInlinePoliciesByAccount**](PolicyClientApi.md#GetAllInlinePoliciesByAccount) | **Get** /api/v1/policy/inline/{accountId} | 
[**GetAllPolicies**](PolicyClientApi.md#GetAllPolicies) | **Get** /api/v1/policy | 
[**GetAllPolicyBelongToAccount**](PolicyClientApi.md#GetAllPolicyBelongToAccount) | **Get** /api/v1/policy/belong-to-account/{accountId} | 
[**GetAllPolicyBelongToUserGroup**](PolicyClientApi.md#GetAllPolicyBelongToUserGroup) | **Get** /api/v1/policy/belong-to-user-group/{userGroupId} | 
[**GetAllPolicyNotIn**](PolicyClientApi.md#GetAllPolicyNotIn) | **Get** /api/v1/policy/not-in-group/{userGroupId} | 
[**GetAllStaticPolicy**](PolicyClientApi.md#GetAllStaticPolicy) | **Get** /api/v1/policy/static | 
[**GetDetailPolicyResponse**](PolicyClientApi.md#GetDetailPolicyResponse) | **Get** /api/v1/policy/{policyId} | 
[**GetPoliciesAccountNotHas**](PolicyClientApi.md#GetPoliciesAccountNotHas) | **Get** /api/v1/policy/account-not-has/{accountId} | 
[**GetPolicyByUser**](PolicyClientApi.md#GetPolicyByUser) | **Get** /api/v1/policy/custom | 
[**SavePolicy**](PolicyClientApi.md#SavePolicy) | **Post** /api/v1/policy | 
[**UpdatePolicy**](PolicyClientApi.md#UpdatePolicy) | **Put** /api/v1/policy | 

# **DeletePolicyByPolicyId**
> DeletePolicyResponse DeletePolicyByPolicyId(ctx, policyId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 

### Return type

[**DeletePolicyResponse**](DeletePolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllInlinePoliciesByAccount**
> GetAllInlinePoliciesByAccountResponse GetAllInlinePoliciesByAccount(ctx, accountId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
 **optional** | ***PolicyClientApiGetAllInlinePoliciesByAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyClientApiGetAllInlinePoliciesByAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllInlinePoliciesByAccountResponse**](GetAllInlinePoliciesByAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPolicies**
> GetAllPoliciesResponse GetAllPolicies(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyClientApiGetAllPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyClientApiGetAllPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainIdContains** | **optional.String**|  | 
 **domainIdDoesNotContain** | **optional.String**|  | 
 **domainIdEquals** | **optional.String**|  | 
 **domainIdNotEquals** | **optional.String**|  | 
 **domainIdSpecified** | **optional.Bool**|  | 
 **domainIdIn** | [**optional.Interface of []string**](string.md)|  | 
 **domainIdNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **isActiveEquals** | **optional.Bool**|  | 
 **isActiveNotEquals** | **optional.Bool**|  | 
 **isActiveSpecified** | **optional.Bool**|  | 
 **isActiveIn** | [**optional.Interface of []bool**](bool.md)|  | 
 **isActiveNotIn** | [**optional.Interface of []bool**](bool.md)|  | 
 **deletedDateSpecified** | **optional.Bool**|  | 
 **deletedDateIn** | [**optional.Interface of []interface{}**](interface{}.md)|  | 
 **deletedDateNotIn** | [**optional.Interface of []interface{}**](interface{}.md)|  | 
 **policyTypeContains** | **optional.String**|  | 
 **policyTypeDoesNotContain** | **optional.String**|  | 
 **policyTypeEquals** | **optional.String**|  | 
 **policyTypeNotEquals** | **optional.String**|  | 
 **policyTypeSpecified** | **optional.Bool**|  | 
 **policyTypeIn** | [**optional.Interface of []string**](string.md)|  | 
 **policyTypeNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **policyScopeContains** | **optional.String**|  | 
 **policyScopeDoesNotContain** | **optional.String**|  | 
 **policyScopeEquals** | **optional.String**|  | 
 **policyScopeNotEquals** | **optional.String**|  | 
 **policyScopeSpecified** | **optional.Bool**|  | 
 **policyScopeIn** | [**optional.Interface of []string**](string.md)|  | 
 **policyScopeNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **nameContains** | **optional.String**|  | 
 **nameDoesNotContain** | **optional.String**|  | 
 **nameEquals** | **optional.String**|  | 
 **nameNotEquals** | **optional.String**|  | 
 **nameSpecified** | **optional.Bool**|  | 
 **nameIn** | [**optional.Interface of []string**](string.md)|  | 
 **nameNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllPoliciesResponse**](GetAllPoliciesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPolicyBelongToAccount**
> GetAllInlineAndUserGroupPoliciesResponse GetAllPolicyBelongToAccount(ctx, accountId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
 **optional** | ***PolicyClientApiGetAllPolicyBelongToAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyClientApiGetAllPolicyBelongToAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllInlineAndUserGroupPoliciesResponse**](GetAllInlineAndUserGroupPoliciesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPolicyBelongToUserGroup**
> GetAllPolicyBelongToUserGroupResponse GetAllPolicyBelongToUserGroup(ctx, userGroupId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **string**|  | 
 **optional** | ***PolicyClientApiGetAllPolicyBelongToUserGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyClientApiGetAllPolicyBelongToUserGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllPolicyBelongToUserGroupResponse**](GetAllPolicyBelongToUserGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPolicyNotIn**
> GetAllPoliciesNotInUserGroupResponse GetAllPolicyNotIn(ctx, userGroupId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **string**|  | 
 **optional** | ***PolicyClientApiGetAllPolicyNotInOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyClientApiGetAllPolicyNotInOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllPoliciesNotInUserGroupResponse**](GetAllPoliciesNotInUserGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllStaticPolicy**
> GetAllStaticPolicyResponse GetAllStaticPolicy(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyClientApiGetAllStaticPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyClientApiGetAllStaticPolicyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainIdContains** | **optional.String**|  | 
 **domainIdDoesNotContain** | **optional.String**|  | 
 **domainIdEquals** | **optional.String**|  | 
 **domainIdNotEquals** | **optional.String**|  | 
 **domainIdSpecified** | **optional.Bool**|  | 
 **domainIdIn** | [**optional.Interface of []string**](string.md)|  | 
 **domainIdNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **isActiveEquals** | **optional.Bool**|  | 
 **isActiveNotEquals** | **optional.Bool**|  | 
 **isActiveSpecified** | **optional.Bool**|  | 
 **isActiveIn** | [**optional.Interface of []bool**](bool.md)|  | 
 **isActiveNotIn** | [**optional.Interface of []bool**](bool.md)|  | 
 **deletedDateSpecified** | **optional.Bool**|  | 
 **deletedDateIn** | [**optional.Interface of []interface{}**](interface{}.md)|  | 
 **deletedDateNotIn** | [**optional.Interface of []interface{}**](interface{}.md)|  | 
 **policyTypeContains** | **optional.String**|  | 
 **policyTypeDoesNotContain** | **optional.String**|  | 
 **policyTypeEquals** | **optional.String**|  | 
 **policyTypeNotEquals** | **optional.String**|  | 
 **policyTypeSpecified** | **optional.Bool**|  | 
 **policyTypeIn** | [**optional.Interface of []string**](string.md)|  | 
 **policyTypeNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **policyScopeContains** | **optional.String**|  | 
 **policyScopeDoesNotContain** | **optional.String**|  | 
 **policyScopeEquals** | **optional.String**|  | 
 **policyScopeNotEquals** | **optional.String**|  | 
 **policyScopeSpecified** | **optional.Bool**|  | 
 **policyScopeIn** | [**optional.Interface of []string**](string.md)|  | 
 **policyScopeNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **nameContains** | **optional.String**|  | 
 **nameDoesNotContain** | **optional.String**|  | 
 **nameEquals** | **optional.String**|  | 
 **nameNotEquals** | **optional.String**|  | 
 **nameSpecified** | **optional.Bool**|  | 
 **nameIn** | [**optional.Interface of []string**](string.md)|  | 
 **nameNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllStaticPolicyResponse**](GetAllStaticPolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailPolicyResponse**
> GetDetailPolicyResponse GetDetailPolicyResponse(ctx, policyId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 

### Return type

[**GetDetailPolicyResponse**](GetDetailPolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPoliciesAccountNotHas**
> GetAllPoliciesNotInAccountResponses GetPoliciesAccountNotHas(ctx, accountId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
 **optional** | ***PolicyClientApiGetPoliciesAccountNotHasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyClientApiGetPoliciesAccountNotHasOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllPoliciesNotInAccountResponses**](GetAllPoliciesNotInAccountResponses.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPolicyByUser**
> GetPoliciesByUserResponse GetPolicyByUser(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyClientApiGetPolicyByUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyClientApiGetPolicyByUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainIdContains** | **optional.String**|  | 
 **domainIdDoesNotContain** | **optional.String**|  | 
 **domainIdEquals** | **optional.String**|  | 
 **domainIdNotEquals** | **optional.String**|  | 
 **domainIdSpecified** | **optional.Bool**|  | 
 **domainIdIn** | [**optional.Interface of []string**](string.md)|  | 
 **domainIdNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **isActiveEquals** | **optional.Bool**|  | 
 **isActiveNotEquals** | **optional.Bool**|  | 
 **isActiveSpecified** | **optional.Bool**|  | 
 **isActiveIn** | [**optional.Interface of []bool**](bool.md)|  | 
 **isActiveNotIn** | [**optional.Interface of []bool**](bool.md)|  | 
 **deletedDateSpecified** | **optional.Bool**|  | 
 **deletedDateIn** | [**optional.Interface of []interface{}**](interface{}.md)|  | 
 **deletedDateNotIn** | [**optional.Interface of []interface{}**](interface{}.md)|  | 
 **policyTypeContains** | **optional.String**|  | 
 **policyTypeDoesNotContain** | **optional.String**|  | 
 **policyTypeEquals** | **optional.String**|  | 
 **policyTypeNotEquals** | **optional.String**|  | 
 **policyTypeSpecified** | **optional.Bool**|  | 
 **policyTypeIn** | [**optional.Interface of []string**](string.md)|  | 
 **policyTypeNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **policyScopeContains** | **optional.String**|  | 
 **policyScopeDoesNotContain** | **optional.String**|  | 
 **policyScopeEquals** | **optional.String**|  | 
 **policyScopeNotEquals** | **optional.String**|  | 
 **policyScopeSpecified** | **optional.Bool**|  | 
 **policyScopeIn** | [**optional.Interface of []string**](string.md)|  | 
 **policyScopeNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **nameContains** | **optional.String**|  | 
 **nameDoesNotContain** | **optional.String**|  | 
 **nameEquals** | **optional.String**|  | 
 **nameNotEquals** | **optional.String**|  | 
 **nameSpecified** | **optional.Bool**|  | 
 **nameIn** | [**optional.Interface of []string**](string.md)|  | 
 **nameNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetPoliciesByUserResponse**](GetPoliciesByUserResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SavePolicy**
> SavePolicyResponse SavePolicy(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SavePolicyRequest**](SavePolicyRequest.md)|  | 

### Return type

[**SavePolicyResponse**](SavePolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicy**
> UpdatePolicyResponse UpdatePolicy(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePolicyRequest**](UpdatePolicyRequest.md)|  | 

### Return type

[**UpdatePolicyResponse**](UpdatePolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

