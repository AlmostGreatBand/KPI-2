package lab2

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)
type OutTest struct {
	called bool
}

func (o *OutTest) Write(p []byte) (n int, err error) {
	o.called = true
	return 0, nil
}

//Function that tests writing to output (stab)
func TestComputeHandler_Compute(t *testing.T) {
	expression := " 4 2 - 3 * 5 + "
	input := strings.NewReader(expression)
	output := OutTest{}

	handler := ComputeHandler{
		Input: input,
		Output: &output,
	}

	handler.Compute()

	assert.Equal(t,true, output.called)
}

//Functions that tests syntax errors
func TestComputeHandler_ComputeFailed(t *testing.T) {
	expressionFail := " 4 m - 3 * 5 + "
	input := strings.NewReader(expressionFail)
	output := OutTest{}

	handler := ComputeHandler{
		Input: input,
		Output: &output,
	}

	err := handler.Compute()

	assert.NotNil(t, err)

}