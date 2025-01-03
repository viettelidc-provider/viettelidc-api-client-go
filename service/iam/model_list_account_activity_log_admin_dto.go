/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type ListAccountActivityLogAdminDto struct {
	Id                        string `json:"id,omitempty"`
	Username                  string `json:"username,omitempty"`
	Email                     string `json:"email,omitempty"`
	Action                    string `json:"action,omitempty"`
	Url                       string `json:"url,omitempty"`
	Method                    string `json:"method,omitempty"`
	VttServiceName            string `json:"vtt_service_name,omitempty"`
	AccountId                 string `json:"account_id,omitempty"`
	DomainId                  string `json:"domain_id,omitempty"`
	DomainEmail               string `json:"domain_email,omitempty"`
	AccountRole               string `json:"account_role,omitempty"`
	IsRootAccount             bool   `json:"is_root_account,omitempty"`
	ResourceType              string `json:"resource_type,omitempty"`
	EndpointDescription       string `json:"endpoint_description,omitempty"`
	RequestBodyJson           string `json:"request_body_json,omitempty"`
	ResponseBodyJson          string `json:"response_body_json,omitempty"`
	HttpStatus                int32  `json:"http_status,omitempty"`
	HttpMessage               string `json:"http_message,omitempty"`
	RequestedTime             string `json:"requested_time,omitempty"`
	RespondedTime             string `json:"responded_time,omitempty"`
	RequestProcessingDuration int32  `json:"request_processing_duration,omitempty"`
	ThreadId                  string `json:"thread_id,omitempty"`
	ThreadName                string `json:"thread_name,omitempty"`
	ClientIpAddress           string `json:"client_ip_address,omitempty"`
	CurrentNodeIpAddress      string `json:"current_node_ip_address,omitempty"`
}
