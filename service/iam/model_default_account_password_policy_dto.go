/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type DefaultAccountPasswordPolicyDto struct {
	MinimumLength                 int32 `json:"minimum_length,omitempty"`
	IsRequiredUppercaseCharacters bool  `json:"is_required_uppercase_characters,omitempty"`
	IsRequiredLowercaseCharacters bool  `json:"is_required_lowercase_characters,omitempty"`
	IsRequiredNumberCharacters    bool  `json:"is_required_number_characters,omitempty"`
	IsRequiredSymbolCharacters    bool  `json:"is_required_symbol_characters,omitempty"`
	IsExpired                     bool  `json:"is_expired,omitempty"`
	ExpiredPeriod                 int32 `json:"expired_period,omitempty"`
	IsPreventReused               bool  `json:"is_prevent_reused,omitempty"`
	PreventReusedPeriod           int32 `json:"prevent_reused_period,omitempty"`
	IsUsingLockPolicy             bool  `json:"is_using_lock_policy,omitempty"`
	AccountMaxLoginAttempts       int32 `json:"account_max_login_attempts,omitempty"`
	AccountLockoutDurations       int32 `json:"account_lockout_durations,omitempty"`
	IsAllowUserChangePassword     bool  `json:"is_allow_user_change_password,omitempty"`
}
