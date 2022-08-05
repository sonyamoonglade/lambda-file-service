package lambda_file_service

import "context"

type LambdaResponse struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

func NewLambdaResponse(code int, body interface{}) *LambdaResponse {
	return &LambdaResponse{
		StatusCode: code,
		Body:       body,
	}
}

type LambdaHandlerFunc func(ctx context.Context) (*LambdaResponse, error)
