package validation

import (
	"errors"
	"github.com/sonyamoonglade/lambda-file-service/pkg/types"
)

var InvalidTarget = errors.New("invalid target")

func ValidateTarget(t string) error {

	hit := 0
	for _, tar := range types.Targets {
		if tar == t {
			hit += 1
		}
	}

	if hit == 0 {
		return InvalidTarget
	}

	return nil
}
