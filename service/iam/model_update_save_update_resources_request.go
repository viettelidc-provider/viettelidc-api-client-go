/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type UpdateSaveUpdateResourcesRequest struct {
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	Id                string `json:"id,omitempty"`
	ResourceTypeId    string `json:"resource_type_id,omitempty"`
	ServiceEndpointId string `json:"service_endpoint_id"`
}
