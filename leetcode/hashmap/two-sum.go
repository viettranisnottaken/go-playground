package hashmap

import (
	"fmt"
	"reflect"
)

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		curr := nums[i]

		if num, valid := hashMap[curr]; valid {
			return []int{num, i}
		}

		hashMap[target-curr] = i
	}

	return nil
}

// test cases

func TestTwoSum() {
	test1 := func() {
		expectation := []int{0, 1}
		result := twoSum([]int{2, 7, 11, 15}, 9)

		if !reflect.DeepEqual(result, expectation) {
			fmt.Println("test1 failed")
		} else {
			fmt.Println("test1 passed")
		}
	}

	test2 := func() {
		expectation := []int{1, 2}
		result := twoSum([]int{3, 2, 4}, 6)

		if !reflect.DeepEqual(result, expectation) {
			fmt.Println("test2 failed")
		} else {
			fmt.Println("test2 passed")
		}
	}

	test3 := func() {
		expectation := []int{0, 1}
		result := twoSum([]int{3, 3}, 6)

		if !reflect.DeepEqual(result, expectation) {
			fmt.Println("test3 failed")
		} else {
			fmt.Println("test3 passed")
		}
	}

	test1()
	test2()
	test3()
}
