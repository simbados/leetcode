package main

import (
	"slices"
	"strconv"
)

var operators = []string{"/", "*", "+", "-"}

func evalRPN(tokens []string) int {
	stack := StackImpl{values: []int{}}
	for _, token := range tokens {
		if isOperator(token) {
			valueTwo := stack.Pop()
			valueOne := stack.Pop()
			stack.Push(doOperation(valueOne, valueTwo, token))
		} else {
			converted, _ := strconv.Atoi(token)
			stack.Push(converted)
		}
	}
	return stack.Pop()
}

func doOperation(v1 int, v2 int, token string) int {
	if token == "/" {
		return v1 / v2
	} else if token == "*" {
		return v1 * v2
	} else if token == "+" {
		return v1 + v2
	} else {
		return v1 - v2
	}
}

func isOperator(token string) bool {
	return slices.Contains(operators, token)
}

type StackImpl struct {
	values []int
}

func (receiver *StackImpl) Push(val int) {
	receiver.values = append(receiver.values, val)
}

func (receiver *StackImpl) Pop() int {
	if len(receiver.values) > 0 {
		value := receiver.values[len(receiver.values)-1]
		receiver.values = receiver.values[0 : len(receiver.values)-1]
		return value
	}
	return 0
}
