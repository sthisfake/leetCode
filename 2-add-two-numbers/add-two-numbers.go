/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	result := dummyHead

	sizeListNode1 := 0
	sizeListNode2 := 0
	lenghtToItterate := 0

	currentListNode1 := l1
	currentListNode2 := l2

	for currentListNode1 != nil || currentListNode2 != nil {
		if currentListNode1 != nil {
			sizeListNode1++
			currentListNode1 = currentListNode1.Next
		}
		if currentListNode2 != nil {
			sizeListNode2++
			currentListNode2 = currentListNode2.Next
		}
	}
	if sizeListNode1 >= sizeListNode2 {
		lenghtToItterate = sizeListNode1
	} else {
		lenghtToItterate = sizeListNode2
	}

	var shouldAddOne bool
	for i := 0; i <= lenghtToItterate+1; i++ {
		if l1 == nil && l2 != nil {
			result.Val = l2.Val
			if shouldAddOne {
				result.Val++
				shouldAddOne = false
			}
			if result.Val > 9 {
				result.Val = result.Val - 10
				shouldAddOne = true
			}
			l2 = l2.Next
			if l2 != nil || shouldAddOne {
				result.Next = &ListNode{}
				result = result.Next
			}
		} else if l2 == nil && l1 != nil {
			result.Val = l1.Val
			if shouldAddOne {
				result.Val++
				shouldAddOne = false
			}
			if result.Val > 9 {
				result.Val = result.Val - 10
				shouldAddOne = true
			}
			l1 = l1.Next
			if l1 != nil || shouldAddOne {
				result.Next = &ListNode{}
				result = result.Next
			}
		} else {
			if l1 == nil && l2 == nil {
				if shouldAddOne {
					result.Val++
					shouldAddOne = false
				}
				break
			}
			result.Val = l1.Val + l2.Val
			if shouldAddOne {
				result.Val++
				shouldAddOne = false
			}
			if result.Val > 9 {
				result.Val = result.Val - 10
				shouldAddOne = true
			}
			l1 = l1.Next
			l2 = l2.Next
			if l1 != nil || l2 != nil || shouldAddOne {
				result.Next = &ListNode{}
				result = result.Next
			}

		}
	}

	return dummyHead
}
