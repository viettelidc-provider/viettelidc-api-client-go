/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type RequiredChangePasswordRequest struct {
	NewPassword string `json:"newPassword,omitempty"`
	OldPassword string `json:"oldPassword,omitempty"`
}