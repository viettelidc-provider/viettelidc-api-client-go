/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type LoginViaPageWithMfaResponse struct {
	Code                     int32    `json:"code,omitempty"`
	Key                      string   `json:"key,omitempty"`
	PageIndex                int32    `json:"pageIndex,omitempty"`
	PageSize                 int32    `json:"pageSize,omitempty"`
	TotalItems               int64    `json:"totalItems,omitempty"`
	Data                     string   `json:"data,omitempty"`
	Items                    []string `json:"items,omitempty"`
	ChangePasswordToken      string   `json:"change_password_token,omitempty"`
	IsRequiredChangePassword bool     `json:"is_required_change_password,omitempty"`
}
