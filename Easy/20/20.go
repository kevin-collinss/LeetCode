package Easy

type Node struct {
	Val  rune
	Next *Node
}

type Stack struct {
	head *Node
}

func isValid(s string) bool {
	stack := &Stack{nil}

	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, r := range s {
		if r == '{' || r == '(' || r == '[' {
			stack.push(r)
		} else if r == '}' || r == ')' || r == ']' {
			if stack.head == nil {
				return false
			}
			top := stack.peek()
			if bracketPairs[r] == top {
				stack.pop()
			} else {
				return false
			}
		}
	}
	if stack.head == nil {
		return true
	}
	return false
}

func (stack *Stack) push(r rune) {
	newNode := &Node{Val: r, Next: stack.head}
	stack.head = newNode
}

func (stack *Stack) pop() rune {
	r := stack.head.Val
	stack.head = stack.head.Next
	return r
}

func (stack *Stack) peek() rune {
	return stack.head.Val
}
