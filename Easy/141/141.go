package Easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {

		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}
	return false
}
