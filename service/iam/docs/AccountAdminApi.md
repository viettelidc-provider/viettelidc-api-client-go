# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccountToUserGroup**](AccountAdminApi.md#AddAccountToUserGroup) | **Put** /api/v1/admin/account/add-to-group | 
[**AttachDomainToTechSupport**](AccountAdminApi.md#AttachDomainToTechSupport) | **Put** /api/v1/admin/account/tech-support/attach-domain | 
[**AttachPolicyToUser**](AccountAdminApi.md#AttachPolicyToUser) | **Put** /api/v1/admin/account/attach-policies | 
[**ChangeEmailRootAccount**](AccountAdminApi.md#ChangeEmailRootAccount) | **Post** /api/v1/admin/account/root/change-email | 
[**ChangePasswordTechSupport**](AccountAdminApi.md#ChangePasswordTechSupport) | **Put** /api/v1/admin/account/tech-support/change-password | 
[**CreateRootAccount**](AccountAdminApi.md#CreateRootAccount) | **Post** /api/v1/admin/account/create-root | 
[**DeleteTechSupportAccount**](AccountAdminApi.md#DeleteTechSupportAccount) | **Delete** /api/v1/admin/account/tech-support/{accountId} | 
[**DetachDomainForTechSupport**](AccountAdminApi.md#DetachDomainForTechSupport) | **Put** /api/v1/admin/account/tech-support/detach-domain | 
[**DetachPolicyUser**](AccountAdminApi.md#DetachPolicyUser) | **Put** /api/v1/admin/account/detach-policies | 
[**GetAccountBelongToGroup**](AccountAdminApi.md#GetAccountBelongToGroup) | **Get** /api/v1/admin/account/belong-to/{userGroupId} | 
[**GetAccountInfo**](AccountAdminApi.md#GetAccountInfo) | **Get** /api/v1/admin/account/info | 
[**GetAllRootAccountBelongToTechSupport**](AccountAdminApi.md#GetAllRootAccountBelongToTechSupport) | **Get** /api/v1/admin/account/root-account/belong-to-tech-support/{techSupportAccountId} | 
[**GetAllRootAccountNotBelongToTechSupport**](AccountAdminApi.md#GetAllRootAccountNotBelongToTechSupport) | **Get** /api/v1/admin/account/root-account/not-belong-to-tech-support/{techSupportAccountId} | 
[**GetAllRootAccountResponse**](AccountAdminApi.md#GetAllRootAccountResponse) | **Get** /api/v1/admin/account/root | 
[**GetAllUser**](AccountAdminApi.md#GetAllUser) | **Get** /api/v1/admin/account | 
[**GetDetail**](AccountAdminApi.md#GetDetail) | **Get** /api/v1/admin/account/{accountId} | 
[**GetDetailTechSupport**](AccountAdminApi.md#GetDetailTechSupport) | **Get** /api/v1/admin/account/tech-support/{accountId} | 
[**GetListTechSupport**](AccountAdminApi.md#GetListTechSupport) | **Get** /api/v1/admin/account/tech-support | 
[**SaveTechSupportAccount**](AccountAdminApi.md#SaveTechSupportAccount) | **Post** /api/v1/admin/account/tech-support | 
[**SetActiveTechSupport**](AccountAdminApi.md#SetActiveTechSupport) | **Put** /api/v1/admin/account/tech-support/set-active/{accountId} | 
[**SetInactiveTechSupport**](AccountAdminApi.md#SetInactiveTechSupport) | **Put** /api/v1/admin/account/tech-support/set-inactive/{accountId} | 

# **AddAccountToUserGroup**
> AddAccountToUserGroupsResponses AddAccountToUserGroup(ctx, body)


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

# **AttachDomainToTechSupport**
> AttachDomainToTechSupportResponse AttachDomainToTechSupport(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AttachDomainToTechSupportRequest**](AttachDomainToTechSupportRequest.md)|  | 

### Return type

[**AttachDomainToTechSupportResponse**](AttachDomainToTechSupportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttachPolicyToUser**
> AttachAccountPolicyResponse AttachPolicyToUser(ctx, body)


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

# **ChangeEmailRootAccount**
> ChangeEmailRootAccountResponse ChangeEmailRootAccount(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangeEmailRootAccountRequest**](ChangeEmailRootAccountRequest.md)|  | 

### Return type

[**ChangeEmailRootAccountResponse**](ChangeEmailRootAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangePasswordTechSupport**
> ChangePasswordTechSupportResponse ChangePasswordTechSupport(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangePasswordTechSupportRequest**](ChangePasswordTechSupportRequest.md)|  | 

### Return type

[**ChangePasswordTechSupportResponse**](ChangePasswordTechSupportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRootAccount**
> CreateRootAccountResponse CreateRootAccount(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRootAccountRequest**](CreateRootAccountRequest.md)|  | 

### Return type

[**CreateRootAccountResponse**](CreateRootAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTechSupportAccount**
> DeleteTechSupportResponse DeleteTechSupportAccount(ctx, accountId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**DeleteTechSupportResponse**](DeleteTechSupportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachDomainForTechSupport**
> DetachDomainToTechSupportResponse DetachDomainForTechSupport(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DetachDomainToTechSupportRequest**](DetachDomainToTechSupportRequest.md)|  | 

### Return type

[**DetachDomainToTechSupportResponse**](DetachDomainToTechSupportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachPolicyUser**
> DetachAccountPolicyResponse DetachPolicyUser(ctx, body)


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

# **GetAccountBelongToGroup**
> GetAccountsInGroupResponse GetAccountBelongToGroup(ctx, userGroupId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userGroupId** | **string**|  | 
 **optional** | ***AccountAdminApiGetAccountBelongToGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountAdminApiGetAccountBelongToGroupOpts struct
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

# **GetAccountInfo**
> GetAccountInfoFromTokenResponse GetAccountInfo(ctx, )


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

# **GetAllRootAccountBelongToTechSupport**
> GetRootAccountsBelongToTechSupportResponse GetAllRootAccountBelongToTechSupport(ctx, techSupportAccountId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **techSupportAccountId** | **string**|  | 
 **optional** | ***AccountAdminApiGetAllRootAccountBelongToTechSupportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountAdminApiGetAllRootAccountBelongToTechSupportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **email** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetRootAccountsBelongToTechSupportResponse**](GetRootAccountsBelongToTechSupportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllRootAccountNotBelongToTechSupport**
> GetRootAccountsNotBelongToTechSupportResponse GetAllRootAccountNotBelongToTechSupport(ctx, techSupportAccountId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **techSupportAccountId** | **string**|  | 
 **optional** | ***AccountAdminApiGetAllRootAccountNotBelongToTechSupportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountAdminApiGetAllRootAccountNotBelongToTechSupportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **email** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetRootAccountsNotBelongToTechSupportResponse**](GetRootAccountsNotBelongToTechSupportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllRootAccountResponse**
> GetAllRootAccountResponse GetAllRootAccountResponse(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountAdminApiGetAllRootAccountResponseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountAdminApiGetAllRootAccountResponseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetAllRootAccountResponse**](GetAllRootAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllUser**
> GetAllAccountOwnedResponse GetAllUser(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountAdminApiGetAllUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountAdminApiGetAllUserOpts struct
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

# **GetDetail**
> GetDetailAccountResponse GetDetail(ctx, accountId)


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

# **GetDetailTechSupport**
> GetDetailTechSupportResponse GetDetailTechSupport(ctx, accountId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**GetDetailTechSupportResponse**](GetDetailTechSupportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListTechSupport**
> GetListTechSupportResponse GetListTechSupport(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountAdminApiGetListTechSupportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountAdminApiGetListTechSupportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetListTechSupportResponse**](GetListTechSupportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveTechSupportAccount**
> SaveTechSupportResponse SaveTechSupportAccount(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SaveTechSupportRequest**](SaveTechSupportRequest.md)|  | 

### Return type

[**SaveTechSupportResponse**](SaveTechSupportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetActiveTechSupport**
> SetActiveTechSupportResponse SetActiveTechSupport(ctx, accountId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**SetActiveTechSupportResponse**](SetActiveTechSupportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetInactiveTechSupport**
> SetInactiveTechSupportResponse SetInactiveTechSupport(ctx, accountId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**|  | 

### Return type

[**SetInactiveTechSupportResponse**](SetInactiveTechSupportResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

