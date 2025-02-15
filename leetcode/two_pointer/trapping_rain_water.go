package two_pointer

func TrappingRainWater(height []int) int {
	// water at i = min(prefix[i], suffix[i]) - height[i]

	//return prefixSuffix(height)
	return twoPointer(height)
}

func prefixSuffix(height []int) int {
	// create a prefix sum for max left value
	// create a suffix sum for max right value
	// iterate arr, use formula, res +=

	res := 0

	// prefix
	prefix := make([]int, len(height))
	currMaxL := 0
	for i, h := range height {
		if h > currMaxL {
			currMaxL = h
		}

		prefix[i] = currMaxL
	}

	// suffix
	suffix := make([]int, len(height))
	currMaxR := 0
	for i := len(height) - 1; i >= 0; i-- {
		h := height[i]

		if h > currMaxR {
			currMaxR = h
		}

		suffix[i] = currMaxR
	}

	for i, h := range height {
		waterAtI := min(prefix[i], suffix[i]) - h

		res += max(waterAtI, 0)
	}

	return res
}

func twoPointer(height []int) int {
	// left ptr, right ptr
	// currMaxL, currMaxR
	// while left < right
	// if left < right, left++, res += currMaxL, update currMaxL

	res := 0
	left, right := 0, len(height)-1
	currMaxL, currMaxR := height[left], height[right]

	for left < right {
		if currMaxL < currMaxR {
			left++
			res += max(currMaxL-height[left], 0)
			currMaxL = max(currMaxL, height[left])
		} else {
			right--
			res += max(currMaxR-height[right], 0)
			currMaxR = max(currMaxR, height[right])
		}
	}

	return res
}
