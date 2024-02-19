package main

type MinStack struct {
	arr []int
	min []int
}

func Constructor2() MinStack {
	stack := MinStack{}
	return stack
}

func (stack *MinStack) Push(val int) {
	stack.arr = append(stack.arr, val)
	if len(stack.min) == 0 || stack.min[len(stack.min)-1] >= val {
		stack.min = append(stack.min, val)
	}
}

func (stack *MinStack) Pop() {
	popElement := stack.arr[len(stack.arr)-1]
	stack.arr = stack.arr[0 : len(stack.arr)-1]
	if len(stack.min) > 0 && stack.min[len(stack.min)-1] == popElement {
		stack.min = stack.min[0 : len(stack.min)-1]
	}
}

func (stack *MinStack) Top() int {
	return stack.arr[len(stack.arr)-1]
}

func (stack *MinStack) GetMin() int {
	return stack.min[len(stack.min)-1]
}

func (stack *MinStack) Len() int {
	return len(stack.arr)
}
