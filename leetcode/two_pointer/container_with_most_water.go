package two_pointer

func ContainerWithMostWater(heights []int) int {
	// start, end pointers
	// res := 0
	// loop heights
	// area := min(nums[start], nums[end]) * (end - start)
	// res = max(area, res)
	// whatever pointer smaller, shifts

	res := 0

	start, end := 0, len(heights)-1

	for start < end {
		area := min(heights[start], heights[end]) * (end - start)
		res = max(area, res)

		if heights[start] <= heights[end] {
			start++
		} else {
			end--
		}
	}

	return res
}
