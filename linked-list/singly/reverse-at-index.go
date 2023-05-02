package singly

/**
 * ReverseList: inserts a new node at given index
 * @head: pointer to the head node of a linked list
 * Returns: head of the reversed list node
*/
func ReverseList(head *List) *List {
	var next *List
	var nextNext *List
	if head == nil {
		return nil
	}
	next = head.next
	head.next = nil
	for next != nil {
		nextNext = next.next
		next.next = head
		head = next
		next = nextNext
	}
	return head
}