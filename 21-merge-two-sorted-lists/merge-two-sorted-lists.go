/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}   
    output := dummy
    for true {

        if list1 == nil && list2 == nil{
            break
        }

        if list1 == nil {
            output.Next = &ListNode{Val: list2.Val}
            list2 = list2.Next    
            output = output.Next    
            continue    
        }

        if list2 == nil {
            output.Next = &ListNode{Val: list1.Val}
            list1 = list1.Next 
            output = output.Next
            continue
        }

        if list1.Val < list2.Val {
            output.Next = &ListNode{Val: list1.Val}
            list1 = list1.Next 
            output = output.Next
            continue
        }else{
            output.Next = &ListNode{Val: list2.Val} 
            list2 = list2.Next   
            output = output.Next    
            continue    
        }
    }

    return dummy.Next
}