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

type AccountSessionDto struct {
	Id             string      `json:"id,omitempty"`
	AccountId      string      `json:"account_id,omitempty"`
	TokenType      string      `json:"token_type,omitempty"`
	EncryptToken   string      `json:"encrypt_token,omitempty"`
	ExpiredTime    time.Time   `json:"expired_time,omitempty"`
	IsValid        bool        `json:"is_valid,omitempty"`
	BrowserName    string      `json:"browser_name,omitempty"`
	BrowserVersion string      `json:"browser_version,omitempty"`
	DeviceName     string      `json:"device_name,omitempty"`
	OsName         string      `json:"os_name,omitempty"`
	OsVersion      string      `json:"os_version,omitempty"`
	Country        string      `json:"country,omitempty"`
	City           string      `json:"city,omitempty"`
	FirstLoginTime time.Time   `json:"first_login_time,omitempty"`
	AccountDto     *AccountDto `json:"account_dto,omitempty"`
}
