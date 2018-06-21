package main

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	if stack == nil {
		t.Error("stack should not be empty")
	}
	if stack.Size() != 0 {
		t.Error("stack size should be zero")
	}
}

func TestIsStackEmpty(t *testing.T) {
	stack := NewStack()
	isEmpty := stack.IsStackEmpty()
	if !isEmpty {
		t.Error("stack is expecd to be empty")
	}
	stack.Push(1)
	isEmpty = stack.IsStackEmpty()
	if isEmpty {
		t.Error("stack is expecd to be not empty")
	}
}

func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	if stack.Size() != 2 {
		t.Error("stack size should be 2")
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	element, err := stack.Pop()
	if err == nil || element != 0 {
		t.Error("There should be a error and element should be 0")
	}
	stack.Push(1)
	stack.Push(2)
	element, err = stack.Pop()
	if err != nil || element != 2 {
		t.Error("There should be no error and element should be 2")
	}
	if stack.Size() != 1 {
		t.Error("stack size should be 1")
	}
}

func TestPeak(t *testing.T) {
	stack := NewStack()
	element, err := stack.Peak()
	if err == nil || element != 0 {
		t.Error("There should be a error and element should be 0")
	}
	stack.Push(1)
	stack.Push(2)
	element, err = stack.Peak()
	if err != nil || element != 2 {
		t.Error("There should be no error and element should be 2")
	}
	if stack.Size() != 2 {
		t.Error("stack size should be 2")
	}
}
