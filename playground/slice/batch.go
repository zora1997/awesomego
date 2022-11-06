package main

import (
	"fmt"
)

func main() {
	slice1 := []uint32{1, 2, 3}
	slice2 := []uint32{4, 5, 6, 7, 8, 9, 10, 11}
	all := 12
	b1 := 1
	b2 := 4

	if all < b1+b2 {
		fmt.Println("\"all\" wrong")
	}

	cards := make([]uint32, 0, all)
	s1 := make([]uint32, 0, b1)
	s2 := make([]uint32, 0, b2)

	for len(slice1) != 0 && len(slice2) != 0 {
		if len(slice1) > b1 {
			s1 = slice1[:b1]
			// 出队
			slice1 = slice1[b1:]
		} else {
			s1 = slice1
		}
		fmt.Printf("s1: %v", s1)

		if len(slice2) > b2 {
			s2 = slice2[:b2]
			// 出队
			slice2 = slice2[:b2]
		} else {
			s2 = slice2
		}
		fmt.Printf("s2: %v", s2)

		cards = append(cards, s1...)
		fmt.Printf("cards1: %v", cards)
		cards = append(cards, s2...)
		fmt.Printf("cards2: %v", cards)
	}
}
