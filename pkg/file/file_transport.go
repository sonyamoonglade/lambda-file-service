package file

import (
	"context"
	"encoding/json"
	"errors"
	dto "github.com/sonyamoonglade/lambda-file-service/pkg/file/dto"
	"github.com/sonyamoonglade/lambda-file-service/pkg/headers"
	"github.com/sonyamoonglade/lambda-file-service/pkg/types"
	"github.com/sonyamoonglade/lambda-file-service/pkg/validation"
	"log"
)

type Transport interface {
	Router(ctx context.Context, input []byte) (*types.Response, error)
	PutFile(ctx context.Context, r types.Request) (*types.Response, error)
	PseudoDelete(ctx context.Context, r types.Request) (*types.Response, error)
	Delete(ctx context.Context, r types.Request) (*types.Response, error)
}

type transport struct {
	service        Service
	logger         *log.Logger
	headerProvider headers.Provider
}

func NewTransport(logger *log.Logger, service Service) Transport {
	return &transport{service: service, logger: logger, headerProvider: headers.NewHeaderProvider(logger)}
}

func (t *transport) Router(ctx context.Context, input []byte) (*types.Response, error) {

	var req types.Request

	err := json.Unmarshal(input, &req)
	if err != nil {
		return nil, err
	}

	target, ok := req.Query[types.RoutingTarget]
	//No specified target provided
	if !ok {
		return nil, errors.New("empty target")
	}

	err = validation.ValidateTarget(target)
	if err != nil {
		return nil, err
	}

	switch target {
	case types.PutFile:
		return t.PutFile(ctx, req)
	case types.PseudoDelete:
		return t.PseudoDelete(ctx, req)
	default: //delete
		return t.Delete(ctx, req)
	}
}

func (t *transport) PutFile(ctx context.Context, r types.Request) (*types.Response, error) {

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

	return &types.Response{
		StatusCode: 201,
		Body:       out,
	}, nil

}

func (t *transport) PseudoDelete(ctx context.Context, r types.Request) (*types.Response, error) {

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

	return &types.Response{
		StatusCode: 200,
		Body:       nil,
	}, nil

}

func (t *transport) Delete(ctx context.Context, r types.Request) (*types.Response, error) {
	//TODO implement me
	panic("implement me")
}
