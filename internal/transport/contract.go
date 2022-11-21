package transport

import "context"

const RoutingTarget = "target"

// Routing paths
var PutFile = "put_file"
var PseudoDelete = "pseudo_delete"
var Delete = "delete"

var Targets = []string{PutFile, PseudoDelete, Delete}

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

type Request struct {
	HttpMethod string            `json:"httpMethod,omitempty"`
	Body       []byte            `json:"body,omitempty"`
	Query      map[string]string `json:"queryStringParameters,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	IsBase64   bool              `json:"IsBase64Encoded,omitempty"`
}

func NewResponse(code int, body interface{}) *Response {
	return &Response{
		StatusCode: code,
		Body:       body,
	}
}

type HandlerFunc func(ctx context.Context, body interface{}) (*Response, error)
