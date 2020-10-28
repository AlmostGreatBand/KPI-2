package lab2

import (
	"errors"
	"math"
	"strconv"
	s "strings"
)

func CalculatePostfix(input string) (float64, error) {
	stack := &Stack{nil, 0}
	operators := Filter(s.Split(input, " "), func(c string) bool {
		return c != "" && c != " "
	})

	for _, v := range operators {
		if num, e := strconv.ParseFloat(v, 64); e == nil {
			stack.push(&Node{num, nil})
		} else {
			left, _ := stack.pop()
			right, err := stack.pop()
			if err != nil {
				return 0.0, err
			}
			switch v {
			case "-":
				node := &Node{right.value - left.value, nil}
				stack.push(node)
			case "+":
				node := &Node{right.value + left.value, nil}
				stack.push(node)
			case "*":
				node := &Node{right.value * left.value, nil}
				stack.push(node)
			case "/":
				if left.value == 0 {
					return 0.0, errors.New("denominator cannot be 0")
				}
				node := &Node{right.value / left.value, nil}
				stack.push(node)
			case "^":
				node := &Node{math.Pow(right.value, left.value), nil}
				stack.push(node)
			default:
				return 0.0, errors.New("unexpected type")
			}
		}
	}
	if res, err := stack.pop(); err == nil {
		if stack.length != 0 {
			return 0.0, errors.New("stack is not empty; mistake in expresion")
		}
		return res.value, nil
	} else {
		return 0.0, err
	}
}
