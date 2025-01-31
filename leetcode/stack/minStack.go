package stack

import (
	"fmt"
)

type MinStack struct {
	min   Stack[int]
	stack Stack[int]
}

func NewMinStack() *MinStack {
	return &MinStack{
		min:   make([]int, 0),
		stack: make([]int, 0),
	}
}

func (s *MinStack) Push(value int) {
	// if value < min, append to min stack
	// append to slice

	if minTop, isMinTopExist := s.min.Peek(); !isMinTopExist || value <= minTop {
		s.min = append(s.min, value)
	}

	s.stack = append(s.stack, value)
}

func (s *MinStack) Pop() {
	// if stack length === 0, return
	// pop
	// if popped el === min stack top, pop min stack

	popped, _ := s.stack.Peek()

	if minVal, _ := s.min.Peek(); popped == minVal {
		s.min.Pop()
	}

	s.stack.Pop()
}

func (s *MinStack) Top() (int, bool) {
	value, isExist := s.stack.Peek()
	return value, isExist
}

func (s *MinStack) GetMin() (int, bool) {
	minVal, isExist := s.min.Peek()
	return minVal, isExist
}

func MinStackExample() {
	minStack := NewMinStack()

	minStack.Push(4)
	minStack.Push(5)
	minStack.Push(6)
	minStack.Push(6)
	minStack.Push(5)
	minStack.Push(4)
	minStack.Push(3)
	minStack.Push(1)
	minStack.Push(2)

	fmt.Println("Min stack", minStack)

	minStack.Pop()
	minStack.Pop()

	fmt.Println("Min stack", minStack)

	minStack.Pop()
	minStack.Pop()

	fmt.Println("Min stack", minStack)

	minStack.Push(3)
	minStack.Push(1)
	minStack.Push(2)

	fmt.Println("Min stack", minStack)
}

func testTop(minStack *MinStack, testNo int, expect int) {
	top, isExist := minStack.Top()

	if len(minStack.stack) == 0 && isExist {
		panic(fmt.Sprintf("Test %d failed, Top() returning %t instead of %t\n", testNo, isExist, false))
	}

	if top != expect {
		panic(fmt.Sprintf("Test %d failed, Top() returning %d instead of %d\n", testNo, top, expect))
	}
}

func testMin(minStack *MinStack, testNo int, expect int) {
	minVal, isExist := minStack.GetMin()

	if len(minStack.min) == 0 && isExist {
		panic(fmt.Sprintf("Test %d failed, GetMin() returning %t instead of %t\n", testNo, isExist, false))
	}

	if minVal != expect {
		panic(fmt.Sprintf("Test %d failed, GetMin() returning %d instead of %d\n", testNo, minVal, expect))
	}
}

func testCase1() {
	minStack := NewMinStack()

	minStack.Push(4)
	minStack.Push(5)
	minStack.Push(6)
	minStack.Push(6)
	minStack.Push(5)
	minStack.Push(4)
	minStack.Push(3)
	minStack.Push(1)
	minStack.Push(2)

	testTop(minStack, 1, 2)
	testMin(minStack, 1, 1)

	minStack.Pop()
	minStack.Pop()

	testTop(minStack, 1, 3)
	testMin(minStack, 1, 3)

	minStack.Pop()
	minStack.Pop()

	testTop(minStack, 1, 5)
	testMin(minStack, 1, 4)

	minStack.Push(3)
	minStack.Push(1)
	minStack.Push(2)

	testTop(minStack, 1, 2)
	testMin(minStack, 1, 1)

	// empty the stack

	i := 0
	for i < 8 {
		minStack.Pop()
		i++
	}

	testMin(minStack, 1, 0)
	testTop(minStack, 1, 0)

	// should not do anything
	minStack.Pop()
	minStack.Top()

	fmt.Println(minStack)

	fmt.Printf("Test %d passed\n", 1)
}

func TestMinStack() {
	testCase1()
}
