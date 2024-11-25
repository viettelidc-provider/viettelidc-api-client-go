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

type DetailServiceEndpointDto struct {
	Id               string    `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Description      string    `json:"description,omitempty"`
	Url              string    `json:"url,omitempty"`
	Method           string    `json:"method,omitempty"`
	BodyProperties   string    `json:"body_properties,omitempty"`
	IsActive         bool      `json:"is_active,omitempty"`
	VttServiceId     string    `json:"vtt_service_id,omitempty"`
	LastModifiedDate time.Time `json:"last_modified_date,omitempty"`
}
