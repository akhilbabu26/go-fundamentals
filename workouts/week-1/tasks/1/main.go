package main

import "fmt"

// Insert element at given index
func insertAtIndex(arr []int, index int, value int) []int {
	if index < 0 || index > len(arr) {
		fmt.Println("Invalid index for insertion")
		return arr
	}

	// Increase slice size by 1
	arr = append(arr, 0)

	// Shift elements to the right
	for i := len(arr) - 1; i > index; i-- {
		arr[i] = arr[i-1]
	}

	// Insert value
	arr[index] = value

	return arr
}

// Delete element at given index
func deleteAtIndex(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		fmt.Println("Invalid index for deletion")
		return arr
	}

	// Shift elements to the left
	for i := index; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}

	// Remove last element
	arr = arr[:len(arr)-1]

	return arr
}

func main() {
	arr := []int{10, 20, 30, 40, 50}

	fmt.Println("Original array:", arr)

	// Insert 99 at index 2
	arr = insertAtIndex(arr, 2, 99)
	fmt.Println("After insertion:", arr)

	// Delete element at index 3
	arr = deleteAtIndex(arr, 3)
	fmt.Println("After deletion:", arr)
}
