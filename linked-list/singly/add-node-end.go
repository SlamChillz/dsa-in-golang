package singly

/**
 * AddEndNode: adds a new node to the end of the list
 * @head: head of the linked list
 * @name: str field of the new node head node
 * Returns: the new head node
 */
 func AddEndNode(head *List, name string) *List {
    h := head
 	if name == "" {
 		return head
 	}
    if head == nil {
        return NewList(name, head)
    } else {
        for head.next != nil {
            head = head.next
        }
        head.next = NewList(name, nil)
    }
    return h
 }