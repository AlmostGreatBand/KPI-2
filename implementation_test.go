package lab2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculatePostfix(t *testing.T) {
	res, err := CalculatePostfix(" 4 2 - 3 * 5 + ")
	if assert.Nil(t, err) {
		assert.Equal(t, 11.0, res)
	}

	res, err = CalculatePostfix(" 2 6 ^ 3 6 / +")
	if assert.Nil(t, err) {
		assert.Equal(t, 64.5, res)
	}

	res, err = CalculatePostfix(" 3 12 3 / 6 4 - ^ * ")
	if assert.Nil(t, err) {
		assert.Equal(t, 48.0, res)
	}

	res, err = CalculatePostfix(" 9 4 * 147 57 - 5 / - 35 7 / + 4 7 + 11 - - ")
	if assert.Nil(t, err) {
		assert.Equal(t, 23.0, res)
	}
}

func TestCalculatePostfixFailed(t *testing.T) {
	_, err := CalculatePostfix(" 9 + ")
	assert.NotNil(t, err)

	_, err = CalculatePostfix(" + ")
	assert.NotNil(t, err)

	_, err = CalculatePostfix(" 9 + ")
	assert.NotNil(t, err)

	_, err = CalculatePostfix("qwerty")
	assert.NotNil(t, err)

	_, err = CalculatePostfix("1 2")
	assert.NotNil(t, err)
}

//Example of calculating postfix
func ExampleCalculatePostfix() {
	res, _ := CalculatePostfix(" 2 2 +")
	fmt.Println(res)
	// Output:
	// 4
}
