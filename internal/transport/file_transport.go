package transport

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/sonyamoonglade/lambda-file-service/internal/headers"
	"github.com/sonyamoonglade/lambda-file-service/internal/lambda_errors"
	"github.com/sonyamoonglade/lambda-file-service/internal/service"
	"github.com/sonyamoonglade/lambda-file-service/internal/transport/dto"
)

type Transport interface {
	Router(ctx context.Context, input []byte) (*Response, error)
	PutFile(ctx context.Context, r *Request) (*Response, error)
	PseudoDelete(ctx context.Context, r *Request) (*Response, error)
	Delete(ctx context.Context, r *Request) (*Response, error)
}

type transport struct {
	service        service.FileService
	logger         *log.Logger
	headerProvider headers.Provider
}

func NewTransport(logger *log.Logger, service service.FileService) Transport {
	return &transport{service: service, logger: logger, headerProvider: headers.NewHeaderProvider(logger)}
}

func (t *transport) Router(ctx context.Context, input []byte) (*Response, error) {
	var req *Request

	err := json.Unmarshal(input, &req)
	if err != nil {
		return nil, err
	}

	target, ok := req.Query[RoutingTarget]
	//No specified target provided
	if !ok {
		return nil, errors.New("empty target")
	}

	err = validateTarget(target)
	if err != nil {
		return nil, err
	}

	m := req.HttpMethod

	switch {
	case target == PutFile && m == http.MethodPost:
		return t.PutFile(ctx, req)
	case target == PseudoDelete && m == http.MethodPost:
		return t.PseudoDelete(ctx, req)
	case target == Delete && m == http.MethodPost:
		return t.Delete(ctx, req)
	default:
		return nil, lambda_errors.MethodOrTargetIsNotAllowed
	}
}

func (t *transport) PutFile(ctx context.Context, r *Request) (*Response, error) {

	var inp dto.PutFileDto

	//Which headers to require from headerProvider
	hspec := []string{headers.XDestination, headers.XFileName, headers.XContentType}

	h, err := t.headerProvider.GetSpecific(r.Headers, hspec)
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

	return &Response{
		StatusCode: 201,
		Body:       out,
	}, nil

}

func (t *transport) PseudoDelete(ctx context.Context, r *Request) (*Response, error) {

	var inp dto.DeleteFileDto

	//Which headers to require from headerProvider
	hspec := []string{headers.XRoot, headers.XDestination}
	h, err := t.headerProvider.GetSpecific(r.Headers, hspec)
	if err != nil {
		return nil, err
	}

	inp.Destination = h.Destination

	storage, err := t.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	//Assume only work with images through this tool.
	images := storage.Images

	//Find the latest file by .LastModified attr by root
	latest, err := t.service.FindOldestByRoot(images, h.Root)
	if err != nil {
		return nil, err
	}

	//Fulfill exact deleted filename after finding the oldest file
	inp.Filename = latest.Name

	err = t.service.Delete(ctx, inp)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: 200,
		Body:       nil,
	}, nil

}

func (t *transport) Delete(ctx context.Context, r *Request) (*Response, error) {

	var inp dto.DeleteFileDto

	hspec := []string{headers.XDestination, headers.XFileName}
	h, err := t.headerProvider.GetSpecific(r.Headers, hspec)
	if err != nil {
		return nil, err
	}

	inp.Destination = h.Destination
	inp.Filename = h.Filename

	err = t.service.Delete(ctx, inp)
	if err != nil {
		return nil, err
	}

	return NewResponse(http.StatusOK, nil), nil
}
