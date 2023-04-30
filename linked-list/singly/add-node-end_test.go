package singly

import "testing"

func TestAddEndNode(t *testing.T) {
	var head *List
	names := []string{"Mendy", "Abeeb", "Slam", "Samuel"}
	for _, name := range names {
		head = AddEndNode(head, name)
	}
	for namesIndex := uint64(0); head != nil && namesIndex < uint64(len(names)); head, namesIndex = head.next, namesIndex + 1 {
		if head.str != names[namesIndex] {
			t.Fatalf(`At index %v: head.str = %v, want %v`, namesIndex, head.str, names[namesIndex])
		}
	}
}