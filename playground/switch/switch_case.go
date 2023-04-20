package main

import "fmt"

func main() {
	isA := true
	isB := true
	switch {
	case isA:
		fmt.Printf("isA\n")
		fallthrough
	case isB:
		fmt.Printf("isB\n")
	}

	fmt.Printf("last\n")
}
