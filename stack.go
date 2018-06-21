package main

import (
	"errors"
	"fmt"
)

//Stack
type Stack []interface{}

func (stack *Stack) Push(element interface{}) {
	*stack = append(*stack, element)
}

func (stack *Stack) Peak() (interface{}, error) {
	isStackEmpty := (*stack).IsStackEmpty()
	if isStackEmpty {
		return 0, errors.New("no elements on stack")
	}
	noOfEntries := len(*stack)
	return (*stack)[noOfEntries-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	isStackEmpty := (*stack).IsStackEmpty()
	if isStackEmpty {
		return 0, errors.New("no elements on stack")
	}
	noOfEntries := len(*stack)
	popedItem := (*stack)[noOfEntries-1]
	*stack = (*stack)[0 : noOfEntries-1]
	return popedItem, nil
}

func (stack *Stack) IsStackEmpty() bool {
	noOfEntries := len(*stack)
	if noOfEntries == 0 {
		return true
	}
	return false
}

//New Stack
func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) Size() int {
	return len(*stack)
}

//PrintStack
func (stack *Stack) PrintStack() {
	fmt.Print("Stack Content(s): ")
	for _, element := range *stack {
		fmt.Printf("%v ", element)
	}
	fmt.Println()
}

func (stack *Stack) Clear() {
	*stack = (*stack)[:0]
}
