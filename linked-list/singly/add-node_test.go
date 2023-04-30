package singly

import "testing"

func TestAddNode(t *testing.T) {
	var head *List
	names := []string{"Mendy", "Abeeb", "Slam", "Samuel"}
	for _, name := range names {
		head = AddNode(head, name)
	}
	for namesIndex := uint64(len(names)) - 1; head != nil && namesIndex >= 0; head, namesIndex = head.next, namesIndex - 1 {
		if head.str != names[namesIndex] {
			t.Fatalf(`At index %v: head.str = %v, want %v`, namesIndex, head.str, names[namesIndex])
		}
	}
}