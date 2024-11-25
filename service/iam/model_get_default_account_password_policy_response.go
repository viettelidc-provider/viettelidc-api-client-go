/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type GetDefaultAccountPasswordPolicyResponse struct {
	Code       int32                             `json:"code,omitempty"`
	Key        string                            `json:"key,omitempty"`
	PageIndex  int32                             `json:"pageIndex,omitempty"`
	PageSize   int32                             `json:"pageSize,omitempty"`
	TotalItems int64                             `json:"totalItems,omitempty"`
	Data       *DefaultAccountPasswordPolicyDto  `json:"data,omitempty"`
	Items      []DefaultAccountPasswordPolicyDto `json:"items,omitempty"`
}