# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccountsToGroup**](UsergroupClientApi.md#AddAccountsToGroup) | **Put** /api/v1/user-group/add-accounts | 
[**AttachPolicyToUserGroup**](UsergroupClientApi.md#AttachPolicyToUserGroup) | **Put** /api/v1/user-group/attach-policies | 
[**DeleteUserGroupByGroupId**](UsergroupClientApi.md#DeleteUserGroupByGroupId) | **Delete** /api/v1/user-group/{userGroupId} | 
[**DetachPoliciesFromUserGroup**](UsergroupClientApi.md#DetachPoliciesFromUserGroup) | **Put** /api/v1/user-group/detach-policies | 
[**GetAllGroupUserBelong**](UsergroupClientApi.md#GetAllGroupUserBelong) | **Get** /api/v1/user-group/has-account/{accountId} | 
[**GetDetailUserGroupResponse**](UsergroupClientApi.md#GetDetailUserGroupResponse) | **Get** /api/v1/user-group/{userGroupId} | 
[**GetGroupUserAccountNotIn**](UsergroupClientApi.md#GetGroupUserAccountNotIn) | **Get** /api/v1/user-group/account-not-in/{accountId} | 
[**GetUserGroupByUserId**](UsergroupClientApi.md#GetUserGroupByUserId) | **Get** /api/v1/user-group | 
[**RemoveAccounts**](UsergroupClientApi.md#RemoveAccounts) | **Put** /api/v1/user-group/remove-accounts | 
[**SaveUserGroup**](UsergroupClientApi.md#SaveUserGroup) | **Post** /api/v1/user-group | 
[**UpdateUserGroup**](UsergroupClientApi.md#UpdateUserGroup) | **Put** /api/v1/user-group | 

# **AddAccountsToGroup**
> AddAccountsToUserGroupResponse AddAccountsToGroup(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddAccountsToUserGroupRequest**](AddAccountsToUserGroupRequest.md)|  | 

### Return type

[**AddAccountsToUserGroupResponse**](AddAccountsToUserGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttachPolicyToUserGroup**
> AttachPolicyToUserGroupResponse AttachPolicyToUserGroup(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AttachPolicyToUserGroupRequest**](AttachPolicyToUserGroupRequest.md)|  | 

### Return type

[**AttachPolicyToUserGroupResponse**](AttachPolicyToUserGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserGroupByGroupId**
> DeleteUserGroupResponse DeleteUserGroupByGroupId(ctx, userGroupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **string**|  | 

### Return type

[**DeleteUserGroupResponse**](DeleteUserGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachPoliciesFromUserGroup**
> DetachPolicyFromUserGroupResponse DetachPoliciesFromUserGroup(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DetachPolicyFromUserGroupRequest**](DetachPolicyFromUserGroupRequest.md)|  | 

### Return type

[**DetachPolicyFromUserGroupResponse**](DetachPolicyFromUserGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllGroupUserBelong**
> GetAllGroupUserBelongResponse GetAllGroupUserBelong(ctx, accountId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
 **optional** | ***UsergroupClientApiGetAllGroupUserBelongOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupClientApiGetAllGroupUserBelongOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllGroupUserBelongResponse**](GetAllGroupUserBelongResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetailUserGroupResponse**
> GetDetailUserGroupResponse GetDetailUserGroupResponse(ctx, userGroupId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **string**|  | 

### Return type

[**GetDetailUserGroupResponse**](GetDetailUserGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGroupUserAccountNotIn**
> GetAllUserGroupsNotInAccountResponse GetGroupUserAccountNotIn(ctx, accountId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
 **optional** | ***UsergroupClientApiGetGroupUserAccountNotInOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupClientApiGetGroupUserAccountNotInOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllUserGroupsNotInAccountResponse**](GetAllUserGroupsNotInAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserGroupByUserId**
> GetAllUserGroupResponse GetUserGroupByUserId(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsergroupClientApiGetUserGroupByUserIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupClientApiGetUserGroupByUserIdOpts struct
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

[**GetAllUserGroupResponse**](GetAllUserGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAccounts**
> RemoveAccountsUserGroupResponses RemoveAccounts(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveAccountsUserGroupRequest**](RemoveAccountsUserGroupRequest.md)|  | 

### Return type

[**RemoveAccountsUserGroupResponses**](RemoveAccountsUserGroupResponses.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveUserGroup**
> SaveUserGroupResponse SaveUserGroup(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SaveUserGroupRequest**](SaveUserGroupRequest.md)|  | 

### Return type

[**SaveUserGroupResponse**](SaveUserGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserGroup**
> UpdateUserGroupResponse UpdateUserGroup(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateUserGroupRequest**](UpdateUserGroupRequest.md)|  | 

### Return type

[**UpdateUserGroupResponse**](UpdateUserGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

