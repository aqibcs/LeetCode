package main

import "fmt"

func main() {
	// Create an empty slice
	var dynamicArray []int

	// Append elements to the dynamic array
	dynamicArray = append(dynamicArray, 1)
	dynamicArray = append(dynamicArray, 2, 3, 4)

	// Display the dynamic array
	fmt.Println("Dynamic Array:", dynamicArray)

	// Access elements by index
	fmt.Println("Element at index 1:", dynamicArray[1])

	// Update an element
	dynamicArray[1] = 99
	fmt.Println("Updated Dynamic Array:", dynamicArray)

	// length and capacity of the slice 
	fmt.Println("Length:", len(dynamicArray))
	fmt.Println("Capacity;", cap(dynamicArray))

	// Creating a slice with make function
	anotherArray := make([]int, 5, 10)
	fmt.Println("Another Array:", anotherArray)
	fmt.Println("length:", len(anotherArray))
	fmt.Println("Capacity:", cap(anotherArray))

	// Appending to the slice with elements
	anotherArray = append(anotherArray, 6, 7, 8)
	fmt.Println("Updated Another Array:", anotherArray)
	fmt.Println("Length:", len(anotherArray))
	fmt.Println("Capacity:", cap(anotherArray))
}
