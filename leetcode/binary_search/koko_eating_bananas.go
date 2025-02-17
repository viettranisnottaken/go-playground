package binary_search

import "math"

func KokoEatingBananas(piles []int, h int) int {
	// we search speed from 1 to max(piles)
	// loop through arr to calculate and check if total time <= h
	return kokoEatingBananasRedo(piles, h)
}

func kokoEatingBananasRedo(piles []int, h int) int {
	// search speed with left ptr = min speed (1 is smallest possible), and right ptr = max(piles)
	// mid; total time at speed mid is valid when <= h
	// if valid (fast enough), res = mid, then move to left side to find smaller speed (we need the smallest possible mid)
	// if not valid (too slow), move to the right to find faster speed

	minSpeed, maxSpeed := 1, getMax(piles)
	res := maxSpeed

	for minSpeed <= maxSpeed {
		speed := minSpeed + (maxSpeed-minSpeed)/2
		totalTime := 0

		// calculate how much time to finish all piles at speed
		for _, pile := range piles {
			totalTime += int(math.Ceil(float64(pile) / float64(speed)))
		}

		if totalTime <= h {
			res = speed
			maxSpeed = speed - 1
		} else {
			minSpeed = speed + 1
		}
	}

	return res
}

func kokoEatingBananasSolution(piles []int, h int) int {
	minSpeed, maxSpeed := 1, getMax(piles)
	res := maxSpeed

	for minSpeed <= maxSpeed {
		speed := minSpeed + (maxSpeed-minSpeed)/2
		totalTime := 0

		// loop through arr to calculate
		for _, pile := range piles {
			totalTime += int(math.Ceil(float64(pile) / float64(speed)))
		}

		if totalTime <= h {
			// if valid, we move to the left to find even smaller number
			res = speed
			maxSpeed = speed - 1
		} else {
			// if invalid, we move to the right to find valid
			minSpeed = speed + 1
		}
	}

	return res
}

func getMax(nums []int) int {
	res := 0
	for _, num := range nums {
		if num > res {
			res = num
		}
	}

	return res
}
