package validation

import (
	"errors"
	"github.com/sonyamoonglade/lambda-file-service"
)

var InvalidTarget = errors.New("invalid target")

func ValidateTarget(t string) (error) {

	hit := 0
	for _, tar := range lambda.Targets {
		if tar == t {
			hit += 1
		}
	}

	if hit == 0 {
		return InvalidTarget
	}

	return nil
}
