package transport

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		actual := validateTarget(t.v)
		assert.Equal(test, t.exp, actual)
	}

}
