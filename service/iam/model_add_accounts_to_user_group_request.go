/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type AddAccountsToUserGroupRequest struct {
	AccountIds  []string `json:"account_ids,omitempty"`
	UserGroupId string   `json:"user_group_id,omitempty"`
}
