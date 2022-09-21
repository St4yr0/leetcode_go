package tasks

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
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

// // LinkedList represents a linked list
// type LinkedList struct {
// 	Head *ListNode
// 	Len  int
// }

// // Insert inserts new node at the end of  from linked list
// func (l *LinkedList) Insert(val int) {
// 	n := ListNode{}
// 	n.Val = val
// 	if l.Len == 0 {
// 		l.Head = &n
// 		l.Len++
// 		return
// 	}
// 	ptr := l.Head
// 	for i := 0; i < l.Len; i++ {
// 		if ptr.Next == nil {
// 			ptr.Next = &n
// 			l.Len++
// 			return
// 		}
// 		ptr = ptr.Next
// 	}
// }

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	merged_list := LinkedList{}

// 	for {
// 		if list1 == nil {
// 			for list2 != nil {
// 				merged_list.Insert(list2.Val)
// 				list2 = list2.Next
// 			}
// 			break
// 		} else if list2 == nil {
// 			for list1 != nil {
// 				merged_list.Insert(list1.Val)
// 				list1 = list1.Next
// 			}
// 			break
// 		}
// 		if list1.Val < list2.Val {
// 			merged_list.Insert(list1.Val)
// 			list1 = list1.Next
// 			continue
// 		} else if list2.Val < list1.Val {
// 			merged_list.Insert(list2.Val)
// 			list2 = list2.Next
// 			continue
// 		} else if list2.Val == list1.Val {
// 			merged_list.Insert(list1.Val)
// 			merged_list.Insert(list2.Val)
// 			list1 = list1.Next
// 			list2 = list2.Next
// 			continue
// 		}
// 		break

// 	}
// 	return merged_list.Head
// }
