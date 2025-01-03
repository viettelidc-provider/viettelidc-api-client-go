/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type GetAccountInfoFromTokenDto struct {
	Id            string `json:"id,omitempty"`
	Email         string `json:"email,omitempty"`
	Username      string `json:"username,omitempty"`
	Role          string `json:"role,omitempty"`
	CustomerId    string `json:"customer_id,omitempty"`
	DomainId      string `json:"domain_id,omitempty"`
	IsRootAccount bool   `json:"is_root_account,omitempty"`
	IsUsingMfa    bool   `json:"is_using_mfa,omitempty"`
}
