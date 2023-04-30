package singly

/**
 * AddNode: adds a new node to the head of the list
 * @head: head of the linked list
 * @name: str field of the new node head node
 * Returns: the new head node
 */
 func AddNode(head *List, name string) *List {
 	if name == "" {
 		return head
 	}
 	head = NewList(name, head)
 	return head
 }