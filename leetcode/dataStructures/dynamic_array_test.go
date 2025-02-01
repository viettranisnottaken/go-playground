package dataStructures

import (
	"testing"
)

func TestDynamicArray_InsertFirst(t *testing.T) {
	arr := NewDynamicArray[int]()

	arr.InsertFirst(10)
	arr.InsertFirst(20)

	if arr.currSize != 2 {
		t.Errorf("Expected size 2, got %d", arr.currSize)
	}

	if val, _ := arr.GetItem(0); val != 20 {
		t.Errorf("Expected first element to be 20, got %d", val)
	}
}

func TestDynamicArray_InsertLast(t *testing.T) {
	arr := NewDynamicArray[int]()
	arr.InsertLast(10)
	arr.InsertLast(20)

	if arr.currSize != 2 {
		t.Errorf("Expected size 2, got %d", arr.currSize)
	}

	if val, _ := arr.GetItem(1); val != 20 {
		t.Errorf("Expected last element to be 20, got %d", val)
	}
}

func TestDynamicArray_InsertAt(t *testing.T) {
	arr := NewDynamicArray[int](10, 20, 30)
	success := arr.InsertAt(15, 1)

	if !success {
		t.Errorf("InsertAt failed")
	}

	if arr.currSize != 4 {
		t.Errorf("Expected size 4, got %d", arr.currSize)
	}

	if val, _ := arr.GetItem(1); val != 15 {
		t.Errorf("Expected element at index 1 to be 15, got %d", val)
	}
}

func TestDynamicArray_DeleteFirst(t *testing.T) {
	arr := NewDynamicArray[int](10, 20, 30)
	success := arr.DeleteFirst()

	if !success {
		t.Errorf("DeleteFirst failed")
	}

	if arr.currSize != 2 {
		t.Errorf("Expected size 2, got %d", arr.currSize)
	}

	if val, _ := arr.GetItem(0); val != 20 {
		t.Errorf("Expected first element to be 20, got %d", val)
	}
}

func TestDynamicArray_DeleteLast(t *testing.T) {
	arr := NewDynamicArray[int](10, 20, 30)
	success := arr.DeleteLast()

	if !success {
		t.Errorf("DeleteLast failed")
	}

	if arr.currSize != 2 {
		t.Errorf("Expected size 2, got %d", arr.currSize)
	}

	if val, _ := arr.GetItem(1); val != 20 {
		t.Errorf("Expected last element to be 20, got %d", val)
	}
}

func TestDynamicArray_DeleteAt(t *testing.T) {
	arr := NewDynamicArray[int](10, 20, 30)
	success := arr.DeleteAt(1)

	if !success {
		t.Errorf("DeleteAt failed")
	}

	if arr.currSize != 2 {
		t.Errorf("Expected size 2, got %d", arr.currSize)
	}

	if val, _ := arr.GetItem(1); val != 30 {
		t.Errorf("Expected element at index 1 to be 30, got %d", val)
	}
}

func TestDynamicArray_GetItem(t *testing.T) {
	arr := NewDynamicArray[int](10, 20, 30)

	val, exists := arr.GetItem(1)
	if !exists || val != 20 {
		t.Errorf("Expected to get 20 at index 1, got %d", val)
	}

	_, exists = arr.GetItem(5)
	if exists {
		t.Errorf("Expected false for out-of-bounds index")
	}
}

func TestDynamicArray_Grow(t *testing.T) {
	arr := NewDynamicArray[int](10, 20, 30)
	initialCap := arr.capacity

	arr.InsertLast(40)
	arr.InsertLast(50)
	arr.InsertLast(60)
	arr.InsertLast(70) // This should trigger `grow()`

	if arr.capacity <= initialCap {
		t.Errorf("Expected capacity to grow, but it didn't")
	}
}

func TestDynamicArray_Shrink(t *testing.T) {
	arr := NewDynamicArray[int](10, 20, 30, 40)
	initialCap := arr.capacity

	arr.DeleteLast()
	arr.DeleteLast()
	arr.DeleteLast() // Should trigger `shrink()`

	if arr.capacity >= initialCap {
		t.Errorf("Expected capacity to shrink, but it didn't")
	}
}
