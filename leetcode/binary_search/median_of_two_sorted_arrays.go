package binary_search

//func MedianOfTwoSortedArrays(nums1 []int, nums2 []int) float32 {
//	// [1, 2, 3, 4, 5] [1, 2, 3, 4, 5, 6]
//	// To find median, we need to partition the array into left and right partition
//	// We have 2 sorted arrays, so left partition should generally be on the left side of both
//	// LeftPartition[len(LeftPartition) - 1] <= RightPartition[0]
//	// We work on smaller array of the 2
//	// If LeftPartition[len(LeftPartition) - 1] > RightPartition[0], find el in an smaller array that
//	// satisfies LeftPartition[len(LeftPartition) - 1] <= RightPartition[0]
//	// LeftPartition = Left ArrayA + Left ArrayB, so if we have left ArrayA, we can find left ArrayB
//	// Find median: if len(arrayA) + len(arrayB) is odd, then its min RightPartition,
//	// if even, then its (max of LeftPartition + min of RightPartition) / 2 ,
//
//	totalLen := len(nums1) + len(nums2)
//	half := (totalLen + 1) / 2
//
//	var selectedArr *[]int
//	var remainingArr *[]int
//	if len(nums1) <= len(nums2) {
//		selectedArr = &nums1
//		remainingArr = &nums2
//	} else {
//		selectedArr = &nums2
//		remainingArr = &nums1
//	}
//
//	l, r := 0, len(*selectedArr)-1
//
//	for l <= r {
//		mid1 := l + (r-l)/2
//		mid2 := half - mid1
//
//		//leftPart1 :=
//	}
//}
