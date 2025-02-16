package binary_search

func SearchA2DMatrix(matrix [][]int, target int) bool {
	// search the ranges first, then search in range
	// top, bot for the ranges
	// range := matrix[row]
	// if range[0] > target, top = row - 1
	// if range[len(range) - 1] < target, bot = row + 1
	// else break

	// if cant find (top > bot), return false

	// then binary search in the selected range

	//return solution(matrix, target)
	return redo(matrix, target)
}

func redo(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	rowsCount, rowLength := len(matrix), len(matrix[0])

	top, bot := 0, rowsCount-1
	var foundRow int

	for top <= bot {
		foundRow = top + (bot-top)/2

		if matrix[foundRow][0] > target {
			// first in range
			bot = foundRow - 1
		} else if matrix[foundRow][rowLength-1] < target {
			// last in range
			top = foundRow + 1
		} else {
			break
		}
	}

	if top > bot {
		return false
	}

	l, r := 0, rowLength-1

	for l <= r {
		mid := l + (r-l)/2

		if matrix[foundRow][mid] > target {
			r = mid - 1
		} else if matrix[foundRow][mid] < target {
			l = mid + 1
		} else {
			return true
		}
	}

	return false
}

func solution(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	rows := len(matrix)

	top, bot := 0, rows-1
	var row int

	for top <= bot {
		row = bot + (top-bot)/2
		arr := matrix[row]

		if arr[0] > target {
			bot = row - 1
		} else if arr[len(arr)-1] < target {
			top = row + 1
		} else {
			break
		}
	}

	// value out of bound
	if top > bot {
		return false
	}

	arr := matrix[row]
	l, r := 0, len(arr)-1

	for l <= r {
		mid := l + (r-l)/2

		if arr[mid] > target {
			r = mid - 1
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			return true
		}
	}

	return false
}
