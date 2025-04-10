package tree

import (
	"testing"
)

var bst *BinarySearchTree[int]

func setupBSTTest() {
	bst = NewBST(10)
}

func teardownBSTTest() {
	bst = nil
}

func TestBinarySearchTree_Clear(t *testing.T) {
	setupBSTTest()

	teardownBSTTest()
}

func TestBinarySearchTree_IsEmpty(t *testing.T) {
	setupBSTTest()

	teardownBSTTest()
}

func TestBinarySearchTree_Contains(t *testing.T) {
	setupBSTTest()

	teardownBSTTest()
}

func TestBinarySearchTree_Find(t *testing.T) {
	setupBSTTest()

	teardownBSTTest()
}

func TestBinarySearchTree_FindMax(t *testing.T) {
	setupBSTTest()

	teardownBSTTest()
}

func TestBinarySearchTree_FindMin(t *testing.T) {
	setupBSTTest()

	teardownBSTTest()
}

func TestBinarySearchTree_Insert(t *testing.T) {
	setupBSTTest()
	defer teardownBSTTest()

	// Insert right
	insertRight := 15
	if err := bst.Insert(insertRight); err != nil {
		t.Errorf("Unexpected error inserting %d: %v", insertRight, err)
	}
	if bst.Root.R == nil || bst.Root.R.Value != insertRight {
		t.Errorf("Expected %d on the right of root", insertRight)
	}

	// Insert left
	insertLeft := 5
	if err := bst.Insert(insertLeft); err != nil {
		t.Errorf("Unexpected error inserting %d: %v", insertLeft, err)
	}
	if bst.Root.L == nil || bst.Root.L.Value != insertLeft {
		t.Errorf("Expected %d on the left of root", insertLeft)
	}

	// Insert deeper left
	insertLeftDeeper := 2
	if err := bst.Insert(insertLeftDeeper); err != nil {
		t.Errorf("Unexpected error inserting %d: %v", insertLeftDeeper, err)
	}
	if bst.Root.L.L == nil || bst.Root.L.L.Value != insertLeftDeeper {
		t.Errorf("Expected %d to be left child of node %d", insertLeftDeeper, insertLeft)
	}

	// Insert deeper right
	insertRightDeeper := 17
	if err := bst.Insert(insertRightDeeper); err != nil {
		t.Errorf("Unexpected error inserting %d: %v", insertRightDeeper, err)
	}
	if bst.Root.R.R == nil || bst.Root.R.R.Value != insertRightDeeper {
		t.Errorf("Expected %d to be right child of node %d", insertRightDeeper, insertRight)
	}

	// Insert into middle layer
	insertRightLeft := 12
	if err := bst.Insert(insertRightLeft); err != nil {
		t.Errorf("Unexpected error inserting %d: %v", insertRightLeft, err)
	}
	if bst.Root.R.L == nil || bst.Root.R.L.Value != insertRightLeft {
		t.Errorf("Expected %d to be left child of node %d", insertRightLeft, insertRight)
	}

	// Insert duplicate
	if err := bst.Insert(insertLeft); err == nil {
		t.Errorf("Expected error when inserting duplicate %d", insertLeft)
	}
}

func TestBinarySearchTree_Delete(t *testing.T) {
	setupBSTTest()

	teardownBSTTest()
}
