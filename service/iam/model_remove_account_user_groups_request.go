/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type RemoveAccountUserGroupsRequest struct {
	AccountId string   `json:"account_id,omitempty"`
	GroupIds  []string `json:"group_ids,omitempty"`
}
