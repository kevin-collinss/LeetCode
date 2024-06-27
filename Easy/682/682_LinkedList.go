package Easy

import (
	"strconv"
)

type Stack struct {
	head *ListNode
}

func calPoints(operations []string) int {
	stack := Stack{nil}

	for _, val := range operations {
		if num, err := strconv.Atoi(val); err == nil {
			stack.push(num)
		} else if val == "+" {
			prev := stack.pop()
			prev2 := stack.peek()
			stack.push(prev)
			stack.push(prev + prev2)
		} else if val == "D" {
			x := stack.peek()
			stack.push(x * 2)
		} else if val == "C" {
			stack.pop()
		}
	}

	total := 0
	for stack.head != nil {
		total += stack.head.Val
		stack.head = stack.head.Next
	}
	return total
}

func (stack *Stack) push(val int) {
	newNode := &ListNode{Val: val, Next: stack.head}
	stack.head = newNode
}

func (stack *Stack) pop() int {
	x := stack.head.Val
	stack.head = stack.head.Next
	return x
}

func (stack *Stack) peek() int {
	return stack.head.Val
}
