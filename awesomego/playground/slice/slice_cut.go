package main

import "fmt"

func main() {
	sl1 := []uint32{1, 2, 3, 4, 5}
	sl2 := []uint32{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	index1 := 0
	index2 := 0
	b1 := 2
	b2 := 3
	res := make([]uint32, 0)

	for index1 < len(sl1) || index2 < len(sl2) {
		if index1 < len(sl1) {
			fmt.Printf("index1 before: %d\n", index1)
			if index1+b1 <= len(sl1) {
				res = append(res, sl1[index1:index1+b1]...)
			} else {
				res = append(res, sl1[index1:]...)
			}
			index1 += b1
			fmt.Printf("index1 after: %d\n", index1)
		}

		if index2 < len(sl2) {
			fmt.Printf("index2 before: %d\n", index2)
			if index2+b2 <= len(sl2) {
				res = append(res, sl2[index2:index2+b2]...)
			} else {
				res = append(res, sl2[index2:]...)
			}
			index2 += b2
			fmt.Printf("index2 after: %d\n", index2)
		}
		fmt.Printf("%d res: %+v\n", len(res), res)
	}
}
