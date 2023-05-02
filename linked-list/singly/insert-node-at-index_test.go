package singly

import "testing"


func TestInsertAtFirstIndex(t *testing.T) {
	var head *List
	seed := []string{"Mendy", "Abeeb", "Slam", "Samuel"}
	for _, name := range seed {
		head = AddEndNode(head, name)
	}
	node := InsertNodeAtIndex(head, "uwemedimo", 0)
	if node.next != head {
		t.Fatalf(`Expected head.next == %+v, got %+v`, head, node.next)
	}
}

func TestInsertAtLastIndex(t *testing.T) {
	var head *List
	seed := []string{"Mendy", "Abeeb", "Slam", "Samuel"}
	for _, name := range seed {
		head = AddEndNode(head, name)
	}
	node := InsertNodeAtIndex(head, "uwemedimo", uint64(len(seed)) - 1)
	nodeAtPrevLast := FindNodeAtIndex(head, uint64(len(seed)) - 1)
	if node != nodeAtPrevLast {
		t.Fatalf(`Expected %+v, got %+v`, nodeAtPrevLast, node)
	}
}

func TestInsertAtMiddleIndex(t *testing.T) {
	var head *List
	seed := []string{"Mendy", "Abeeb", "Slam", "Samuel"}
	for _, name := range seed {
		head = AddEndNode(head, name)
	}
	mid := uint64(len(seed)/2)
	node := InsertNodeAtIndex(head, "uwemedimo", mid)
	nodeAtPrevLast := FindNodeAtIndex(head, mid)
	if node != nodeAtPrevLast {
		t.Fatalf(`Expected %+v, got %+v`, nodeAtPrevLast, node)
	}
}