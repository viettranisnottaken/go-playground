package binary_search

func FindMinimumInRotatedSortedArray(nums []int) int {
	// l, r, mid
	// array is rotated, so we can split it into 2 parts: before min and after min
	// Ex: [4, 5, 6, 0, 1, 2, 3]: before 0 and after 0
	// while l < r (because when l = r, we found the min, read below)
	// we find min, so we need to be moving toward min
	// mid is in rotated part, then nums[mid] > nums[r] and min is on the right -> move right
	// if array is not rotated, then nums[mid] <= nums[r] and min is on the left -> move left
	// to cover the case where nums[mid] = min, when nums[mid] <= nums[r] we move left and r = mid
	// return nums[l]

	return findMinimumInRotatedSortedArrayRedo(nums)
}

func findMinimumInRotatedSortedArrayRedo(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}

func findMinimumInRotatedSortedArraySolution(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}
