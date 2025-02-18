package binary_search

func SearchInRotatedSortedArray(nums []int, target int) int {
	return searchInRotatedSortedArraySolution(nums, target)
}

func searchInRotatedSortedArrayRedo(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		}

		// check from mid, to what end is sorted
		// if nums[mid] <= nums[r], then its right side is sorted
		// else its left side is sorted
		// if right side is sorted, check if target is between mid and r, if so, move right, else move left
		// if left side is sorted, heck if target is between mid and l, if so, move left, else move right
		// we check target <= nums[r] and target >= nums[l] for small arrays, or when we finally cut
		// down to small arrays (test case 4 - 6)

		if nums[mid] < nums[r] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[mid] > target && target >= nums[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}

func searchInRotatedSortedArraySolution(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		}

		// mid < right -> sorted
		//if nums[r] >= nums[mid] { -> same thing
		if nums[r] > nums[mid] {
			// right side of pivot is sorted
			// check what half target belongs
			if nums[mid] < target && target <= nums[r] {
				// if on right side: mid < target < right
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			// left side of pivot is sorted
			if nums[l] <= target && target < nums[mid] {
				// if on left side: left < target < mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}
