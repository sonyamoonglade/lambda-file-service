package dto

type PutFileDto struct {
	Filename    string //e.g. myfile.png
	Destination string //e.g. /static/images
	Bytes       []byte
	ContentType string //e.g. image/png
}

//Pseudo stands for delete files that have same root but leave the latest one alive.
type PseudoDeleteFileDto struct {
	Root        string
	Destination string
}
