package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{5, 6, 7, 8, 9, 10}
	// copy(slice1, slice2)
	slice2 = []int{}
	slice2 = slice1
	// copy(slice2, slice1)

	fmt.Printf("slice1: %+v\n", slice1)
	fmt.Printf("slice2: %+v\n", slice2)
}
