package ds

import (
	"testing"
)

// Helper function to check node values
func assertNodeValue[T comparable](t *testing.T, node *Node[T], expected T) {
	if node == nil {
		t.Errorf("Expected node value %v, but got nil", expected)
	} else if node.Value != expected {
		t.Errorf("Expected node value %v, but got %v", expected, node.Value)
	}
}

// ✅ Test list initialization
func TestNewDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList[int](1, 2, 3)

	if list.Length != 3 {
		t.Errorf("Expected length 3, got %d", list.Length)
	}

	assertNodeValue(t, list.Head.Next, 1)
	assertNodeValue(t, list.Tail.Prev, 3)
}

// ✅ Test InsertFirst
func TestInsertFirst(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.InsertFirst(10)
	list.InsertFirst(20)

	assertNodeValue(t, list.Head.Next, 20)
	assertNodeValue(t, list.Head.Next.Next, 10)

	if list.Length != 2 {
		t.Errorf("Expected length 2, got %d", list.Length)
	}
}

// ✅ Test InsertLast
func TestInsertLast(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.InsertLast(10)
	list.InsertLast(20)

	assertNodeValue(t, list.Tail.Prev, 20)
	assertNodeValue(t, list.Tail.Prev.Prev, 10)

	if list.Length != 2 {
		t.Errorf("Expected length 2, got %d", list.Length)
	}
}

// ✅ Test InsertAt
func TestInsertAt(t *testing.T) {
	list := NewDoublyLinkedList[int](1, 3)
	success := list.InsertAt(2, 1)

	if !success {
		t.Errorf("InsertAt failed")
	}

	assertNodeValue(t, list.Head.Next.Next, 2)
	assertNodeValue(t, list.Head.Next.Next.Next, 3)

	if list.Length != 3 {
		t.Errorf("Expected length 3, got %d", list.Length)
	}
}

// ✅ Test DeleteFirst
func TestDeleteFirst(t *testing.T) {
	list := NewDoublyLinkedList[int](1, 2, 3)
	list.DeleteFirst()

	assertNodeValue(t, list.Head.Next, 2)

	if list.Length != 2 {
		t.Errorf("Expected length 2, got %d", list.Length)
	}
}

// ✅ Test DeleteLast
func TestDeleteLast(t *testing.T) {
	list := NewDoublyLinkedList[int](1, 2, 3)
	list.DeleteLast()

	assertNodeValue(t, list.Tail.Prev, 2)

	if list.Length != 2 {
		t.Errorf("Expected length 2, got %d", list.Length)
	}
}

// ✅ Test DeleteAt
func TestDeleteAt(t *testing.T) {
	list := NewDoublyLinkedList[int](1, 2, 3)
	success := list.DeleteAt(1)

	if !success {
		t.Errorf("DeleteAt failed")
	}

	assertNodeValue(t, list.Head.Next.Next, 3)

	if list.Length != 2 {
		t.Errorf("Expected length 2, got %d", list.Length)
	}
}

// ✅ Test GetFirst
func TestGetFirst(t *testing.T) {
	list := NewDoublyLinkedList[int](10, 20, 30)

	firstNode, exists := list.GetFirst()
	if !exists {
		t.Errorf("Expected node, but got nil")
	}
	assertNodeValue(t, firstNode, 10)
}

// ✅ Test GetLast
func TestGetLast(t *testing.T) {
	list := NewDoublyLinkedList[int](10, 20, 30)

	lastNode, exists := list.GetLast()
	if !exists {
		t.Errorf("Expected node, but got nil")
	}
	assertNodeValue(t, lastNode, 30)
}

// ✅ Test GetItem
func TestGetItem(t *testing.T) {
	list := NewDoublyLinkedList[int](10, 20, 30)

	node, exists := list.GetItem(1)
	if !exists {
		t.Errorf("Expected node at index 1, but got nil")
	}
	assertNodeValue(t, node, 20)
}

// ✅ Test UpdateItem
func TestUpdateItem(t *testing.T) {
	list := NewDoublyLinkedList[int](10, 20, 30)
	success := list.UpdateItem(1, 25)

	if !success {
		t.Errorf("UpdateItem failed")
	}

	node, _ := list.GetItem(1)
	assertNodeValue(t, node, 25)
}

// ✅ Test IsEmpty
func TestIsEmpty(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	if !list.IsEmpty() {
		t.Errorf("Expected empty list, but got non-empty")
	}

	list.InsertFirst(10)
	if list.IsEmpty() {
		t.Errorf("Expected non-empty list, but got empty")
	}
}

// ✅ **NEW: Test edge cases on an empty list**
func TestEmptyListOperations(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	// Deletion operations should fail
	if list.DeleteFirst() {
		t.Errorf("Expected DeleteFirst to return false on an empty list")
	}
	if list.DeleteLast() {
		t.Errorf("Expected DeleteLast to return false on an empty list")
	}
	if list.DeleteAt(0) {
		t.Errorf("Expected DeleteAt to return false on an empty list")
	}

	// Getting items should fail
	if node, exists := list.GetFirst(); exists || node != nil {
		t.Errorf("Expected GetFirst to return nil, got %v", node)
	}
	if node, exists := list.GetLast(); exists || node != nil {
		t.Errorf("Expected GetLast to return nil, got %v", node)
	}
	if node, exists := list.GetItem(0); exists || node != nil {
		t.Errorf("Expected GetItem to return nil, got %v", node)
	}

	// Updating should fail
	if list.UpdateItem(0, 100) {
		t.Errorf("Expected UpdateItem to return false on an empty list")
	}

	// Confirm list is empty
	if !list.IsEmpty() {
		t.Errorf("Expected list to be empty")
	}
}
