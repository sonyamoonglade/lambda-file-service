package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Testing struct {
	v   string
	exp error
}

func TestValidation(test *testing.T) {

	tt := []Testing{
		{v: "put_file", exp: nil},
		{v: "pseudo_delete", exp: nil},
		{v: "delete", exp: nil},
		{v: "deletee", exp: InvalidTarget},
		{v: "", exp: InvalidTarget},
		{v: "5", exp: InvalidTarget},
		{v: ":)", exp: InvalidTarget},
		{v: "*", exp: InvalidTarget},
		{v: "?!@?#>!?@#>!@?#>@!?>#!@?>#?@!#@!?#@!?#>@!?#>@!?>#@!?#>@!?#>", exp: InvalidTarget},
	}

	for _, t := range tt {
		actual := ValidateTarget(t.v)
		assert.Equal(test, t.exp, actual)
	}

}
