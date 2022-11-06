package main

import (
	"fmt"
)

func main() {
	sliceA := []string{"1", "2", "3", "4", "5"}
	sliceB := []string{"6", "2", "3", "4", "7"}

	result := unDupTwoItems(sliceA, sliceB, 4)
	fmt.Printf("result unDup: %s \n", result)
}

// 取并集，去重
func unDupTwoItems(itemA, itemB []string, num uint32) []string {
	if uint32(len(itemA)) >= num {
		return itemA[:num]
	}

	if len(itemA) == 0 {
		return itemB
	}

	if len(itemB) == 0 {
		return itemA
	}

	mHash := make(map[string]struct{})
	result := make([]string, 0)

	for i := range itemA {
		result = append(result, itemA[i])
		mHash[itemA[i]] = struct{}{}
	}

	for j := range itemB {
		if uint32(len(result)) >= num {
			break
		}
		if _, ok := mHash[itemB[j]]; ok {
			continue
		}
		result = append(result, itemB[j])
	}

	return result
}
