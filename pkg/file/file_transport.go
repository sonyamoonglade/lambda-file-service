package file

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/sonyamoonglade/lambda-file-service"
	dto "github.com/sonyamoonglade/lambda-file-service/pkg/file/dto"
	"github.com/sonyamoonglade/lambda-file-service/pkg/headers"
	"github.com/sonyamoonglade/lambda-file-service/pkg/validation"
)

type Transport interface {
	Router(ctx context.Context, input []byte) (*lambda.Response, error)
	PutFile(ctx context.Context, r lambda.Request) (*lambda.Response, error)
	PseudoDelete(ctx context.Context, r lambda.Request) (*lambda.Response, error)
	Delete(ctx context.Context, r lambda.Request) (*lambda.Response, error)
}

type transport struct {
	service Service
}

func NewTransport(service Service) Transport {
	return &transport{service: service}
}

func (t *transport) PutFile(ctx context.Context, r lambda.Request) (*lambda.Response, error) {

	var inp dto.PutFileDto

	h, err := headers.FromRequest(r.Headers)
	if err != nil {
		return nil, err
	}

	inp.Destination = h.Destination
	inp.ContentType = h.ContentType
	inp.Filename = h.Filename
	inp.Bytes = r.Body

	out, err := t.service.Put(ctx, inp)
	if err != nil {
		return nil, err
	}

	return &lambda.Response{
		StatusCode: 201,
		Body:       out,
	}, nil

}

func (t *transport) PseudoDelete(ctx context.Context, r lambda.Request) (*lambda.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (t *transport) Delete(ctx context.Context, r lambda.Request) (*lambda.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (t *transport) Router(ctx context.Context, input []byte) (*lambda.Response, error) {

	var req lambda.Request

	err := json.Unmarshal(input, &req)
	if err != nil {
		return nil, err
	}

	target, ok := req.Query[lambda.RoutingTarget]
	//No specified target provided
	if !ok {
		return nil, errors.New("empty target")
	}

	err = validation.ValidateTarget(target)
	if err != nil {
		return nil, err
	}

	switch target {
	case lambda.PutFile:
		return t.PutFile(ctx, req)
	case lambda.PseudoDelete:
		return t.PseudoDelete(ctx, req)
	default: //delete
		return t.Delete(ctx, req)
	}
}
