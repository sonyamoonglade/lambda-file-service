package headers

import (
	"errors"
	"log"
)

var XFileNameMissing = errors.New("x-filename header is missing")
var XDestinationMissing = errors.New("x-destination header is missing")
var XContentTypeMissing = errors.New("content-type header is missing")

type BaseHeaders struct {
	Filename    string
	Destination string
	ContentType string
}

//Yandex provides upper-cased headers
const XFileName = "X-Filename"
const XDestination = "X-Destination"
const XContentType = "X-Content-Type"

func FromRequest(logger *log.Logger, h map[string]string) (*BaseHeaders, error) {

	logger.Printf("headers: %v\n", h)

	filename, ok := h[XFileName]
	logger.Printf("filename: %s, ok: %t\n", filename, ok)
	if !ok {
		return nil, XFileNameMissing
	}

	destination, ok := h[XDestination]
	logger.Printf("destination: %s, ok: %t\n", destination, ok)
	if !ok {
		return nil, XDestinationMissing
	}

	contentType, ok := h[XContentType]
	logger.Printf("x-contentType: %s, ok: %t\n", contentType, ok)
	if !ok {
		return nil, XContentTypeMissing
	}

	return &BaseHeaders{
		Filename:    filename,
		Destination: destination,
		ContentType: contentType,
	}, nil

}
