package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestItems struct {
	input    string
	expected string
}

var expectedResult string = "*simple test*"

var testItems = []TestItems{
	{"simple test", expectedResult},
	{" simple test ", expectedResult},
	{"", "**"},
}

func TestFormatText(t *testing.T) {

	assert := assert.New(t)

	for _, test := range testItems {

		assert.Equal(FormatText(test.input), test.expected)

		// if test.expected != FormatText(test.input) {
		// 	t.Errorf("Expected simple test '%s' to be '%s'", test.input, test.expected)
		// }
	}

}
