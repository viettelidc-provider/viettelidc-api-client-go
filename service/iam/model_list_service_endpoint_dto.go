/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type ListServiceEndpointDto struct {
	Id             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	Url            string `json:"url,omitempty"`
	Method         string `json:"method,omitempty"`
	BodyProperties string `json:"body_properties,omitempty"`
	IsActive       bool   `json:"is_active,omitempty"`
	VttServiceId   string `json:"vtt_service_id,omitempty"`
}
