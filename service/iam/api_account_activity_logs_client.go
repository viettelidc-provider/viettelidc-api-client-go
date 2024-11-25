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

type AccountActivityLogsClientApiService service

/*
AccountActivityLogsClientApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *AccountActivityLogsClientApiGetListAccountActivityLog1Opts - Optional Parameters:
     * @param "VttServiceName" (optional.String) -
     * @param "Name" (optional.String) -
     * @param "ResourceType" (optional.String) -
     * @param "Action" (optional.String) -
     * @param "HttpMessage" (optional.String) -
     * @param "PageIndex" (optional.Int32) -  Zero-based page index (0..N)
     * @param "PageSize" (optional.Int32) -  The size of the page to be returned
     * @param "Sort" (optional.Interface of []string) -  Sorting criteria in the format: property,(asc|desc). Default sort order is ascending. Multiple sort criteria are supported.
@return GetListAccountActivityLogClientResponse
*/

type AccountActivityLogsClientApiGetListAccountActivityLog1Opts struct {
	VttServiceName optional.String
	Name           optional.String
	ResourceType   optional.String
	Action         optional.String
	HttpMessage    optional.String
	PageIndex      optional.Int32
	PageSize       optional.Int32
	Sort           optional.Interface
}

func (a *AccountActivityLogsClientApiService) GetListAccountActivityLog1(ctx context.Context, localVarOptionals *AccountActivityLogsClientApiGetListAccountActivityLog1Opts) (GetListAccountActivityLogClientResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue GetListAccountActivityLogClientResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v1/account-activity-log"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.VttServiceName.IsSet() {
		localVarQueryParams.Add("vttServiceName", parameterToString(localVarOptionals.VttServiceName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Name.IsSet() {
		localVarQueryParams.Add("name", parameterToString(localVarOptionals.Name.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ResourceType.IsSet() {
		localVarQueryParams.Add("resourceType", parameterToString(localVarOptionals.ResourceType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Action.IsSet() {
		localVarQueryParams.Add("action", parameterToString(localVarOptionals.Action.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HttpMessage.IsSet() {
		localVarQueryParams.Add("httpMessage", parameterToString(localVarOptionals.HttpMessage.Value(), ""))
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
			var v GetListAccountActivityLogClientResponse
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
