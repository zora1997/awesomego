package main

import (
	"fmt"
	"strings"
)

const (
	strEx = "srcarea 'untrust' desarea 'trust' src 'A1 A2' des 'B1' log on "
)

func main() {
	parsedOne := strings.Split(strEx, " ")
	for i := range parsedOne {
		fmt.Printf("%d param: %+v\n", i, parsedOne[i])
	}
}
