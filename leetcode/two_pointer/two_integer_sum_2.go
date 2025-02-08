package two_pointer

// https://neetcode.io/problems/two-integer-sum-ii

func TwoIntegerSum2(numbers []int, target int) [2]int {
	// 2 pointers start, end
	// while loop, stops when start >= end
	// if start + end == target, break
	// if start + end > target, end--
	// if start + end < target, start++
	// return [2]int{start + 1, end + 1}

	start := 0
	end := len(numbers) - 1

	for start < end {
		if numbers[start]+numbers[end] == target {
			break
		}

		if numbers[start]+numbers[end] > target {
			end--
		}

		if numbers[start]+numbers[end] < target {
			start++
		}
	}

	return [2]int{start + 1, end + 1}
}
