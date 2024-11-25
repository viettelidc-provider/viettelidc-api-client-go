# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateIAMAccount**](AccountClientApi.md#ActivateIAMAccount) | **Put** /api/v1/account/activate/{accountId} | 
[**AddAccountToUserGroup1**](AccountClientApi.md#AddAccountToUserGroup1) | **Put** /api/v1/account/add-to-groups | 
[**AttachPoliciesByService**](AccountClientApi.md#AttachPoliciesByService) | **Put** /api/v1/account/attach-policies-by-service | 
[**AttachPolicyToUser1**](AccountClientApi.md#AttachPolicyToUser1) | **Put** /api/v1/account/detach-policies | 
[**AttachPolicyToUser2**](AccountClientApi.md#AttachPolicyToUser2) | **Put** /api/v1/account/attach-policies | 
[**ChangePasswordIAMAccount**](AccountClientApi.md#ChangePasswordIAMAccount) | **Put** /api/v1/account/change-password | 
[**CreateNewAccount**](AccountClientApi.md#CreateNewAccount) | **Post** /api/v1/account | 
[**DeactivateIAMAccount**](AccountClientApi.md#DeactivateIAMAccount) | **Put** /api/v1/account/deactivate/{accountId} | 
[**DeleteAccount**](AccountClientApi.md#DeleteAccount) | **Delete** /api/v1/account/{accountId} | 
[**DisableMFAForIAMAccount**](AccountClientApi.md#DisableMFAForIAMAccount) | **Put** /api/v1/account/mfa/disable/{accountId} | 
[**DisableMFARootAccount**](AccountClientApi.md#DisableMFARootAccount) | **Put** /api/v1/account/root/mfa/disable | 
[**EnableMFAForIAMAccount**](AccountClientApi.md#EnableMFAForIAMAccount) | **Put** /api/v1/account/mfa/enable/{accountId} | 
[**EnableMFARootAccount**](AccountClientApi.md#EnableMFARootAccount) | **Put** /api/v1/account/root/mfa/enable | 
[**GetAccountBelongToGroup1**](AccountClientApi.md#GetAccountBelongToGroup1) | **Get** /api/v1/account/belong-to/{userGroupId} | 
[**GetAccountInfoClient**](AccountClientApi.md#GetAccountInfoClient) | **Get** /api/v1/account/info | 
[**GetAccountNotInGroup**](AccountClientApi.md#GetAccountNotInGroup) | **Get** /api/v1/account/not-in/{userGroupId} | 
[**GetAccountProfile**](AccountClientApi.md#GetAccountProfile) | **Get** /api/v1/account/profile | 
[**GetAllUser1**](AccountClientApi.md#GetAllUser1) | **Get** /api/v1/account | 
[**GetDetail1**](AccountClientApi.md#GetDetail1) | **Get** /api/v1/account/{accountId} | 
[**GetMFACodeOfIAMAccount**](AccountClientApi.md#GetMFACodeOfIAMAccount) | **Get** /api/v1/account/mfa/code/{accountId} | 
[**GetMFACodeOfRootAccount**](AccountClientApi.md#GetMFACodeOfRootAccount) | **Get** /api/v1/account/root/mfa/code | 
[**RemoveAccountUserGroup**](AccountClientApi.md#RemoveAccountUserGroup) | **Put** /api/v1/account/remove-groups | 
[**ResetPasswordIAMAccount**](AccountClientApi.md#ResetPasswordIAMAccount) | **Put** /api/v1/account/reset-password | 
[**UnlockIAMAccount**](AccountClientApi.md#UnlockIAMAccount) | **Put** /api/v1/account/unlock/{accountId} | 

# **ActivateIAMAccount**
> ActivateIamAccountResponse ActivateIAMAccount(ctx, accountId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**ActivateIamAccountResponse**](ActivateIAMAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddAccountToUserGroup1**
> AddAccountToUserGroupsResponses AddAccountToUserGroup1(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddAccountToUserGroupsRequest**](AddAccountToUserGroupsRequest.md)|  | 

### Return type

[**AddAccountToUserGroupsResponses**](AddAccountToUserGroupsResponses.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttachPoliciesByService**
> AttachPoliciesByVttServiceResponse AttachPoliciesByService(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AttachPoliciesByVttServiceRequest**](AttachPoliciesByVttServiceRequest.md)|  | 

### Return type

[**AttachPoliciesByVttServiceResponse**](AttachPoliciesByVttServiceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttachPolicyToUser1**
> DetachAccountPolicyResponse AttachPolicyToUser1(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DetachAccountPolicyRequest**](DetachAccountPolicyRequest.md)|  | 

### Return type

[**DetachAccountPolicyResponse**](DetachAccountPolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttachPolicyToUser2**
> AttachAccountPolicyResponse AttachPolicyToUser2(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AttachAccountPolicyRequest**](AttachAccountPolicyRequest.md)|  | 

### Return type

[**AttachAccountPolicyResponse**](AttachAccountPolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangePasswordIAMAccount**
> ChangePasswordIamAccountResponse ChangePasswordIAMAccount(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangePasswordIamAccountRequest**](ChangePasswordIamAccountRequest.md)|  | 

### Return type

[**ChangePasswordIamAccountResponse**](ChangePasswordIAMAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNewAccount**
> SaveAccountResponse CreateNewAccount(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SaveAccountRequest**](SaveAccountRequest.md)|  | 

### Return type

[**SaveAccountResponse**](SaveAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeactivateIAMAccount**
> DeactivateIamAccountResponse DeactivateIAMAccount(ctx, accountId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**DeactivateIamAccountResponse**](DeactivateIAMAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccount**
> DeleteAccountResponse DeleteAccount(ctx, accountId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**DeleteAccountResponse**](DeleteAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableMFAForIAMAccount**
> DisableMfaForIamAccountResponse DisableMFAForIAMAccount(ctx, accountId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**DisableMfaForIamAccountResponse**](DisableMFAForIAMAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableMFARootAccount**
> DisableMfaRootAccountResponse DisableMFARootAccount(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DisableMfaRootAccountResponse**](DisableMFARootAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableMFAForIAMAccount**
> EnableMfaForIamAccountResponse EnableMFAForIAMAccount(ctx, accountId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**EnableMfaForIamAccountResponse**](EnableMFAForIAMAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableMFARootAccount**
> EnableMfaRootAccountResponse EnableMFARootAccount(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EnableMfaRootAccountResponse**](EnableMFARootAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountBelongToGroup1**
> GetAccountsInGroupResponse GetAccountBelongToGroup1(ctx, userGroupId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **string**|  | 
 **optional** | ***AccountClientApiGetAccountBelongToGroup1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountClientApiGetAccountBelongToGroup1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAccountsInGroupResponse**](GetAccountsInGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountInfoClient**
> GetAccountInfoFromTokenResponse GetAccountInfoClient(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetAccountInfoFromTokenResponse**](GetAccountInfoFromTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountNotInGroup**
> GetAccountNotInGroupResponse GetAccountNotInGroup(ctx, userGroupId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **string**|  | 
 **optional** | ***AccountClientApiGetAccountNotInGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountClientApiGetAccountNotInGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAccountNotInGroupResponse**](GetAccountNotInGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountProfile**
> GetAccountProfileResponse GetAccountProfile(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetAccountProfileResponse**](GetAccountProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllUser1**
> GetAllAccountOwnedResponse GetAllUser1(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountClientApiGetAllUser1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountClientApiGetAllUser1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainIdContains** | **optional.String**|  | 
 **domainIdDoesNotContain** | **optional.String**|  | 
 **domainIdEquals** | **optional.String**|  | 
 **domainIdNotEquals** | **optional.String**|  | 
 **domainIdSpecified** | **optional.Bool**|  | 
 **domainIdIn** | [**optional.Interface of []string**](string.md)|  | 
 **domainIdNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **nameContains** | **optional.String**|  | 
 **nameDoesNotContain** | **optional.String**|  | 
 **nameEquals** | **optional.String**|  | 
 **nameNotEquals** | **optional.String**|  | 
 **nameSpecified** | **optional.Bool**|  | 
 **nameIn** | [**optional.Interface of []string**](string.md)|  | 
 **nameNotIn** | [**optional.Interface of []string**](string.md)|  | 
 **deletedDateSpecified** | **optional.Bool**|  | 
 **deletedDateIn** | [**optional.Interface of []interface{}**](interface{}.md)|  | 
 **deletedDateNotIn** | [**optional.Interface of []interface{}**](interface{}.md)|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllAccountOwnedResponse**](GetAllAccountOwnedResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDetail1**
> GetDetailAccountResponse GetDetail1(ctx, accountId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**GetDetailAccountResponse**](GetDetailAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMFACodeOfIAMAccount**
> GetMfaCodeOfIamAccountResponse GetMFACodeOfIAMAccount(ctx, accountId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**GetMfaCodeOfIamAccountResponse**](GetMFACodeOfIAMAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMFACodeOfRootAccount**
> GetMfaCodeOfRootAccountResponse GetMFACodeOfRootAccount(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetMfaCodeOfRootAccountResponse**](GetMFACodeOfRootAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAccountUserGroup**
> RemoveAccountUserGroupsResponses RemoveAccountUserGroup(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveAccountUserGroupsRequest**](RemoveAccountUserGroupsRequest.md)|  | 

### Return type

[**RemoveAccountUserGroupsResponses**](RemoveAccountUserGroupsResponses.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetPasswordIAMAccount**
> ResetPasswordIamAccountResponse ResetPasswordIAMAccount(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResetPasswordIamAccountRequest**](ResetPasswordIamAccountRequest.md)|  | 

### Return type

[**ResetPasswordIamAccountResponse**](ResetPasswordIAMAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlockIAMAccount**
> UnlockIamAccountResponse UnlockIAMAccount(ctx, accountId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**UnlockIamAccountResponse**](UnlockIAMAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

