package binary_search

type TimeMap[T string] struct {
	store map[string][]*Pair[T]
}

type Pair[T string] struct {
	timestamp int
	value     T
}

type ITimeMap[T string] interface {
	Set(key string, value T, timestamp int)
	Get(key string, timestamp int)
}

func NewTimeMap[T string]() *TimeMap[T] {
	store := make(map[string][]*Pair[T])
	return &TimeMap[T]{store: store}
}

func (t *TimeMap[T]) Set(key string, value T, timestamp int) {
	pair := &Pair[T]{
		timestamp: timestamp,
		value:     value,
	}

	t.store[key] = append(t.store[key], pair)
}

func (t *TimeMap[T]) Get(key string, timestamp int) T {
	// get: find key.timestamp closest to timestamp
	// if curr > timestamp -> move left
	// if curr <= timestamp
	// if end of array or next to curr > timestamp, return, else move right
	// because if curr < mid, then it's not the max, and if curr == mid, we're not sure of it's max

	if _, ok := t.store[key]; !ok {
		var zero T
		return zero
	}

	l, r := 0, len(t.store[key])-1

	for l <= r {
		mid := l + (r-l)/2
		pair := t.store[key][mid]

		if pair.timestamp > timestamp {
			r = mid - 1
		} else if pair.timestamp <= timestamp {
			if mid == len(t.store[key])-1 || t.store[key][mid+1].timestamp > timestamp {
				return pair.value
			}
			l = mid + 1
		}
	}

	var zero T
	return zero
}
