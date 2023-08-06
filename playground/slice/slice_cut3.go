package main

import "fmt"

func main() {
	slice1 := make([]int, 0, 20)
	fmt.Printf("slice: %+v\n", slice1[0:])
}
