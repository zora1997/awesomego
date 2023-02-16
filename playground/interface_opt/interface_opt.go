package main

import (
	"fmt"
)

/*
type abcc struct {
	name string
	id   string
	age  int64
}

*/

func main() {
	var items []*abcc
	items = append(items, &abcc{
		name: "1",
		id:   "1",
		age:  1,
	})
	items = append(items, &abcc{
		name: "2",
		id:   "2",
		age:  2,
	})
	items = append(items, &abcc{
		name: "3",
		id:   "3",
		age:  3,
	})

	res := abc(items)
	tmp := res.([]*abcc)
	fmt.Println(typeof(res))
	fmt.Println(typeof(tmp))

	for i := range res {
		fmt.Printf("%d tmp: %+v\n", i, tmp[i])
	}

}

func abcddd(items interface{}) interface{} {
	var result interface{}
	tmp := make([]*abcc, 0)

	switch items.(type) {
	case []*abcc:
		item := items.([]*abcc)
		for i := range item {
			tmp = append(tmp, item[i])
			fmt.Printf("%d item: %+v\n", i, item[i])
		}
		result = tmp
	}

	return result
}

// 如果直接使用x.(T)进行断言，如果x不是T类型，那么则会出现panic错误，这显然是不够优雅的，所以建议尽可能的使用convert, ok := x.(T)或者switch + x.(type)的方式来进行类型断言。
