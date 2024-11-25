# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccountsToGroup1**](UsergroupAdminApi.md#AddAccountsToGroup1) | **Put** /api/v1/admin/user-group/add-accounts | 
[**AttachPolicyToUserGroup1**](UsergroupAdminApi.md#AttachPolicyToUserGroup1) | **Put** /api/v1/admin/user-group/attach-policies | 
[**DeleteUserGroupByGroupId1**](UsergroupAdminApi.md#DeleteUserGroupByGroupId1) | **Delete** /api/v1/admin/user-group/{userGroupId} | 
[**DetachPoliciesFromUserGroup1**](UsergroupAdminApi.md#DetachPoliciesFromUserGroup1) | **Put** /api/v1/admin/user-group/detach-policies | 
[**GetAllGroupUserBelong1**](UsergroupAdminApi.md#GetAllGroupUserBelong1) | **Get** /api/v1/admin/user-group/has-account/{accountId} | 
[**GetDetailUserGroupResponse1**](UsergroupAdminApi.md#GetDetailUserGroupResponse1) | **Get** /api/v1/admin/user-group/{userGroupId} | 
[**GetGroupUserAccountNotIn1**](UsergroupAdminApi.md#GetGroupUserAccountNotIn1) | **Get** /api/v1/admin/user-group/account-not-in/{accountId} | 
[**GetUserGroupByUserId1**](UsergroupAdminApi.md#GetUserGroupByUserId1) | **Get** /api/v1/admin/user-group | 
[**RemoveAccounts1**](UsergroupAdminApi.md#RemoveAccounts1) | **Put** /api/v1/admin/user-group/remove-accounts | 
[**SaveUserGroup1**](UsergroupAdminApi.md#SaveUserGroup1) | **Post** /api/v1/admin/user-group | 
[**UpdateUserGroup1**](UsergroupAdminApi.md#UpdateUserGroup1) | **Put** /api/v1/admin/user-group | 

# **AddAccountsToGroup1**
> AddAccountsToUserGroupResponse AddAccountsToGroup1(ctx, body)


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

# **AttachPolicyToUserGroup1**
> AttachPolicyToUserGroupResponse AttachPolicyToUserGroup1(ctx, body)


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

# **DeleteUserGroupByGroupId1**
> DeleteUserGroupResponse DeleteUserGroupByGroupId1(ctx, userGroupId)


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

# **DetachPoliciesFromUserGroup1**
> DetachPolicyFromUserGroupResponse DetachPoliciesFromUserGroup1(ctx, body)


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

# **GetAllGroupUserBelong1**
> GetAllGroupUserBelongResponse GetAllGroupUserBelong1(ctx, accountId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
 **optional** | ***UsergroupAdminApiGetAllGroupUserBelong1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupAdminApiGetAllGroupUserBelong1Opts struct
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

# **GetDetailUserGroupResponse1**
> GetDetailUserGroupResponse GetDetailUserGroupResponse1(ctx, userGroupId)


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

# **GetGroupUserAccountNotIn1**
> GetAllUserGroupsNotInAccountResponse GetGroupUserAccountNotIn1(ctx, accountId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 
 **optional** | ***UsergroupAdminApiGetGroupUserAccountNotIn1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupAdminApiGetGroupUserAccountNotIn1Opts struct
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

# **GetUserGroupByUserId1**
> GetAllUserGroupResponse GetUserGroupByUserId1(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsergroupAdminApiGetUserGroupByUserId1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupAdminApiGetUserGroupByUserId1Opts struct
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

# **RemoveAccounts1**
> RemoveAccountsUserGroupResponses RemoveAccounts1(ctx, body)


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

# **SaveUserGroup1**
> SaveUserGroupResponse SaveUserGroup1(ctx, body)


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

# **UpdateUserGroup1**
> UpdateUserGroupResponse UpdateUserGroup1(ctx, body)


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

