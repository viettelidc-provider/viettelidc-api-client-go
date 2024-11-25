# {{classname}}

All URIs are relative to *http://localhost:7202*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListAccountActivityLog**](AccountActivityLogsAdminApi.md#GetListAccountActivityLog) | **Get** /api/v1/admin/account-activity-log | 

# **GetListAccountActivityLog**
> GetListAccountActivityLogAdminResponse GetListAccountActivityLog(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountActivityLogsAdminApiGetListAccountActivityLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountActivityLogsAdminApiGetListAccountActivityLogOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vttServiceName** | **optional.String**|  | 
 **accountId** | **optional.String**|  | 
 **domainId** | **optional.String**|  | 
 **domainEmail** | **optional.String**|  | 
 **accountRole** | **optional.String**|  | 
 **userName** | **optional.String**|  | 
 **email** | **optional.String**|  | 
 **isRootAccount** | **optional.String**|  | 
 **resourceType** | **optional.String**|  | 
 **action** | **optional.String**|  | 
 **url** | **optional.String**|  | 
 **method** | **optional.String**|  | 
 **requestBodyJson** | **optional.String**|  | 
 **responseBodyJson** | **optional.String**|  | 
 **httpMessage** | **optional.String**|  | 
 **httpStatus** | **optional.Int32**|  | 
 **clientIpAddress** | **optional.String**|  | 
 **currentNodeIpAddress** | **optional.String**|  | 
 **pageIndex** | **optional.Int32**| Zero-based page index (0..N) | [default to 0]
 **pageSize** | **optional.Int32**| The size of the page to be returned | [default to 10]
 **sort** | [**optional.Interface of []string**](string.md)| Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported. | 

### Return type

[**GetListAccountActivityLogAdminResponse**](GetListAccountActivityLogAdminResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

