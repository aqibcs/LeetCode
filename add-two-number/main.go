package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil {
		x := 0
		y := 0

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := x + y + carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next
}

// Helper function to create a linked list from an array of integers
func createLinkedList(nums []int) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead

	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}

	return dummyHead.Next
}

// Helper function to print a linked list
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example 1
	l1 := createLinkedList([]int{2, 4, 3})
	l2 := createLinkedList([]int{5, 6, 4})
	result1 := addTwoNumbers(l1, l2)
	fmt.Print("Example 1: ")
	printLinkedList(result1)

	// Example 2
	l3 := createLinkedList([]int{0})
	l4 := createLinkedList([]int{0})
	result2 := addTwoNumbers(l3, l4)
	fmt.Print("Example 2: ")
	printLinkedList(result2)

	// Example 3
	l5 := createLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	l6 := createLinkedList([]int{9, 9, 9, 9})
	result3 := addTwoNumbers(l5, l6)
	fmt.Print("Example 3: ")
	printLinkedList(result3)
}
