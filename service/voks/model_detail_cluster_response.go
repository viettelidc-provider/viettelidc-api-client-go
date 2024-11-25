/*
 * vOKS IaC API
 *
 * This is required APIs to develop vOKS IaC
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package voks

type DetailClusterResponse struct {
	Id         int32  `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	VpcId      int32  `json:"vpcId,omitempty"`
	VpcNetwork string `json:"vpcNetwork,omitempty"`
	Version    string `json:"version,omitempty"`
	Mode       string `json:"mode,omitempty"`
	Status     string `json:"status,omitempty"`
	ApiAddress string `json:"apiAddress,omitempty"`
}