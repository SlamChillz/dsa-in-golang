package singly

import "testing"

func TestListLengthWithNilHead(t *testing.T) {
	var head *List
	length := ListLength(head)
	if length != 0 {
		t.Fatalf(`ListLength(nil) = %v, want nil`, length)
	}
}

func TestListLengthWithNoneNilHead(t *testing.T) {
	next := NewList("Mendy", nil)
	head := NewList("Slam", next)
	length := ListLength(head)
	if length != 2 {
		t.Fatalf(`ListLength(head) = %v, want 2`, length)
	}
}