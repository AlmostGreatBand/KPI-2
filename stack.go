package lab2

import "errors"

type Node struct {
	value float64
	next  *Node
}

type Stack struct {
	top    *Node
	length int
}

func (stack *Stack) push(node *Node) {
	stack.length++
	node.next = stack.top
	stack.top = node
}

func (stack *Stack) pop() (*Node, error) {
	if stack.length == 0 {
		return nil, errors.New("empty stack")
	}
	stack.length--
	node := stack.top
	stack.top = node.next
	return node, nil
}
