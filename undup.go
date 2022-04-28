package main

import (
	"fmt"
)

func main() {
	sliceC := make([]string, 0)
	sliceA := []string{"1", "2", "3", "4", "5"}
	sliceB := []string{"6", "2", "3", "4", "7"}
	sliceC = append(sliceC, sliceA...)
	sliceC = append(sliceC, sliceB...)
	fmt.Printf("sliceC dup: %s \n", sliceC)

	sliceC = unDupRoomIDs(sliceC)
	fmt.Printf("sliceC unDup: %s \n", sliceC)
}

func unDupRoomIDs(roomIDs []string) []string {
	result := make([]string, 0, len(roomIDs))
	tmp := make(map[string]struct{})

	for _, ele := range roomIDs {
		if _, ok := tmp[ele]; !ok {
			tmp[ele] = struct{}{}
			result = append(result, ele)
		}
	}

	return result
}
