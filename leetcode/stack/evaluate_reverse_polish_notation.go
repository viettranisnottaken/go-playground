package stack

import (
	ds "leetcode/data_structure"
	"strconv"
)

func EvaluateReversePolishNotation(tokens []string) int {
	// create a stack
	// iterate through tokens, push in one by one
	// if encounter an operator, execute what's in the stack, then push result in

	stack := ds.NewStack[int]()

	for _, token := range tokens {
		switch token {
		case "+":
			num1, _ := stack.Pop()
			num2, _ := stack.Pop()

			stack.Push(num2 + num1)
		case "-":
			num1, _ := stack.Pop()
			num2, _ := stack.Pop()

			stack.Push(num2 - num1)
		case "*":
			num1, _ := stack.Pop()
			num2, _ := stack.Pop()

			stack.Push(num2 * num1)
		case "/":
			num1, _ := stack.Pop()
			num2, _ := stack.Pop()

			stack.Push(num2 / num1)
		default:
			num, _ := strconv.Atoi(token)
			stack.Push(num)
		}
	}

	result, _ := stack.Peek()

	return result
}
