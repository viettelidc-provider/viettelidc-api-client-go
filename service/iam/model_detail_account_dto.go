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

type DetailAccountDto struct {
	Id                  string    `json:"id,omitempty"`
	Email               string    `json:"email,omitempty"`
	Username            string    `json:"username,omitempty"`
	IsActive            bool      `json:"is_active,omitempty"`
	LastLoginTime       time.Time `json:"last_login_time,omitempty"`
	CreatedDate         time.Time `json:"created_date,omitempty"`
	LastModifiedDate    time.Time `json:"last_modified_date,omitempty"`
	IsUsingMfa          bool      `json:"is_using_mfa,omitempty"`
	IsLocked            bool      `json:"is_locked,omitempty"`
	LastPasswordChanged time.Time `json:"last_password_changed,omitempty"`
	IsPasswordExpired   bool      `json:"is_password_expired,omitempty"`
}
