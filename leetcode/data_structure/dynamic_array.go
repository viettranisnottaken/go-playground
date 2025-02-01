package ds

type IDynamicArray[T any] interface {
	insertFirst(value T)
	insertLast(value T)
	insertAt(value T, i int) bool
	deleteFirst() bool
	deleteLast() bool
	deleteAt(i int) bool
	getItem(i int) (T, bool)

	grow()
	shrink()
}

type DynamicArray[T any] struct {
	items    []T
	currSize int
	capacity int
}

func NewDynamicArray[T any](values ...T) *DynamicArray[T] {
	// go does not allow run time array size, so I have to use slice
	var defaultCap int

	if len(values) == 0 {
		defaultCap = 2
	} else {
		defaultCap = len(values) * 2
	}

	items := make([]T, defaultCap)

	for i, v := range values {
		items[i] = v
	}

	return &DynamicArray[T]{
		items:    items,
		currSize: len(values),
		capacity: defaultCap,
	}
}

func (d *DynamicArray[T]) insertFirst(value T) {
	// check cap, then grow if needed
	if d.currSize == d.capacity {
		d.grow()
	}

	// move all el to +1 index
	for i := d.currSize - 1; i >= 0; i-- {
		newIndex := i + 1

		d.items[newIndex] = d.items[i]
	}

	// insert to start
	d.items[0] = value

	// update currSize
	d.currSize++
}

func (d *DynamicArray[T]) insertLast(value T) {
	// check cap, then grow if needed
	if d.currSize == d.capacity {
		d.grow()
	}

	// insert
	d.items[d.currSize] = value

	// update currSize
	d.currSize++
}

func (d *DynamicArray[T]) insertAt(value T, i int) bool {
	// if insertion point is out of bound, return error
	if d.currSize-1 < i {
		return false
	}

	// check cap, then grow if needed
	if d.currSize == d.capacity {
		d.grow()
	}

	// move all el from insertion point to +1 index
	for currI := d.currSize - 1; currI >= i; currI-- {
		newIndex := currI + 1

		d.items[newIndex] = d.items[currI]
	}

	// insert
	d.items[i] = value

	// update currSize
	d.currSize++

	return true
}

func (d *DynamicArray[T]) deleteFirst() bool {
	// check if array is empty
	if d.currSize == 0 {
		return false
	}

	// check cap, then shrink if needed
	if d.currSize <= d.capacity/4 {
		d.shrink()
	}

	// move all el to -1 index
	for i := 0; i < d.currSize; i++ {
		currIndex := i + 1

		d.items[i] = d.items[currIndex]
	}

	// delete last value
	d.items = d.items[:d.currSize-1]

	// update currSize
	d.currSize--

	return true
}

func (d *DynamicArray[T]) deleteLast() bool {
	// check if array is empty
	if d.currSize == 0 {
		return false
	}

	// check cap, then shrink if needed
	if d.currSize <= d.capacity/4 {
		d.shrink()
	}

	// delete last value
	d.items = d.items[:d.currSize-1]

	// update currSize
	d.currSize--

	return true
}

func (d *DynamicArray[T]) deleteAt(i int) bool {
	// check if array is empty
	if d.currSize == 0 {
		return false
	}

	// check cap, then shrink if needed
	if d.currSize <= d.capacity/4 {
		d.shrink()
	}

	// move all el from insertion point to -1 index
	for index := i; i < d.currSize; i++ {
		currIndex := index + 1

		d.items[index] = d.items[currIndex]
	}

	// delete last value
	d.items = d.items[:d.currSize-1]

	// update currSize
	d.currSize--

	return true
}

func (d *DynamicArray[T]) getItem(i int) (T, bool) {
	// check if i is out of bound
	if d.currSize-1 < i {
		var zero T
		return zero, false
	}

	return d.items[i], true
}

func (d *DynamicArray[T]) grow() {
	// make new cap = old cap * 2
	newCap := d.capacity * 2
	// create a new array
	newArr := make([]T, newCap)

	// move elements over to new arr
	for i, v := range d.items {
		newArr[i] = v
	}

	// assign new array to dynamic array
	d.items = newArr
	// update capacity
	d.capacity = newCap
}

func (d *DynamicArray[T]) shrink() {
	// make new cap = old cap / 2
	newCap := d.capacity / 2
	// create a new array
	newArr := make([]T, newCap)

	// move elements over to new arr
	for i, v := range d.items {
		newArr[i] = v
	}

	// assign new array to dynamic array
	d.items = newArr
	// update capacity
	d.capacity = newCap
}
