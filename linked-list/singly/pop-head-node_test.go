package singly

import "testing"

func TestPopHeadNode(t *testing.T) {
	var head *List
	seed := []string{"Mendy", "Abeeb", "Slam", "Samuel"}
	for _, name := range seed {
		head = AddEndNode(head, name)
	}
	for _, v := range seed {
		oldHead, newHead := PopHeadNode(head)
		if oldHead.str != v {
			t.Fatalf(`Expected oldHead.str == %v, got %v`, v, oldHead.str)
		}
		if newHead != oldHead.next {
			t.Fatalf(`Expected newHead == %p, got %p`,oldHead.next, newHead)
		}
		head = newHead
	}
}