/*
 * VTT API
 *
 * API
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package iam

type CreateAccountActivityLogRequest struct {
	Url                       string `json:"url,omitempty"`
	Method                    string `json:"method,omitempty"`
	RequestBodyJson           string `json:"request_body_json,omitempty"`
	ResponseBodyJson          string `json:"response_body_json,omitempty"`
	HttpStatus                int32  `json:"http_status,omitempty"`
	HttpMessage               string `json:"http_message,omitempty"`
	RequestedTime             string `json:"requested_time,omitempty"`
	RespondedTime             string `json:"responded_time,omitempty"`
	RequestProcessingDuration int64  `json:"request_processing_duration,omitempty"`
	ThreadId                  string `json:"thread_id,omitempty"`
	ThreadName                string `json:"thread_name,omitempty"`
	ClientIpAddress           string `json:"client_ip_address,omitempty"`
	CurrentNodeIpAddress      string `json:"current_node_ip_address,omitempty"`
}
