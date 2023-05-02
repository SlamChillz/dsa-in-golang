package singly

import "testing"

func TestWithIndexOutOfRange(t *testing.T) {
	var head *List
	seed := []string{"Mendy", "Abeeb", "Slam", "Samuel"}
	for _, name := range seed {
		head = AddEndNode(head, name)
	}
	node := FindNodeAtIndex(head, 5)
	if node != nil {
		t.Fatalf(`Expected <nil>, got %p`, node)
	}
}

func TestWithIndexWithinRange(t *testing.T) {
	var head *List
	seed := []string{"Mendy", "Abeeb", "Slam", "Samuel"}
	for _, name := range seed {
		head = AddEndNode(head, name)
	}
	mid := uint64(len(seed)/2)
	node := FindNodeAtIndex(head, mid)
	if node.str != seed[mid] {
		t.Fatalf(`Expected %v got %v`, seed[mid], node.str)
	}
}