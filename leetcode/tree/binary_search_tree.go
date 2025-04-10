package tree

import (
	"errors"
	"log"
)

// TODO: Self balancing

type ECompareResult int8

const (
	Less                   ECompareResult = -1
	Equal                  ECompareResult = 0
	More                   ECompareResult = 1
	NotFoundErrorMsg       string         = "the value you're trying to find does not exist"
	NotFoundDeleteErrorMsg string         = "the value you're trying to delete does not exist"
	ExistedErrorMsg        string         = "the value you're trying to insert already exists"
)

type BSTNode[T int] struct {
	Value T
	L     *BSTNode[T]
	R     *BSTNode[T]
}

func NewBSTNode[T int](val T) *BSTNode[T] {
	return &BSTNode[T]{
		Value: val,
		L:     nil,
		R:     nil,
	}
}

type BinarySearchTree[T int] struct {
	Root      *BSTNode[T]
	compareFn func(val1 T, val2 T) ECompareResult
}

func NewBST[T int](rootVal T) *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		Root: NewBSTNode(rootVal),
	}
}

func (b *BinarySearchTree[T]) compare(val1 T, val2 T) ECompareResult {
	if b.compareFn != nil {
		return b.compareFn(val1, val2)
	}

	if val1 < val2 {
		return Less
	}

	if val1 > val2 {
		return More
	}

	return Equal
}

func (b *BinarySearchTree[T]) Clear() {
	b.Root.L = nil
	b.Root.R = nil
}

func (b *BinarySearchTree[T]) IsEmpty() bool {
	if b.Root.L == nil && b.Root.R == nil {
		return true
	}

	return false
}

func (b *BinarySearchTree[T]) Contains(val T) bool {
	_, err := b.search(val)

	if err != nil {
		return false
	}

	return true
}

func (b *BinarySearchTree[T]) Find(val T) (T, error) {
	node, err := b.search(val)
	var zero T

	if err != nil {
		return zero, err
	}

	return node.Value, nil
}

func (b *BinarySearchTree[T]) FindMinNodeFrom(node *BSTNode[T]) *BSTNode[T] {
	// go left most
	temp := node

	for {
		if temp.L == nil {
			return temp
		}

		temp = temp.L
	}
}

func (b *BinarySearchTree[T]) FindMin() T {
	// go left most
	node := b.Root
	return b.FindMinNodeFrom(node).Value
}

func (b *BinarySearchTree[T]) FindMaxNodeFrom(node *BSTNode[T]) *BSTNode[T] {
	// go right most
	temp := node

	for {
		if temp.R == nil {
			return temp
		}

		temp = temp.R
	}
}

func (b *BinarySearchTree[T]) FindMax() T {
	// go right most
	node := b.Root
	return b.FindMaxNodeFrom(node).Value
}

func (b *BinarySearchTree[T]) Insert(val T) error {
	// TODO: Self balancing

	node := b.Root

	for {
		compareRes := b.compare(val, node.Value)

		if compareRes == Equal {
			return errors.New(ExistedErrorMsg)
		}

		if compareRes == Less {
			if node.L == nil {
				node.L = NewBSTNode(val)

				return nil
			}

			node = node.L
			continue
		}

		if compareRes == More {
			if node.R == nil {
				node.R = NewBSTNode(val)

				return nil
			}

			node = node.R
			continue
		}
	}
}

func (b *BinarySearchTree[T]) Delete(val T) error {
	// TODO: Self balancing

	node := b.Root

	for {
		compareRes := b.compare(val, node.Value)

		if compareRes == Equal {
			if b.isLeaf(node) {
				node = nil

				return nil
			}

			// largest node on left subtree
			subTreeRoot := node.L
			newNode := b.FindMaxNodeFrom(subTreeRoot)
			node.Value = newNode.Value
			newNode = nil

			return nil
		}

		if compareRes == Less {
			if node.L == nil {
				return errors.New(NotFoundDeleteErrorMsg)
			}

			node = node.L
			continue
		}

		if compareRes == More {
			if node.R == nil {
				return errors.New(NotFoundDeleteErrorMsg)
			}

			node = node.R
			continue
		}
	}
}

func (b *BinarySearchTree[T]) PrintTree() {
	log.Println("Tree")
}

func (b *BinarySearchTree[T]) search(val T) (*BSTNode[T], error) {
	node := b.Root

	for {
		compareRes := b.compare(val, node.Value)

		if compareRes == Equal {
			return node, nil
		}

		if compareRes == Less {
			if node.L == nil {
				return nil, errors.New(NotFoundErrorMsg)
			}

			node = node.L
			continue
		}

		if compareRes == More {
			if node.R == nil {
				return nil, errors.New(NotFoundErrorMsg)
			}

			node = node.R
			continue
		}
	}
}

func (b *BinarySearchTree[T]) isLeaf(node *BSTNode[T]) bool {
	if node.L == nil && node.R == nil {
		return true
	}

	return false
}
