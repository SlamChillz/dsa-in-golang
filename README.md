# Data Structures in Golang

The project focuses on implementing standard data structures in Golang

## [Linked List](./linked-list)

A linked list is a linear data structure, in which the elements are not stored at contiguous memory locations. 

#### [Singly Linked List](./linked-list/singly)

**Data Structure:**
```
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
```