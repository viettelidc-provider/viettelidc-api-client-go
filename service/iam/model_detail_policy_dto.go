/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

import (
	"time"
)

type DetailPolicyDto struct {
	Id               string                        `json:"id,omitempty"`
	Name             string                        `json:"name,omitempty"`
	Description      string                        `json:"description,omitempty"`
	Type_            string                        `json:"type,omitempty"`
	Scope            string                        `json:"scope,omitempty"`
	IsActive         bool                          `json:"is_active,omitempty"`
	Permissions      []PolicyResourcePermissionDto `json:"permissions,omitempty"`
	LastModifiedDate time.Time                     `json:"last_modified_date,omitempty"`
}
