package concurrency

import (
	"sync"
)

func SumOfSquares(nums []int) int {
	// res, wg and mutex
	// wg to determine when all goroutines are done
	// mutex for the res

	res := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, num := range nums {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			res += n * n

		}(num)
	}

	wg.Wait()

	return res
}
