package singly

/**
 * struct List - singly linked list
 * @str: string
 * @len: length of string
 * @next: points to the next node
 * 
 * Description: singly linked list node structure
*/
type List struct {
	str string
	len uint64
	next *List
}

func NewList(str string, next *List) *List {
    len := uint64(len(str))
    if str == "" {
        return nil
    }
    return &List{
        str: str,
        len: len,
        next: next,
    }
}