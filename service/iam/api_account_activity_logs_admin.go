/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

import (
	"context"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type AccountActivityLogsAdminApiService service

/*
AccountActivityLogsAdminApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *AccountActivityLogsAdminApiGetListAccountActivityLogOpts - Optional Parameters:
     * @param "VttServiceName" (optional.String) -
     * @param "AccountId" (optional.String) -
     * @param "DomainId" (optional.String) -
     * @param "DomainEmail" (optional.String) -
     * @param "AccountRole" (optional.String) -
     * @param "UserName" (optional.String) -
     * @param "Email" (optional.String) -
     * @param "IsRootAccount" (optional.String) -
     * @param "ResourceType" (optional.String) -
     * @param "Action" (optional.String) -
     * @param "Url" (optional.String) -
     * @param "Method" (optional.String) -
     * @param "RequestBodyJson" (optional.String) -
     * @param "ResponseBodyJson" (optional.String) -
     * @param "HttpMessage" (optional.String) -
     * @param "HttpStatus" (optional.Int32) -
     * @param "ClientIpAddress" (optional.String) -
     * @param "CurrentNodeIpAddress" (optional.String) -
     * @param "PageIndex" (optional.Int32) -  Zero-based page index (0..N)
     * @param "PageSize" (optional.Int32) -  The size of the page to be returned
     * @param "Sort" (optional.Interface of []string) -  Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported.
@return GetListAccountActivityLogAdminResponse
*/

type AccountActivityLogsAdminApiGetListAccountActivityLogOpts struct {
	VttServiceName       optional.String
	AccountId            optional.String
	DomainId             optional.String
	DomainEmail          optional.String
	AccountRole          optional.String
	UserName             optional.String
	Email                optional.String
	IsRootAccount        optional.String
	ResourceType         optional.String
	Action               optional.String
	Url                  optional.String
	Method               optional.String
	RequestBodyJson      optional.String
	ResponseBodyJson     optional.String
	HttpMessage          optional.String
	HttpStatus           optional.Int32
	ClientIpAddress      optional.String
	CurrentNodeIpAddress optional.String
	PageIndex            optional.Int32
	PageSize             optional.Int32
	Sort                 optional.Interface
}

func (a *AccountActivityLogsAdminApiService) GetListAccountActivityLog(ctx context.Context, localVarOptionals *AccountActivityLogsAdminApiGetListAccountActivityLogOpts) (GetListAccountActivityLogAdminResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GetListAccountActivityLogAdminResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/admin/account-activity-log"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.VttServiceName.IsSet() {
		localVarQueryParams.Add("vttServiceName", parameterToString(localVarOptionals.VttServiceName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccountId.IsSet() {
		localVarQueryParams.Add("accountId", parameterToString(localVarOptionals.AccountId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DomainId.IsSet() {
		localVarQueryParams.Add("domainId", parameterToString(localVarOptionals.DomainId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DomainEmail.IsSet() {
		localVarQueryParams.Add("domainEmail", parameterToString(localVarOptionals.DomainEmail.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AccountRole.IsSet() {
		localVarQueryParams.Add("accountRole", parameterToString(localVarOptionals.AccountRole.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UserName.IsSet() {
		localVarQueryParams.Add("userName", parameterToString(localVarOptionals.UserName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Email.IsSet() {
		localVarQueryParams.Add("email", parameterToString(localVarOptionals.Email.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IsRootAccount.IsSet() {
		localVarQueryParams.Add("isRootAccount", parameterToString(localVarOptionals.IsRootAccount.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ResourceType.IsSet() {
		localVarQueryParams.Add("resourceType", parameterToString(localVarOptionals.ResourceType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Action.IsSet() {
		localVarQueryParams.Add("action", parameterToString(localVarOptionals.Action.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Url.IsSet() {
		localVarQueryParams.Add("url", parameterToString(localVarOptionals.Url.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Method.IsSet() {
		localVarQueryParams.Add("method", parameterToString(localVarOptionals.Method.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RequestBodyJson.IsSet() {
		localVarQueryParams.Add("requestBodyJson", parameterToString(localVarOptionals.RequestBodyJson.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ResponseBodyJson.IsSet() {
		localVarQueryParams.Add("responseBodyJson", parameterToString(localVarOptionals.ResponseBodyJson.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HttpMessage.IsSet() {
		localVarQueryParams.Add("httpMessage", parameterToString(localVarOptionals.HttpMessage.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HttpStatus.IsSet() {
		localVarQueryParams.Add("httpStatus", parameterToString(localVarOptionals.HttpStatus.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientIpAddress.IsSet() {
		localVarQueryParams.Add("clientIpAddress", parameterToString(localVarOptionals.ClientIpAddress.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CurrentNodeIpAddress.IsSet() {
		localVarQueryParams.Add("currentNodeIpAddress", parameterToString(localVarOptionals.CurrentNodeIpAddress.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageIndex.IsSet() {
		localVarQueryParams.Add("pageIndex", parameterToString(localVarOptionals.PageIndex.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("pageSize", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		localVarQueryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v GetListAccountActivityLogAdminResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
