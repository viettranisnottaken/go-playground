package stack

import (
	"fmt"
	ds "leetcode/data_structure"
)

func validParentheses(str string) bool {
	// idea: push open bracket to stack, if encounter a close bracket, pop a corresponding open one
	// create a map that map open to close brackets: close as key, open as value
	// create a stack
	// iterate through string
	// if open, push (check for !exist in map)
	// if close, check if top of stack is valid, if so, pop, if not return false
	// return stack.length === 0;

	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := ds.NewStack[rune]()

	for _, char := range str {
		if bracket, isExist := bracketMap[char]; isExist {
			if top, isStackNotEmpty := stack.Peek(); !isStackNotEmpty || bracket != top {
				return false
			}

			stack.Pop()
		} else {
			stack.Push(char)
		}
	}

	return stack.IsEmpty()
}

func TestValidParentheses() {
	testCase1 := map[string]any{
		"input":  "[]",
		"expect": true,
	}

	testCase2 := map[string]any{
		"input":  "([{}])",
		"expect": true,
	}

	testCase3 := map[string]any{
		"input":  "[(])",
		"expect": false,
	}

	testCase4 := map[string]any{
		"input":  "()[]{}",
		"expect": true,
	}

	testCases := [...]map[string]any{
		testCase1, testCase2, testCase3, testCase4,
	}

	for i := 0; i < len(testCases); i++ {
		testCase := testCases[i]
		var result bool

		input := testCase["input"].(string)
		result = validParentheses(input)

		if result != testCase["expect"] {
			fmt.Printf("Test %d failed, returning %t instead of %t\n", i+1, result, testCase["expect"])
		} else {
			fmt.Printf("Test %d passed\n", i+1)
		}
	}
}
