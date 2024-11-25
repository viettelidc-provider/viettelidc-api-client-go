/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type ListResourceTypeDto struct {
	Id          string               `json:"id,omitempty"`
	Value       string               `json:"value,omitempty"`
	Description string               `json:"description,omitempty"`
	VttService  *DetailVttServiceDto `json:"vtt_service,omitempty"`
	IsDefault   bool                 `json:"is_default,omitempty"`
}