package main

import (
	"fmt"
	ds "leetcode/data_structure"
)

func main() {
	//fmt.Println("Two sum")
	//hashMap.TestTwoSum()
	//
	//fmt.Println("Contains duplicate")
	//hashMap.TestContainsDuplicate()
	//
	//fmt.Println("Valid anagrams")
	//hashMap.TestValidAnagrams()
	//
	//fmt.Println("Longest Consecutive Sequence")
	//hashMap.TestLongestConsecutiveSequence()

	//fmt.Println("Valid Parentheses")
	//stack.TestValidParentheses()

	//fmt.Println("Min stack impl")
	//stack.TestMinStack()

	//fmt.Println("Valid Palindrome")
	//two_pointer.TestValidPalindrome()

	fmt.Println("Doubly linked list")
	list := ds.NewDoublyLinkedList[int](1, 2, 3)
	list.InsertLast(9)
	fmt.Println(list, list.Head.Next.Next.Next.Next)
}
