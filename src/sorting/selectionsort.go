package main

import "fmt"

/*
  Go program for selection sort
*/

// Swap the array element
func swap(arr []int, x int, y int) {
	// x and y are index of array
	var temp int = arr[x]
	arr[x] = arr[y]
	arr[y] = temp
}
func selectionSort(arr []int, n int) {
	var min int = 0
	// Execute loop from 0..n
	for i := 0; i < n; i++ {
		// Get current index
		min = i
		for j := i + 1; j < n; j++ {
			if arr[min] > arr[j] {
				// Get the minimum element index
				min = j
			}
		}
		if i != min {
			// Swap minimum element at i index
			swap(arr, i, min)
		}
	}
}

// Display array elements
func display(arr []int, n int) {
	for i := 0; i < n; i++ {
		// Display element value
		fmt.Print("  ", arr[i])
	}
	fmt.Print("\n")
}
func main() {

	// Array of integer elements
	var arr = []int{
		8,
		2,
		3,
		8,
		1,
		3,
		73,
		121,
		54,
		23,
		84,
		13,
		67,
		23,
		52,
	}
	// Get the size of array
	var n int = len(arr)
	fmt.Println(" Before Sort :")
	display(arr, n)
	// Test
	selectionSort(arr, n)
	fmt.Println(" After Sort :")
	display(arr, n)
}
