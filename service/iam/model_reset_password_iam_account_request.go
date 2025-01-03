/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type ResetPasswordIamAccountRequest struct {
	NewPassword              string `json:"new_password,omitempty"`
	IsRequiredChangePassword bool   `json:"is_required_change_password,omitempty"`
	AccountId                string `json:"account_id,omitempty"`
}
