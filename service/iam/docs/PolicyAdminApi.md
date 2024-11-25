# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivePolicyByPolicyId**](PolicyAdminApi.md#ActivePolicyByPolicyId) | **Put** /api/v1/admin/policy/active/{policyId} | 
[**GetAllInlinePoliciesByAccount1**](PolicyAdminApi.md#GetAllInlinePoliciesByAccount1) | **Get** /api/v1/admin/policy/inline/{accountId} | 
[**GetAllPolicyBelongToUserGroup1**](PolicyAdminApi.md#GetAllPolicyBelongToUserGroup1) | **Get** /api/v1/admin/policy/belong-to-user-group/{userGroupId} | 
[**GetAllPolicyNotIn1**](PolicyAdminApi.md#GetAllPolicyNotIn1) | **Get** /api/v1/admin/policy/not-in-group/{userGroupId} | 
[**GetAllStaticPolicy1**](PolicyAdminApi.md#GetAllStaticPolicy1) | **Get** /api/v1/admin/policy/static | 
[**GetDetailPolicyResponse1**](PolicyAdminApi.md#GetDetailPolicyResponse1) | **Get** /api/v1/admin/policy/{policyId} | 
[**GetPoliciesAccountNotHas1**](PolicyAdminApi.md#GetPoliciesAccountNotHas1) | **Get** /api/v1/admin/policy/account-not-has/{accountId} | 
[**GetPolicyByUser1**](PolicyAdminApi.md#GetPolicyByUser1) | **Get** /api/v1/admin/policy/custom | 
[**InactivePolicyByPolicyId**](PolicyAdminApi.md#InactivePolicyByPolicyId) | **Put** /api/v1/admin/policy/inactive/{policyId} | 
[**SavePolicy1**](PolicyAdminApi.md#SavePolicy1) | **Post** /api/v1/admin/policy | 
[**UpdatePolicy1**](PolicyAdminApi.md#UpdatePolicy1) | **Put** /api/v1/admin/policy | 

# **ActivePolicyByPolicyId**
> ActivePolicyResponse ActivePolicyByPolicyId(ctx, policyId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 

### Return type

[**ActivePolicyResponse**](ActivePolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllInlinePoliciesByAccount1**
> GetAllInlinePoliciesByAccountResponse GetAllInlinePoliciesByAccount1(ctx, accountId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
 **optional** | ***PolicyAdminApiGetAllInlinePoliciesByAccount1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyAdminApiGetAllInlinePoliciesByAccount1Opts struct
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

# **GetAllPolicyBelongToUserGroup1**
> GetAllPolicyBelongToUserGroupResponse GetAllPolicyBelongToUserGroup1(ctx, userGroupId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **string**|  | 
 **optional** | ***PolicyAdminApiGetAllPolicyBelongToUserGroup1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyAdminApiGetAllPolicyBelongToUserGroup1Opts struct
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

# **GetAllPolicyNotIn1**
> GetAllPoliciesNotInUserGroupResponse GetAllPolicyNotIn1(ctx, userGroupId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **string**|  | 
 **optional** | ***PolicyAdminApiGetAllPolicyNotIn1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyAdminApiGetAllPolicyNotIn1Opts struct
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

# **GetAllStaticPolicy1**
> GetAllStaticPolicyResponse GetAllStaticPolicy1(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyAdminApiGetAllStaticPolicy1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyAdminApiGetAllStaticPolicy1Opts struct
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

# **GetDetailPolicyResponse1**
> GetDetailPolicyAdminResponse GetDetailPolicyResponse1(ctx, policyId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 

### Return type

[**GetDetailPolicyAdminResponse**](GetDetailPolicyAdminResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPoliciesAccountNotHas1**
> GetAllPoliciesNotInAccountResponses GetPoliciesAccountNotHas1(ctx, accountId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
 **optional** | ***PolicyAdminApiGetPoliciesAccountNotHas1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyAdminApiGetPoliciesAccountNotHas1Opts struct
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

# **GetPolicyByUser1**
> GetPoliciesByUserResponse GetPolicyByUser1(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PolicyAdminApiGetPolicyByUser1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PolicyAdminApiGetPolicyByUser1Opts struct
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

# **InactivePolicyByPolicyId**
> ActivePolicyResponse InactivePolicyByPolicyId(ctx, policyId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **string**|  | 

### Return type

[**ActivePolicyResponse**](ActivePolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SavePolicy1**
> SavePolicyResponse SavePolicy1(ctx, body)


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

# **UpdatePolicy1**
> UpdatePolicyResponse UpdatePolicy1(ctx, body)


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

