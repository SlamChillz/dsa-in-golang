package singly

/**
 * PopHeadNode: deletes the head node from a linked list
 * @head: pointer to the head node of a linked list
 * Returns: previous head node and new head node
*/
func PopHeadNode(head *List) (List, *List) {
	if head == nil {
		return *new(List), head
	}
	headNode := *head
	head = head.next
	return headNode, head
}