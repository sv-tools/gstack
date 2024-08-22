package gstack_test

import (
	"fmt"
	"github.com/sv-tools/gstack"
	"testing"
)

func TestGStack(t *testing.T) {
	s := gstack.New[int]()
	if l := s.Len(); l != 0 {
		t.Fatalf("expected empty stack, but got %d items", l)
	}
	if v := s.Pop(); v != 0 {
		t.Fatalf("expected zero item, but got %v", v)
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if l := s.Len(); l != 3 {
		t.Fatalf("expected three items, but got %d", l)
	}
	if v := s.Pop(); v != 3 {
		t.Fatalf("expected 3 as item, but got %v", v)
	}
	if l := s.Len(); l != 2 {
		t.Fatalf("expected two items, but got %d", l)
	}
	if v := s.Peek(); v != 2 {
		t.Fatalf("expected 2 as item, but got %v", v)
	}
	if v := s.String(); v != "2, 1" {
		t.Fatalf("expected '1, 2', but got %s", v)
	}
	s.Clear()
	if l := s.Len(); l != 0 {
		t.Fatalf("expected empty stack, but got %d items", l)
	}
}

func TestGStack_New(t *testing.T) {
	s := gstack.New(1, 2, 3, 4)
	if l := s.Len(); l != 4 {
		t.Fatalf("expected four items, but got %d", l)
	}
	if v := s.Pop(); v != 4 {
		t.Fatalf("expected 4 as item, but got %v", v)
	}
	if v := s.Peek(); v != 3 {
		t.Fatalf("expected 3 as item, but got %v", v)
	}
}

func TestGStack_Iter(t *testing.T) {
	s := gstack.New(1, 2, 3, 4)
	i := 4
	for v := range s.Iter() {
		if v != i {
			t.Fatalf("expected %d as item, but got %v", i, v)
		}
		i--
	}
	if s.Len() != 0 {
		t.Fatalf("expected empty stack, but got %d items", s.Len())
	}
}

func ExampleGStack() {
	s := gstack.New(1, 2, 3, 4)
	fmt.Println(s.Len())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Push(5)
	fmt.Println(s.Peek())
	// Output: 4
	// 4
	// 3
	// 2
	// 1
	// 5
}

func ExampleGStack_Iter() {
	s := gstack.New(1, 2, 3, 4)
	for v := range s.Iter() {
		fmt.Println(v)
	}
	fmt.Println("len:", s.Len())
	// Output: 4
	// 3
	// 2
	// 1
	// len: 0
}
