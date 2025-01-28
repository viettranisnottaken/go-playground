package hashmap

import "fmt"

func containsDuplicate(nums []int) bool {
	// create a map
	// convert nums to map
	// if finds a num, return true
	// returns false

	numsMap := make(map[int]bool)

	for _, num := range nums {
		if _, isPresent := numsMap[num]; isPresent {
			return true
		}

		numsMap[num] = true
	}

	return false
}

func TestContainsDuplicate() {
	testCase1 := map[string]any{
		"input":  []int{1, 2, 3, 1},
		"expect": true,
	}

	testCase2 := map[string]any{
		"input":  []int{1, 2, 3, 4},
		"expect": false,
	}

	testCase3 := map[string]any{
		"input":  []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		"expect": true,
	}

	testCases := [...]map[string]any{
		testCase1, testCase2, testCase3,
	}

	for i := 0; i < len(testCases); i++ {
		testCase := testCases[i]
		var result bool

		inputArray := testCase["input"].([]int)   // Use the correct type here
		result = containsDuplicate(inputArray[:]) // Convert array to slice with `[:]`

		if result != testCase["expect"] {
			fmt.Printf("Test %d failed, returning %t instead of %t\n", i+1, result, testCase["expect"])
		} else {
			fmt.Printf("Test %d passed\n", i+1)
		}
	}
}
