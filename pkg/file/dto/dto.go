package dto

type PutFileDto struct {
	Filename    string //e.g. myfile.png
	Destination string //e.g. /static/images
	Bytes       []byte
	ContentType string //e.g. image/png
}

type DeleteFileDto struct {
	Filename    string
	Destination string
}
