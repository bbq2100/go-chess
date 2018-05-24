package pkg

import (
	"testing"
)

func TestAddNewElement(t *testing.T) {
	s := newStack()
	s.add("hallo")
	s.add("welt")

	if s.pop() != "welt" {
		t.Fail()
	}

	if s.pop() != "hallo" {
		t.Fail()
	}

	if s.pop() != nil {
		t.Fail()
	}
}

func TestEmptyStack(t *testing.T) {
	s := newStack()
	if s.pop() != nil {
		t.Fail()
	}
}
