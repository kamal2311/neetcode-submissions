/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
    if list1 == nil {
        return list2
    }

    if list2 == nil {
        return list1
    }

    res := &ListNode{}

    if list1.Val <= list2.Val {        
        res.Val = list1.Val
        res.Next = mergeTwoLists(list1.Next, list2)
    } else {
        res.Val = list2.Val
        res.Next = mergeTwoLists(list1, list2.Next)
    }

    return res

}
