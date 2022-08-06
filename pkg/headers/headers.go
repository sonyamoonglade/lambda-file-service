package headers

import "errors"

var XFileNameMissing = errors.New("x-filename header is missing")
var XDestinationMissing = errors.New("x-destination header is missing")
var XContentTypeMissing = errors.New("content-type header is missing")

type BaseHeaders struct {
	Filename    string
	Destination string
	ContentType string
}

const XFileName = "x-filename"
const XDestination = "x-destination"
const XContentType = "x-content-type"

func FromRequest(h map[string]string) (*BaseHeaders, error) {

	filename, ok := h[XFileName]
	if !ok {
		return nil, XFileNameMissing
	}

	destination, ok := h[XDestination]
	if !ok {
		return nil, XDestinationMissing
	}

	contentType, ok := h[XContentType]
	if !ok {
		return nil, XContentTypeMissing
	}

	return &BaseHeaders{
		Filename:    filename,
		Destination: destination,
		ContentType: contentType,
	}, nil

}
