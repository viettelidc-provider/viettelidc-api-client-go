/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type CreateVcloudOidcRegisterClientRequest struct {
	Name         string   `json:"name,omitempty"`
	RedirectUris []string `json:"redirect_uris,omitempty"`
}
