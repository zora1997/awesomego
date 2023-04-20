package main

import "fmt"

func main() {
	// result := make([]string, 4)
	result := make([]string, 0, 4)
	fmt.Printf("%q, %q, %q, %q \n", result[0], result[1], result[2], result[3])
}
