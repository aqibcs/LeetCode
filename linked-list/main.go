package main

import "fmt"

// Node represents a node in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represent a singly linked list.
type LinkedList struct {
	head *Node
}

// Append adds a new node with the given data to the end of the linked list.
func (ll *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	lastNode := ll.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	lastNode.next = newNode
}

// Display prints the elements of the linked List.
func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	myList := LinkedList{}

	// Append elements to the linked list
	myList.Append(1)
	myList.Append(2)
	myList.Append(3)

	// Display the linked list
	fmt.Println("Linked List:")
	myList.Display()
}
