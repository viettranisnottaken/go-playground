package hashmap

import "fmt"

func longestConsecutiveSequence(nums []int) int {
	// convert nums to hashset
	// maxLength := 0
	// iterate hashset
	// if num is start of sequence (!set[num - 1]), start counting
	// count: length := 1; if !!set[num + length], length++, else break; set maxLength = max(length, maxLength)
	// return maxLength

	set := make(map[int]bool)

	for _, num := range nums {
		set[num] = true
	}

	maxLength := 0

	for num := range set {
		if _, isExist := set[num-1]; !isExist {
			length := 1

			for {
				if _, isPresent := set[num+length]; isPresent {
					length++
				} else {
					break
				}
			}

			if length > maxLength {
				maxLength = length
			}
		}
	}

	return maxLength
}

func TestLongestConsecutiveSequence() {
	testCase1 := map[string]any{
		"input":  []int{2, 20, 4, 10, 3, 4, 5},
		"expect": 4,
	}

	testCase2 := map[string]any{
		"input":  []int{0, 3, 2, 5, 4, 6, 1, 1},
		"expect": 7,
	}

	testCase3 := map[string]any{
		"input":  []int{0, 0, 0, 0, 0},
		"expect": 1,
	}

	testCases := [...]map[string]any{
		testCase1, testCase2, testCase3,
	}

	for i := 0; i < len(testCases); i++ {
		testCase := testCases[i]
		var result int

		inputArray := testCase["input"].([]int)            // Use the correct type here
		result = longestConsecutiveSequence(inputArray[:]) // Convert array to slice with `[:]`

		if result != testCase["expect"] {
			fmt.Printf("Test %d failed, returning %d instead of %d\n", i+1, result, testCase["expect"])
		} else {
			fmt.Printf("Test %d passed\n", i+1)
		}
	}
}
