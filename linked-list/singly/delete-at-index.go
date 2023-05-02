package singly


/**
 * DeleteNodeAtIndex: deletes a new node at given index
 * @head: pointer to the head node of a linked list
 * @index: index of the node to be deleted, starting at 0
 * Returns: new head node and 1 if successful else -1
*/
func DeleteNodeAtIndex(head *List, index uint64) (*List, int) {
	var prev *List
	headNode := head
	if head == nil {
		return head, -1
	}
	position := uint64(0)
	for head != nil && position < index {
		prev = head
		position += 1
		head = head.next
	}
	if head != nil {
		next := head.next
		if next == nil {
			if prev != nil {
				prev.next = nil
				return headNode, 1
			}
			return nil, 1
		}
		*head = *next
		return headNode, 1
	}
	return headNode, -1
}