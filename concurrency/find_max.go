package concurrency

import "sync"

// findMaxRecursive is a helper function that finds the maximum using parallel goroutines.
func findMaxRecursive(nums []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	if len(nums) == 1 {
		ch <- nums[0]
		return
	}

	mid := len(nums) / 2
	leftChan := make(chan int, 1)
	rightChan := make(chan int, 1)

	wg.Add(2)
	go findMaxRecursive(nums[:mid], leftChan, wg)
	go findMaxRecursive(nums[mid:], rightChan, wg)

	leftMax, rightMax := <-leftChan, <-rightChan
	if leftMax > rightMax {
		ch <- leftMax
	} else {
		ch <- rightMax
	}
}

// FindMax finds the maximum number in an array using goroutines.
func FindMax(nums []int) int {
	ch := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(1)

	go findMaxRecursive(nums, ch, &wg)
	wg.Wait()
	return <-ch
}
