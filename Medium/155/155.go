package Medium

type MinStack struct {
	head *ListNode
	min  *ListNode
}

func Constructor() MinStack {
	return MinStack{head: nil, min: nil}
}

func (this *MinStack) Push(val int) {
	newNode := &ListNode{Val: val, Next: this.head}
	this.head = newNode

	newMin := &ListNode{Val: val, Next: this.min}
	if this.min == nil || this.min.Val >= val {
		this.min = newMin
	}
}

func (this *MinStack) Pop() {
	if this.head.Val == this.min.Val {
		this.min = this.min.Next
	}
	this.head = this.head.Next
}

func (this *MinStack) Top() int {
	if this.head != nil {
		return this.head.Val
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if this.min != nil {
		return this.min.Val
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
