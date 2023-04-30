package singly

import (
	"testing"
)

func TestNewListWithEmptyString(t *testing.T) {
	head := NewList("", nil)
	if head != nil {
		t.Fatalf(`NewList("", nil) = %v, want, nil`, head)
	}
}

func TestNewListWithNoneEmptyString(t *testing.T) {
	name := "Mendy"
	head := NewList(name, nil)
	if head == nil || head.len != uint64(len(name)) || head.str != name {
		t.Fatalf(`head.len = %v, want %v.\nhead.str = %q, want %q`, head.len, len(name), head.str, name)
	}
}