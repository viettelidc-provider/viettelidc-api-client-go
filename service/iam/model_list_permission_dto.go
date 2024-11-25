/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type ListPermissionDto struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Scope       string `json:"scope,omitempty"`
	Permission  string `json:"permission,omitempty"`
	IsActive    bool   `json:"is_active,omitempty"`
	AccessLevel string `json:"access_level,omitempty"`
	ResourcesId string `json:"resources_id,omitempty"`
}
