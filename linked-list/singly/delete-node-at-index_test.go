package singly

import "testing"


func TestDeleteAtFirstIndex(t *testing.T) {
	var head *List
	seed := []string{"Mendy", "Abeeb", "Slam", "Samuel"}
	for _, name := range seed {
		head = AddEndNode(head, name)
	}
	head, deleted := DeleteNodeAtIndex(head, 0)
	if deleted == -1 {
		t.Fatalf(`Expected deleted == %v, got %v`, 1, deleted)
	}
	if head.str != seed[1] {
		t.Fatalf(`Expected head.str == %+v, got %+v`, seed[1], head.str)
	}
}

func TestDeleteAtFirstIndexWithOneNode(t *testing.T) {
	var head *List
	seed := []string{"Mendy"}
	for _, name := range seed {
		head = AddEndNode(head, name)
	}
	head, deleted := DeleteNodeAtIndex(head, 0)
	if deleted == -1 {
		t.Fatalf(`Expected deleted == %v, got %v`, 1, deleted)
	}
	if head != nil {
		t.Fatalf(`Expected head == <nil>, got %+v`, head)
	}
}

func TestDeleteAtLastIndex(t *testing.T) {
	var head *List
	seed := []string{"Mendy", "Abeeb", "Slam", "Samuel"}
	for _, name := range seed {
		head = AddEndNode(head, name)
	}
	head, deleted := DeleteNodeAtIndex(head, uint64(len(seed)) - 1)
	nodeAtPrevLast := FindNodeAtIndex(head, uint64(len(seed)) - 1)
	if deleted == -1 {
		t.Fatalf(`Expected deleted == %v, got %v`, 1, deleted)
	}
	if nodeAtPrevLast != nil {
		t.Fatalf(`Expected <nil>, got %+v`, nodeAtPrevLast)
	}
}

func TestDeleteAtMiddleIndex(t *testing.T) {
	var head *List
	seed := []string{"Mendy", "Abeeb", "Slam", "Samuel"}
	for _, name := range seed {
		head = AddEndNode(head, name)
	}
	mid := uint64(len(seed)/2)
	head, deleted := DeleteNodeAtIndex(head, mid)
	after := FindNodeAtIndex(head, mid)
	if deleted == -1 {
		t.Fatalf(`Expected deleted == %v, got %v`, 1, deleted)
	}
	if after.str == seed[mid] {
		t.Fatalf(`Expected %+v != %+v`, after.str, seed[mid])
	}
}