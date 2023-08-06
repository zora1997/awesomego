package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4}
	// slice2 := []int{}
	var slice3 []int
	slice3 = nil

	slice1 = append(slice1, slice3...)
	fmt.Printf("slice1 after 3: %+v\n", slice1)

	/*
		fmt.Printf("slice1 ori: %+v\n", slice1)
		fmt.Printf("slice2: %+v\n", slice2)

		// slice1 = append(slice1, slice2...)
		// fmt.Printf("slice1 final: %+v\n", slice1)
		slice2 = append(slice2, slice1...)
		fmt.Printf("slice2 final: %+v\n", slice2)

	*/
}
