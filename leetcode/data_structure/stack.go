package ds

import "fmt"

type IStack[T any] interface {
	IsEmpty() bool
	Push(el T) T
	Pop() (T, bool)
	Peek() (T, bool)
}

type Stack[T any] []T

func NewStack[T any](values ...T) Stack[T] {
	return Stack[T](values)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(el T) T {
	*s = append(*s, el)

	return el
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T

		return zero, false
	} else {
		index := len(*s) - 1
		poppedEl := (*s)[index]
		*s = (*s)[:index]

		return poppedEl, true
	}
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T

		return zero, false
	}

	index := len(*s) - 1
	lastEl := (*s)[index]

	return lastEl, true
}

func Example() {
	stack := NewStack[string]("0", "9", "8")

	stack.Push("1")
	stack.Push("2")
	stack.Push("3")

	fmt.Println("stack", stack)

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}

	fmt.Println("stack popped", stack)
}
