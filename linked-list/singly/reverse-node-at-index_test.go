package singly

import "testing"

func TestReverseListWithOneNode(t *testing.T) {
	var head *List
	seed := []string{"Mendy"}
	for _, name := range seed {
		head = AddEndNode(head, name)
	}
	newHead := ReverseList(head)
	if newHead != head {
		t.Fatalf(`Expected ReverseList(head) == %v, got %v`, head, newHead)
	}
}

func TestReverseListWithTwoNodes(t *testing.T) {
	var head *List
	seed := []string{"Mendy", "Abeeb"}
	for _, name := range seed {
		head = AddEndNode(head, name)
	}
	newHead := ReverseList(head)
	if newHead == head {
		t.Fatalf(`Expected ReverseList(head) == %v, got %v`, head.next, newHead)
	}
	if newHead.next != head {
		t.Fatalf(`Expected newHead.next == %v, got %v`, head, newHead.next)
	}
}

func TestReverseListWithSeveralNodes(t *testing.T) {
	var head *List
	seed := []string{"Mendy", "Abeeb", "Slam", "Samuel"}
	for _, name := range seed {
		head = AddNode(head, name)
	}
	prevHead := head
	newHead := ReverseList(head)
	if prevHead.next != nil {
		t.Fatalf(`Expected prevHead.next == <nil>, got %+v`, prevHead.next)
	}
	if prevHead == newHead {
		t.Fatalf(`Expected prevHead != newHead, got prevHead: %v == newHead: %v.`, prevHead, newHead)
	}
	for index, name := range seed {
		if newHead.str != name {
			t.Fatalf(`Expected newHead.str == %v at index %d, got %v`, name, index, newHead.str)
		}
		newHead = newHead.next
	}
}