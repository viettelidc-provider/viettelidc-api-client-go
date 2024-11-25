/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type AccountSessionCriteria struct {
	Name           *StringFilter      `json:"name,omitempty"`
	AccountId      *StringFilter      `json:"accountId,omitempty"`
	DeviceName     *StringFilter      `json:"deviceName,omitempty"`
	OsName         *StringFilter      `json:"osName,omitempty"`
	BrowserName    *StringFilter      `json:"browserName,omitempty"`
	Location       *StringFilter      `json:"location,omitempty"`
	FirstLoginTime *RangeFilterString `json:"firstLoginTime,omitempty"`
	IsValid        *BooleanFilter     `json:"isValid,omitempty"`
}