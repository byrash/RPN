package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type UndoItem struct {
	rhs int
	lhs int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	stack := NewStack()
	undoStack := NewStack()
	for {
		fmt.Println("Enter the expression pattern")
		line, _ := reader.ReadString('\n')
		line = strings.Replace(line, "\n", "", -1)
		if strings.EqualFold(line, "exit") {
			os.Exit(0)
		}
		parseInputLineAndProcess(line, stack, undoStack)
		fmt.Println()
		stack.PrintStack()
	}
}

func parseInputLineAndProcess(expression string, stack *Stack, undoStack *Stack) {
	expressionElements := strings.Split(expression, " ")
	for _, element := range expressionElements {
		if number, err := strconv.Atoi(element); err == nil {
			stack.Push(number)
		} else {
			switch element {
			case "undo":
				err := UndoOpertion(stack, undoStack)
				PrintErrorIfExist(err)
			case "clear":
				stack.Clear()
				undoStack.Clear()
			default:
				err := DoMathOperation(stack, undoStack, element)
				PrintErrorIfExist(err)
			}
		}
	}
}

func PrintErrorIfExist(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func UndoOpertion(stack, undoStack *Stack) error {
	if undoStack.IsStackEmpty() {
		return fmt.Errorf("No more actions left to undo")
	}
	if stack.IsStackEmpty() {
		return fmt.Errorf("No actions performed on stack")
	}
	_, _ = stack.Pop() // Just pop last value
	tmp, _ := undoStack.Pop()
	undoItem := tmp.(UndoItem)
	stack.Push(undoItem.lhs)
	stack.Push(undoItem.rhs)
	return nil
}
