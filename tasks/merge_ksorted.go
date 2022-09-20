package tasks

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// solution is to devide lists in to pairs and merge them
func mergeKLists(lists []*ListNode) *ListNode {
	merged_list := LinkedList{}
	dev := len(lists) / 2

	for {

	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	resNode := ListNode{}

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		resNode = *list1
		resNode.Next = mergeTwoLists(list1.Next, list2)
	} else {
		resNode = *list2
		resNode.Next = mergeTwoLists(list1, list2.Next)
	}
	return &resNode

}
