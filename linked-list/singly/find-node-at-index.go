package singly

/**
 * FindNodeAtIndex: finds the node at given index
 * @head: pointer to the head node of a linked list
 * @index: index of the node, starting at 0
 * Returns: the node at index
*/
func FindNodeAtIndex(head *List, index uint64) *List {
	if head == nil {
		return head
	}
	position := uint64(0)
	for head != nil && position < index {
		position += 1
		head = head.next
	}
	return head
}