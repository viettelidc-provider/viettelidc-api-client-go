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
	"os"
)

type UploadSamlMetadataFileRequest struct {
	Name     string    `json:"name,omitempty"`
	File     **os.File `json:"file,omitempty"`
	DomainId string    `json:"domainId,omitempty"`
}