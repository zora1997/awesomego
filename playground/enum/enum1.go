package main

import (
	"fmt"
	"github.com/spf13/cast"
)

type Code int

const Invalid Code = 3

func main() {
	fmt.Println(cast.ToString(Invalid))
}
