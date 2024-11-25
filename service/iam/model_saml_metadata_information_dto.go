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

type SamlMetadataInformationDto struct {
	Id          string    `json:"id,omitempty"`
	DomainId    string    `json:"domainId,omitempty"`
	Name        string    `json:"name,omitempty"`
	EntityId    string    `json:"entityId,omitempty"`
	FileName    string    `json:"fileName,omitempty"`
	FileData    []string  `json:"fileData,omitempty"`
	CreatedDate time.Time `json:"createdDate,omitempty"`
	Active      bool      `json:"active,omitempty"`
	IdpMetadata bool      `json:"idpMetadata,omitempty"`
}