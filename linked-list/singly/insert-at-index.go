package singly

/**
 * InsertNodeAtIndex: inserts a new node at given index
 * @head: pointer to the head node of a linked list
 * @name: data for the new node
 * @index: index of the node, starting at 0
 * Returns: the newly inserted node or nil if not possible
*/
func InsertNodeAtIndex(head *List, name string, index uint64) *List {
	var prev *List
	var newNode *List
	if head == nil || len(name) == 0 {
		return new(List)
	}
	position := uint64(0)
	for head != nil && position < index {
		prev = head
		position += 1
		head = head.next
	}
	if head != nil {
		newNode = NewList(name, head)
		if prev != nil {
			prev.next = newNode
		}
	}
	return newNode
}