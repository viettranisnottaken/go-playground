package hashmap

import "fmt"

func validAnagrams(s string, t string) bool {
	// convert both strings into map
	// iterate a map, if one k, v is different, return false
	// return true

	if len(s) != len(t) {
		return false
	}

	sMap, tMap := make(map[rune]int), make(map[rune]int)

	for i, char := range s {
		sMap[char]++
		// get t[i], then convert to rune
		tMap[rune(t[i])]++
	}

	for char, count := range sMap {
		if count != tMap[char] {
			return false
		}
	}

	return true
}

func TestValidAnagrams() {
	testCase1 := map[string]any{
		"input": map[string]string{
			"s": "anagram",
			"t": "nagaram",
		},
		"expect": true,
	}

	testCase2 := map[string]any{
		"input": map[string]string{
			"s": "rat",
			"t": "car",
		},
		"expect": false,
	}

	testCase3 := map[string]any{
		"input": map[string]string{
			"s": "eh",
			"t": "ehe",
		},
		"expect": false,
	}

	testCases := [...]map[string]any{
		testCase1, testCase2, testCase3,
	}

	for i := 0; i < len(testCases); i++ {
		testCase := testCases[i]
		var result bool

		input := testCase["input"].(map[string]string)
		result = validAnagrams(input["s"], input["t"])

		if result != testCase["expect"] {
			fmt.Printf("Test %d failed, returning %t instead of %t\n", i+1, result, testCase["expect"])
		} else {
			fmt.Printf("Test %d passed\n", i+1)
		}
	}
}
