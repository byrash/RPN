package main

import (
	"testing"
)

func TestGetLhsAndRhs(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	lhs, rhs, err := GetLhsAndRhs(stack)
	if lhs != 1 && rhs != 2 && err != nil {
		t.Error("expected lsh to be 1 and rhs to be 2 and no error")
	}
	stack.Pop()
	lhs, rhs, err = GetLhsAndRhs(stack)
	if err == nil {
		t.Error("expected to have error")
	}
}

func TestDoMathOperation(t *testing.T) {
	stack := NewStack()
	undoStack := NewStack()
	// Add
	stack.Push(1)
	stack.Push(2)
	err := DoMathOperation(stack, undoStack, "+")
	if err != nil {
		t.Error("there should be no error")
	}
	topItem, _ := stack.Peak()
	if topItem != 3 {
		t.Error("Expected 3")
	}
	stack.Clear()
	undoStack.Clear()

	// Sub
	stack.Push(2)
	stack.Push(1)
	err = DoMathOperation(stack, undoStack, "-")
	if err != nil {
		t.Error("there should be no error")
	}
	topItem, _ = stack.Peak()
	if topItem != 1 {
		t.Error("Expected 1")
	}
	stack.Clear()
	undoStack.Clear()

	// Sub
	stack.Push(2)
	stack.Push(1)
	err = DoMathOperation(stack, undoStack, "*")
	if err != nil {
		t.Error("there should be no error")
	}
	topItem, _ = stack.Peak()
	if topItem != 2 {
		t.Error("Expected 2")
	}
	stack.Clear()
	undoStack.Clear()

	// Div
	stack.Push(2)
	stack.Push(2)
	err = DoMathOperation(stack, undoStack, "/")
	if err != nil {
		t.Error("there should be no error")
	}
	topItem, _ = stack.Peak()
	if topItem != 1 {
		t.Error("Expected 1")
	}
	stack.Clear()
	undoStack.Clear()

	//Unkown operation
	err = DoMathOperation(stack, undoStack, "%")
	if err == nil {
		t.Error("there should be a error")
	}
}
