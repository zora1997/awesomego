package main

import "fmt"

func main() {
	tmp1 := make([]string, 0)
	tmp2 := make([]string, 0)
	res := make([]string, 0)

	tmp2 = append(tmp2, "tmp2", "tmp22", "tmp222")

	res = append(tmp1, tmp2...)
	fmt.Printf("%d res: %+v\n", len(res), res)
}
