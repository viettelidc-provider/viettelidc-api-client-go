/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type ChangeEmailRootAccountRequest struct {
	CustomerId   int32  `json:"customerId,omitempty"`
	CurrentEmail string `json:"currentEmail,omitempty"`
	NewEmail     string `json:"newEmail,omitempty"`
}
