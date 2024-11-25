package viettelidc

import (
	"net/http"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "ctx " + string(c)
}

var (
	// ContextOAuth2 takes a oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextBasicAuth takes BasicAuth as authentication for the request.
	ContextBasicAuth = contextKey("basic")

	// ContextAPIKey takes an APIKey as authentication for the request
	ContextAPIKey = contextKey("apikey")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("access_token")

	ContextXAccessToken = contextKey("x_access_token")

	ContextId = contextKey("id")

	ContextCustomerId = contextKey("customer_id")

	ContextDomainId = contextKey("domain_id")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	Host          string            `json:"host,omitempty"`
	Scheme        string            `json:"scheme,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`

	Id           string `json:"id,omitempty"`
	DomainId     string `json:"domainId,omitempty"`
	CustomerId   string `json:"customerId,omitempty"`
	AccessToken  string `json:"accessToken,omitempty"`
	XAccessToken string `json:"x-access-token,omitempty"`

	HTTPClient *http.Client
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		BasePath:      "http://localhost:7202",
		DefaultHeader: make(map[string]string),
		UserAgent:     "viettelidc/cmp/iac",
	}
	return cfg
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}
