package ds

type IDynamicArray[T any] interface {
	InsertFirst(value T)
	InsertLast(value T)
	InsertAt(value T, i int) bool
	DeleteFirst() bool
	DeleteLast() bool
	DeleteAt(i int) bool
	GetItem(i int) (T, bool)
	IsEmpty() bool

	grow()
	shrink()
}

type DynamicArray[T any] struct {
	Items    []T
	CurrSize int
	Capacity int
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
		Items:    items,
		CurrSize: len(values),
		Capacity: defaultCap,
	}
}

func (d *DynamicArray[T]) InsertFirst(value T) {
	// check cap, then grow if needed
	if d.CurrSize == d.Capacity {
		d.grow()
	}

	// move all el to +1 index
	for i := d.CurrSize - 1; i >= 0; i-- {
		newIndex := i + 1

		d.Items[newIndex] = d.Items[i]
	}

	// insert to start
	d.Items[0] = value

	// update CurrSize
	d.CurrSize++
}

func (d *DynamicArray[T]) InsertLast(value T) {
	// check cap, then grow if needed
	if d.CurrSize == d.Capacity {
		d.grow()
	}

	// insert
	d.Items[d.CurrSize] = value

	// update CurrSize
	d.CurrSize++
}

func (d *DynamicArray[T]) InsertAt(value T, i int) bool {
	// if insertion point is out of bound, return error
	if d.CurrSize-1 < i {
		return false
	}

	// check cap, then grow if needed
	if d.CurrSize == d.Capacity {
		d.grow()
	}

	// move all el from insertion point to +1 index
	for currI := d.CurrSize - 1; currI >= i; currI-- {
		newIndex := currI + 1

		d.Items[newIndex] = d.Items[currI]
	}

	// insert
	d.Items[i] = value

	// update CurrSize
	d.CurrSize++

	return true
}

func (d *DynamicArray[T]) DeleteFirst() bool {
	// check if array is empty
	if d.CurrSize == 0 {
		return false
	}

	// check cap, then shrink if needed
	if d.CurrSize <= d.Capacity/4 {
		d.shrink()
	}

	// move all el to -1 index
	for i := 0; i < d.CurrSize; i++ {
		currIndex := i + 1

		d.Items[i] = d.Items[currIndex]
	}

	// delete last Value
	d.Items = d.Items[:d.CurrSize-1]

	// update CurrSize
	d.CurrSize--

	return true
}

func (d *DynamicArray[T]) DeleteLast() bool {
	// check if array is empty
	if d.CurrSize == 0 {
		return false
	}

	// check cap, then shrink if needed
	if d.CurrSize <= d.Capacity/4 {
		d.shrink()
	}

	// delete last Value
	d.Items = d.Items[:d.CurrSize-1]

	// update CurrSize
	d.CurrSize--

	return true
}

func (d *DynamicArray[T]) DeleteAt(i int) bool {
	// check if array is empty
	if d.CurrSize == 0 {
		return false
	}

	// check cap, then shrink if needed
	if d.CurrSize <= d.Capacity/4 {
		d.shrink()
	}

	// move all el from insertion point to -1 index
	for index := i; i < d.CurrSize; i++ {
		currIndex := index + 1

		d.Items[index] = d.Items[currIndex]
	}

	// delete last Value
	d.Items = d.Items[:d.CurrSize-1]

	// update CurrSize
	d.CurrSize--

	return true
}

func (d *DynamicArray[T]) GetItem(i int) (T, bool) {
	// check if i is out of bound
	if d.CurrSize-1 < i {
		var zero T
		return zero, false
	}

	return d.Items[i], true
}

func (d *DynamicArray[T]) IsEmpty() bool {
	return d.CurrSize == 0
}

func (d *DynamicArray[T]) grow() {
	// make new cap = old cap * 2
	newCap := d.Capacity * 2
	// create a new array
	newArr := make([]T, newCap)

	// move elements over to new arr
	for i, v := range d.Items {
		newArr[i] = v
	}

	// assign new array to dynamic array
	d.Items = newArr
	// update Capacity
	d.Capacity = newCap
}

func (d *DynamicArray[T]) shrink() {
	// make new cap = old cap / 2
	newCap := d.Capacity / 2
	// create a new array
	newArr := make([]T, newCap)

	// move elements over to new arr
	for i, v := range d.Items {
		newArr[i] = v
	}

	// assign new array to dynamic array
	d.Items = newArr
	// update Capacity
	d.Capacity = newCap
}
