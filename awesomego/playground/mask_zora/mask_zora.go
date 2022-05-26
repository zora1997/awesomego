package main

import (
	"fmt"
	"strconv"
)

// 获取所有开关默认对应规则
func main() {
	rule := make(map[string]uint64, 0)
	rule["0x1"] = 1
	rule["0x2"] = 1
	rule["0x4"] = 0
	rule["0x8"] = 0
	rule["0x16"] = 1

	var result uint64 = 0
	// v为1，每一位1代表关，0代表开
	// v为0，每一位1代表开，0代表关
	for k, v := range rule {
		// 如果v为1，就没必要计算mask和做或运算，因此continue
		if v == 1 {
			fmt.Printf("%s mask's value: %d\n", k, v)
			continue
		}

		mask, err := strconv.ParseUint(k, 0, 64)
		if err != nil {
			fmt.Println("fail to parse mask")
			continue
		}

		result |= mask
	}
	fmt.Printf("result: %d!", result)
}
