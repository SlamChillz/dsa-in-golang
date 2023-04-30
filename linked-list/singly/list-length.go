package singly

/**
 * ListLength: calculates the length of a list
 * @head: head node of the list
 * Returns: The length of the list
*/
func ListLength(head *List) uint64 {
	var length uint64
	length = 0
	for head != nil {
		length += 1
		head = head.next
	}
	return length
}