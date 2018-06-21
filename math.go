package main

import (
	"errors"
	"fmt"
)

func GetLhsAndRhs(stack *Stack) (int, int, error) {
	if stack.Size() < 2 {
		return 0, 0, errors.New("insufficient operands")
	}
	rhs, _ := stack.Pop()
	lhs, _ := stack.Pop()
	return lhs.(int), rhs.(int), nil
}

func DoMathOperation(stack *Stack, undoStack *Stack, operation string) error {
	lhs, rhs, err := GetLhsAndRhs(stack)
	if err != nil {
		return fmt.Errorf("%v, ignoring operation %v", err, operation)
	}
	switch operation {
	case "+":
		stack.Push(lhs + rhs)
	case "-":
		stack.Push(lhs - rhs)
	case "*":
		stack.Push(lhs * rhs)
	case "/":
		stack.Push(lhs / rhs)
	default:
		return fmt.Errorf("unkown operation, ignoring %v", operation)
	}
	undoStack.Push(UndoItem{lhs: lhs, rhs: rhs})
	return nil
}
