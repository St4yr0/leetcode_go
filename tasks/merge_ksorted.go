package tasks

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// solution is to devide lists in to pairs and merge them
func mergeKLists(lists []*ListNode) *ListNode {
	var resNode *ListNode
	if len(lists) > 0 {
		resNode = lists[0]
		for i := 1; i < len(lists); i++ {
			resNode = mergeTwoLists(resNode, lists[i])
		}
	}
	return resNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var resNode *ListNode

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		resNode = list1
		resNode.Next = mergeTwoLists(list1.Next, list2)
	} else {
		resNode = list2
		resNode.Next = mergeTwoLists(list1, list2.Next)
	}
	return resNode

}
