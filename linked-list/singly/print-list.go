package singly

/**
 * PrintList: returns a map of str length as key and str as value
 * @head: pointer to the head node of the list
 * Returns: map[len(str)]str
*/
func PrintList(head *List) (r map[uint64]string) {
	h := head
	if h == nil {
		return nil
	}
	r = map[uint64]string{}
	for h != nil {
		r[h.len] = h.str
		h = h.next
	}
	return
}