package transport

import (
	"errors"
)

var InvalidTarget = errors.New("invalid target")

func validateTarget(t string) error {

	hit := 0
	for _, tar := range Targets {
		if tar == t {
			hit += 1
		}
	}

	if hit == 0 {
		return InvalidTarget
	}

	return nil
}
