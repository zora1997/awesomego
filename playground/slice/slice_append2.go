package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}

	var tmp []int
	for i := range slice1 {
		if slice1[i]%2 != 0 {
			continue
		}
		tmp = append(tmp, slice1[i])
	}

	slice1 = tmp
	fmt.Printf("slice1: %+v\n", slice1)
	// slice1: [2 4]
}
