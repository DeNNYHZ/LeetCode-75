package Linked_List

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, current, next *ListNode
	current = head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
		head = prev
	}
	return head
}
