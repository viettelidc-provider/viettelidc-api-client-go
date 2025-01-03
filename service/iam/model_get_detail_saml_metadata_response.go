/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type GetDetailSamlMetadataResponse struct {
	Code       int32                        `json:"code,omitempty"`
	Key        string                       `json:"key,omitempty"`
	PageIndex  int32                        `json:"pageIndex,omitempty"`
	PageSize   int32                        `json:"pageSize,omitempty"`
	TotalItems int64                        `json:"totalItems,omitempty"`
	Data       *SamlMetadataInformationDto  `json:"data,omitempty"`
	Items      []SamlMetadataInformationDto `json:"items,omitempty"`
}
