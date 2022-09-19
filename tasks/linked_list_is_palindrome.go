package tasks

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	stack := []int{}
	next := head

	for {
		if next == nil {
			break
		} else {
			stack = append(stack, next.Val)
			next = next.Next
		}
	}

	for len(stack) > 0 {
		n := len(stack) - 1
		if stack[n] != head.Val {
			return false
		}
		stack = stack[:n]
		head = head.Next
	}
	return true
}
