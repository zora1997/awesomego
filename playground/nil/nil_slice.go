package main

import "fmt"

func main() {
	a := make([]string, 0, 5)
	a = nil
	fmt.Println("nil len: %d", len(a))
}
