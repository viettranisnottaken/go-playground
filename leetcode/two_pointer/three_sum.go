package two_pointer

import "sort"

func ThreeSum(nums []int) [][3]int {
	// sort nums
	// res := [][3]int
	// loop sorted
	// for each i, create 2 pointers
	// if nums[i] > 0, break
	// if curr num === last iteration num, continue
	// if i + start + end < 0, start++
	// if i + start + end > 0, end--
	// if i + start + end == 0, push, start++, end--
	// skip duplications
	// return res

	res := make([][3]int, 0)
	sort.Ints(nums)

	for i, num := range nums {
		if num > 0 {
			break
		}

		if i > 0 && num == nums[i-1] {
			continue
		}

		// two sum
		start, end := i+1, len(nums)-1

		for start < end {
			sum := num + nums[start] + nums[end]

			if sum > 0 {
				end--
			} else if sum < 0 {
				start++
			} else {
				res = append(res, [3]int{num, nums[start], nums[end]})

				start++
				end--

				// skip duplication
				for start < end && nums[start] == nums[start-1] {
					start++
				}
			}
		}
	}

	return res
}
