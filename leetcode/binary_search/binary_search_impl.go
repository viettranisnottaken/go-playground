package binary_search

func BinarySearchImpl(nums []int, target int) int {
	// left and right ptr
	// while l <= r
	// mid : (l + r) / 2
	// if mid < target, left = mid + 1
	// if mid > target, right = mid - 1
	// if mid == target, return mid
	// return -1
	l, r := 0, len(nums)-1

	for l <= r {
		// unsafe for large int, may be out of bound
		//mid := (l + r) / 2

		// safe
		mid := l + (r-l)/2
		curr := nums[mid]
		if curr == target {
			return mid
		} else if curr < target {
			l = mid + 1
		} else if curr > target {
			r = mid - 1
		}
	}

	return -1
}
