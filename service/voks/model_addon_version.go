/*
 * vOKS IaC API
 *
 * This is required APIs to develop vOKS IaC
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package voks

type AddonVersion struct {
	VersionId   string `json:"versionId,omitempty"`
	VersionName string `json:"versionName,omitempty"`
	Status      string `json:"status,omitempty"`
}