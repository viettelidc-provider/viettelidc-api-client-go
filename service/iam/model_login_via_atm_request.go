/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type LoginViaAtmRequest struct {
	Email    string `json:"email,omitempty"`
	AtmToken string `json:"atmToken,omitempty"`
}
