package types

import "context"

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

type Request struct {
	HttpMethod string            `json:"httpMethod"`
	Body       []byte            `json:"body"`
	Query      map[string]string `json:"queryStringParameters"`
	Headers    map[string]string `json:"headers"`
	IsBase64   bool              `json:"IsBase64Encoded"`
}

func NewResponse(code int, body interface{}) *Response {
	return &Response{
		StatusCode: code,
		Body:       body,
	}
}

type HandlerFunc func(ctx context.Context, body interface{}) (*Response, error)

const RoutingTarget = "target"

// Routing paths
var PutFile = "put_file"
var PseudoDelete = "pseudo_delete"
var Delete = "delete"

var Targets = []string{PutFile, PseudoDelete, Delete}
