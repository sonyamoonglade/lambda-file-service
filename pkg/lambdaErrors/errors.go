package lambdaErrors

import "errors"

var UnableToDeleteFile = errors.New("unable to delete file due to it has not any root siblings")
