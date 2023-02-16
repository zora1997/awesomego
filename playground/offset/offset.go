package main

import "fmt"

type Req struct {
	id     string
	offset uint64
}

func main() {
	req := &Req{}
	req.offset = 3
	fmt.Println(Offset(req.offset == 0))
}

func Offset(ifZero bool) string {
	switch ifZero {
	case true:
		return "0 offset"
	default:
		return "default"
	}
}
