package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := []int{11, 12, 13, 12}
	slice1 = slice2
	fmt.Printf("slice1: %+v\n", slice1)

}
