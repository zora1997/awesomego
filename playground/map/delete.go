package main

import "fmt"

func main() {

	mp := make(map[string]int)
	mp["a"] = 1
	mp["b"] = 2
	mp["c"] = 3

	for k, v := range mp {
		fmt.Printf("before delete, key: %s, value: %d\n", k, v)
	}

	delete(mp, "a")

	for k, v := range mp {
		fmt.Printf("after delete, key: %s, value: %d\n", k, v)
	}
}
