package main

import "fmt"

func main() {
	/*
		a := make([]string, 0, 5)
		a = nil
		fmt.Println("nil len: %d", len(a))
	*/

	var items []string
	fmt.Printf("items: %+v\n", items)
	items = append(items, "1")
	fmt.Printf("items: %+v\n", items)
}
