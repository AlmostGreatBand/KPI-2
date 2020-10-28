package lab2

import (
	"fmt"
	"io"
	"io/ioutil"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression.txt from input and write the computed result to the output.
type ComputeHandler struct {
	Input io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	data, err := ioutil.ReadAll(ch.Input)
	if err != nil {
		return err
	}

	res, err := CalculatePostfix(string(data))

	if err != nil {
		return err
	}

	out := []byte(fmt.Sprintf("%f", res))
	_, err = ch.Output.Write(out)

	return err
}
