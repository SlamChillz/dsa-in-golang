package singly

import "testing"

func TestPrintListWithNilHead(t *testing.T) {
	var head *List
	r := PrintList(head)
	if r != nil {
		t.Fatalf(`PrintList(nil) = %v, want <nil>`, r)
	}
}

func TestPrintListWithNoneNilHead(t *testing.T) {
	next := NewList("Mendy", nil)
	head := NewList("Slam", next)
	r := PrintList(head)
	if r[head.len] != head.str || r[next.len] != next.str {
		t.Fatalf(`PrintList(head) = %#v, want map[4: "Slam" 5: "Mendy"]`, r)
	}
}