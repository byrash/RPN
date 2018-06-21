package main

import (
	"testing"
)

func TestParseInputLineAndProcess(t *testing.T) {
	stack := NewStack()
	undoStack := NewStack()
	parseInputLineAndProcess("1 2 3 4 5", stack, undoStack)
	if stack.Size() != 5 {
		t.Error("expected 5 times on stack")
	}
	if !undoStack.IsStackEmpty() {
		t.Error("undo stack should be empty")
	}
	parseInputLineAndProcess("* * * *", stack, undoStack)
	if stack.Size() != 1 {
		t.Error("expected 1 time on stack")
	}
	if undoStack.Size() != 4 {
		t.Error("undo stack should have 4 times")
	}
	parseInputLineAndProcess("undo undo undo undo", stack, undoStack)
	if stack.Size() != 5 {
		t.Error("expected 5 times on stack")
	}
	if !undoStack.IsStackEmpty() {
		t.Error("undo stack should be empty")
	}
	parseInputLineAndProcess("undo", stack, undoStack)
	if stack.Size() != 5 {
		t.Error("expected 5 times on stack")
	}
	if !undoStack.IsStackEmpty() {
		t.Error("undo stack should be empty")
	}
}

func TestUndoOpertion(t *testing.T) {
	stack := NewStack()
	undoStack := NewStack()
	parseInputLineAndProcess("1 2 3 4 5", stack, undoStack)
	err := UndoOpertion(stack, undoStack)
	if err == nil {
		t.Error("expected error")
	}
	parseInputLineAndProcess("*", stack, undoStack)
	err = UndoOpertion(stack, undoStack)
	if err != nil {
		t.Error("expected no error")
	}
}
