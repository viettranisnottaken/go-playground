package two_pointer

import (
	"fmt"
	"unicode"
)

func isAlphaNum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func validPalindrome(s string) bool {
	// start, end pointers
	// if start sees !isAlphaNum, advance
	// if end sees !isAlphaNum, advance
	// if s[start] != s[end], return false
	// end loop when start >= end
	// return true

	start := 0
	end := len(s) - 1

	for start < end {
		for start < end && !isAlphaNum(rune(s[start])) {
			start++
		}

		for end > start && !isAlphaNum(rune(s[end])) {
			end--
		}

		if unicode.ToLower(rune(s[start])) != unicode.ToLower(rune(s[end])) {
			return false
		}

		start++
		end--
	}

	return true
}

func TestValidPalindrome() {
	testCase1 := map[string]any{
		"input":  "A man, a plan, a canal: Panama",
		"expect": true,
	}

	testCase2 := map[string]any{
		"input":  "race a car",
		"expect": false,
	}

	testCase3 := map[string]any{
		"input":  " ",
		"expect": true,
	}

	testCases := [...]map[string]any{
		testCase1, testCase2, testCase3,
	}

	for i := 0; i < len(testCases); i++ {
		testCase := testCases[i]
		var result bool

		input := testCase["input"].(string)
		result = validPalindrome(input)

		if result != testCase["expect"] {
			fmt.Printf("Test %d failed, returning %t instead of %t\n", i+1, result, testCase["expect"])
		} else {
			fmt.Printf("Test %d passed\n", i+1)
		}
	}
}
