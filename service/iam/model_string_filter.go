/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type StringFilter struct {
	Equals         string   `json:"equals,omitempty"`
	NotEquals      string   `json:"notEquals,omitempty"`
	Specified      bool     `json:"specified,omitempty"`
	In             []string `json:"in,omitempty"`
	NotIn          []string `json:"notIn,omitempty"`
	Contains       string   `json:"contains,omitempty"`
	DoesNotContain string   `json:"doesNotContain,omitempty"`
}
