package gstack

import (
	"bytes"
	"fmt"
	"iter"
)

type node[T any] struct {
	Value T
	Next  *node[T]
}

// GStack is generic Stack implementation using linked list
type GStack[T any] struct {
	top *node[T]
	len int
}

// New creates an instance of GStack and pushes all given items to the created stack in reversed order
func New[T any](items ...T) GStack[T] {
	s := GStack[T]{}
	for _, item := range items {
		s.Push(item)
	}
	return s
}

// Len returns the length of the stack
func (s *GStack[T]) Len() int {
	return s.len
}

// Peek returns the top item without removing it or a zero object of given type
func (s *GStack[T]) Peek() T {
	if s.top == nil {
		var zero T
		return zero
	}
	return s.top.Value
}

// Push puts the given item into the stack
func (s *GStack[T]) Push(item T) {
	s.top = &node[T]{
		Value: item,
		Next:  s.top,
	}
	s.len++
}

// Pop returns the top item and removes it from the stack or returns a zero object of given type
func (s *GStack[T]) Pop() T {
	if s.top == nil {
		var zero T
		return zero
	}

	item := s.top.Value
	s.top = s.top.Next
	s.len--
	return item
}

// String returns the list of all items in text format
func (s *GStack[T]) String() string {
	if s.len == 0 {
		return ""
	}
	var buf bytes.Buffer
	top := s.top
	for top != nil {
		buf.WriteString(fmt.Sprintf("%v", top.Value))
		if top.Next != nil {
			buf.WriteString(", ")
		}
		top = top.Next
	}
	return buf.String()
}

// Clear clears the stack
func (s *GStack[T]) Clear() {
	s.top = nil
	s.len = 0
}

// Iter returns a consuming iterator for the stack
func (s *GStack[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for s.top != nil {
			if !yield(s.Pop()) {
				break
			}
		}
	}
}

// IsEmpty returns true if the stack is empty
func (s *GStack[T]) IsEmpty() bool {
	return s.top == nil
}
