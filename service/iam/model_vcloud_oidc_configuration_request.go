/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type VcloudOidcConfigurationRequest struct {
	HostId                                       int32  `json:"host_id,omitempty"`
	HostIp                                       string `json:"host_ip,omitempty"`
	PublicIp                                     string `json:"public_ip,omitempty"`
	OrgId                                        string `json:"org_id,omitempty"`
	VdcOrgId                                     string `json:"vdc_org_id,omitempty"`
	ClientId                                     string `json:"client_id,omitempty"`
	ClientSecret                                 string `json:"client_secret,omitempty"`
	AtmAccountName                               string `json:"atm_account_name,omitempty"`
	IsUpdateRequest                              bool   `json:"is_update_request,omitempty"`
	OpenIdProviderWellknownConfigurationEndpoint string `json:"open_id_provider_wellknown_configuration_endpoint,omitempty"`
}