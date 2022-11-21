package lambda_errors

import "errors"

var UnableToDeleteFile = errors.New("unable to delete file due to it has not any root siblings")
var MethodOrTargetIsNotAllowed = errors.New("http method or target is not allowed")
