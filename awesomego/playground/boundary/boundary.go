package main

import "fmt"

type st struct {
	num uint32
	tp  string
}

func main() {
	size := 4
	ex := make([]*st, 0, size)
	ex = append(ex, &st{
		num: 1,
		tp:  "b",
	},
		&st{
			num: 2,
			tp:  "a",
		},
		&st{
			num: 3,
			tp:  "1",
		},
		&st{
			num: 4,
			tp:  "b",
		},
		&st{
			num: 5,
			tp:  "a",
		},
		&st{
			num: 6,
			tp:  "b",
		},
		&st{
			num: 7,
			tp:  "b",
		},
		&st{
			num: 8,
			tp:  "b",
		},
		&st{
			num: 9,
			tp:  "b",
		})

	resA := make([]uint32, 0, size)
	resB := make([]uint32, 0, size)

	for i := range ex {
		if len(resA) < size {
			if ex[i].tp == "a" {
				resA = append(resA, ex[i].num)
			}
		}

		if len(resB) < size {
			if ex[i].tp == "b" {
				resB = append(resB, ex[i].num)
			}
		}
	}
	fmt.Printf("resA: %+v\n", resA)
	fmt.Printf("resB: %+v\n", resB)
}
