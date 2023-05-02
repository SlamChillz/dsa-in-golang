package singly

import "testing"


func TestWithoutLoop(t *testing.T) {
	var head *List
	seed := []string{"Mendy", "Abeeb", "Slam", "Samuel"}
	for _, name := range seed {
		head = AddNode(head, name)
	}
	loopNode := FindLoopNode(head)
	if loopNode != nil {
		t.Fatalf(`Expected loopNode == <nil>, got %+v`, loopNode)
	}
}

func TestWithLoopNode(t *testing.T) {
	var head *List
	tail := AddEndNode(head, "Jenkins")
	head = tail
	loopIndex := 2
	seed := []string{"Mendy", "Abeeb", "Slam", "Samuel"}
	for i, name := range seed {
		head = AddNode(head, name)
		if i == loopIndex {
			tail.next = head
		}
	}
	loopNode := FindLoopNode(head)
	if loopNode != tail.next {
		t.Fatalf(`Expected loopNode == %+v, got %+v`, tail.next, loopNode)
	}
	if loopNode.str != seed[loopIndex] {
		t.Fatalf(`Expected name at loopNode == %v, got %v`, seed[loopIndex], loopNode.str)
	}
}