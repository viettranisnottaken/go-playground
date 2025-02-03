package ds

// Node TODO: add multiple types support
type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type DoublyLinkedList[T any] struct {
	Length int
	Head   *Node[T]
	Tail   *Node[T]
}

type IDoublyLinkedList[T any] interface {
	InsertFirst(value T)
	InsertLast(value T)
	InsertAt(value T, i int) bool
	DeleteFirst() bool
	DeleteLast() bool
	DeleteAt(i int) bool
	GetFirst() (*Node[T], bool)
	GetLast() (*Node[T], bool)
	GetItem(i int) (*Node[T], bool)
	UpdateItem(i int, value T) bool
	IsEmpty() bool

	// value_n_from_end(n) - returns the value of the node at the nth position from the end of the list
	// reverse() - reverses the list
	// delete, insert multiple
}

func newNode[T any](value T, next *Node[T], prev *Node[T]) *Node[T] {
	return &Node[T]{
		Value: value,
		Next:  next,
		Prev:  prev,
	}
}

func NewDoublyLinkedList[T any](values ...T) *DoublyLinkedList[T] {
	var zero T

	head := newNode[T](zero, nil, nil)
	tail := newNode[T](zero, nil, nil)

	if len(values) == 0 {
		head.Next = tail
		tail.Prev = head
	} else {
		// create initial nodes
		nodes := make([]*Node[T], 0)
		for _, v := range values {
			nodes = append(nodes, newNode(v, nil, nil))
		}

		// link them together
		for i := 0; i < len(nodes); i++ {
			if i == len(nodes)-1 {
				break
			}

			node1 := nodes[i]
			node2 := nodes[i+1]

			node1.Next = node2
			node2.Prev = node1
		}

		// link the list with head and tail
		nodes[0].Prev = head
		head.Next = nodes[0]

		nodes[len(nodes)-1].Next = tail
		tail.Prev = nodes[len(nodes)-1]
	}

	return &DoublyLinkedList[T]{
		Length: len(values),
		Head:   head,
		Tail:   tail,
	}
}

func (d *DoublyLinkedList[T]) InsertFirst(value T) {
	// create new node
	node := newNode(value, nil, nil)

	// link to head and the next node
	temp := d.Head.Next

	d.Head.Next = node
	node.Prev = d.Head

	node.Next = temp
	temp.Prev = node

	// update length
	d.Length++
}

func (d *DoublyLinkedList[T]) InsertLast(value T) {
	// create new node
	node := newNode(value, nil, nil)

	// link to tail and the next node
	temp := d.Tail.Prev

	d.Tail.Prev = node
	node.Next = d.Tail

	node.Prev = temp
	temp.Next = node

	// update length
	d.Length++
}

func (d *DoublyLinkedList[T]) InsertAt(value T, i int) bool {
	if d.IsEmpty() {
		return false
	}

	node := newNode(value, nil, nil)
	nextNode := d.Head.Next

	for j := 0; j < i; j++ {
		nextNode = nextNode.Next
	}

	prevNode := nextNode.Prev

	prevNode.Next = node
	node.Prev = prevNode

	nextNode.Prev = node
	node.Next = nextNode

	d.Length++

	return true
}

func (d *DoublyLinkedList[T]) DeleteFirst() bool {
	if d.IsEmpty() {
		return false
	}

	nodeToDelete := d.Head.Next
	newFirstNode := nodeToDelete.Next

	// link to head and the next node
	d.Head.Next = newFirstNode
	newFirstNode.Prev = d.Head

	// update length
	d.Length--

	return true
}

func (d *DoublyLinkedList[T]) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}

	nodeToDelete := d.Tail.Prev
	newLastNode := nodeToDelete.Prev

	// link to head and the next node
	d.Tail.Prev = newLastNode
	newLastNode.Next = d.Tail

	// update length
	d.Length--

	return true
}

func (d *DoublyLinkedList[T]) DeleteAt(i int) bool {
	if d.IsEmpty() {
		return false
	}

	nodeToDelete := d.Head.Next
	for j := 0; j < i; j++ {
		nodeToDelete = nodeToDelete.Next
	}

	newPrevNode := nodeToDelete.Prev
	newNextNode := nodeToDelete.Next

	// link to head and the next node
	newPrevNode.Next = newNextNode
	newNextNode.Prev = newPrevNode

	// update length
	d.Length--

	return true
}

func (d *DoublyLinkedList[T]) GetFirst() (*Node[T], bool) {
	if d.IsEmpty() {
		return nil, false
	}

	return d.Head.Next, true
}

func (d *DoublyLinkedList[T]) GetLast() (*Node[T], bool) {
	if d.IsEmpty() {
		return nil, false
	}

	return d.Tail.Prev, true
}

func (d *DoublyLinkedList[T]) GetItem(i int) (*Node[T], bool) {
	if d.IsEmpty() {
		return nil, false
	}

	node := d.Head.Next
	for j := 0; j < i; j++ {
		node = node.Next
	}

	return node, true
}

func (d *DoublyLinkedList[T]) UpdateItem(i int, value T) bool {
	if d.IsEmpty() {
		return false
	}

	node := d.Head.Next
	for j := 0; j < i; j++ {
		node = node.Next
	}

	node.Value = value

	return true
}

func (d *DoublyLinkedList[T]) IsEmpty() bool {
	return d.Length == 0
}
