package headers

import (
	"errors"
	"github.com/sonyamoonglade/lambda-file-service/pkg/types"
	"log"
)

var MissingXFileName = errors.New("x-filename header is missing")
var MissingXDestination = errors.New("x-destination header is missing")
var MissingXContentType = errors.New("content-type header is missing")
var AskForInvalidHeaders = errors.New("invalid headers asked")
var MissingXRoot = errors.New("x-root header is missing")

type BaseHeaders struct {
	Filename    string
	Destination string
	ContentType string
	Root        types.Root
}

type Provider interface {
	GetSpecific(h map[string]string, spec []string) (*BaseHeaders, error)
}

type provider struct {
	logger *log.Logger
}

func NewHeaderProvider(logger *log.Logger) Provider {
	return &provider{
		logger: logger,
	}
}

//Yandex provides upper-cased headers
const XFileName = "X-Filename"
const XDestination = "X-Destination"
const XContentType = "X-Content-Type"
const XRoot = "X-Root"

func (p *provider) GetSpecific(h map[string]string, spec []string) (*BaseHeaders, error) {

	b := &BaseHeaders{}

	for _, s := range spec {
		v, ok := h[s]
		switch s {
		case XFileName:
			if !ok {
				return nil, MissingXFileName
			}
			b.Filename = v
			break
		case XDestination:
			if !ok {
				return nil, MissingXDestination
			}
			b.Destination = v
			break
		case XContentType:
			if !ok {
				return nil, MissingXContentType
			}
			b.ContentType = v
			break
		case XRoot:
			if !ok {
				return nil, MissingXRoot
			}
			b.Root = types.Root(v)
			break
		default:
			return nil, AskForInvalidHeaders
		}

	}
	return b, nil
}
