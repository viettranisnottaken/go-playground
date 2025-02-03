package ds

type Deque[T any] struct {
	Items *DoublyLinkedList[T]
}

type IDeque[T any] interface {
	InsertFirst(value T)
	InsertLast(value T)
	DeleteFirst() bool
	DeleteLast() bool
	GetFirst() (T, bool)
	GetLast() (T, bool)
	IsEmpty() bool
	Size() int
}

func NewDeque[T any](values ...T) *Deque[T] {
	return &Deque[T]{
		Items: NewDoublyLinkedList[T](values...),
	}
}

func (d *Deque[T]) InsertFirst(value T) {
	d.InsertFirst(value)
}

func (d *Deque[T]) InsertLast(value T) {
	d.InsertLast(value)
}

func (d *Deque[T]) DeleteFirst() bool {
	return d.DeleteFirst()
}

func (d *Deque[T]) DeleteLast() bool {
	return d.DeleteLast()
}

func (d *Deque[T]) GetFirst() (T, bool) {
	return d.GetFirst()
}

func (d *Deque[T]) GetLast() (T, bool) {
	return d.GetLast()
}

func (d *Deque[T]) IsEmpty() bool {
	return d.IsEmpty()
}

func (d *Deque[T]) Size() int {
	return d.Items.Length
}
