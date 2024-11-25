/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type GetAllEndpointsBelongServiceNoPagingResponse struct {
	Code       int32                    `json:"code,omitempty"`
	Key        string                   `json:"key,omitempty"`
	PageIndex  int32                    `json:"pageIndex,omitempty"`
	PageSize   int32                    `json:"pageSize,omitempty"`
	TotalItems int64                    `json:"totalItems,omitempty"`
	Data       *ListServiceEndpointDto  `json:"data,omitempty"`
	Items      []ListServiceEndpointDto `json:"items,omitempty"`
}