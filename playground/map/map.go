package main

import "fmt"

type abc struct {
	name string
	age  int
}

func main() {
	mapTest := make(map[string][]*abc)
	mapTest["a"] = append(mapTest["a"], &abc{
		name: "a",
		age:  1,
	})
	mapTest["a"] = append(mapTest["a"], &abc{
		name: "a",
		age:  11,
	})
	mapTest["a"] = append(mapTest["a"], &abc{
		name: "a",
		age:  111,
	})
	mapTest["b"] = append(mapTest["b"], &abc{
		name: "b",
		age:  2,
	})
	mapTest["b"] = append(mapTest["b"], &abc{
		name: "b",
		age:  22,
	})
	mapTest["c"] = append(mapTest["c"], &abc{
		name: "c",
		age:  3,
	})

	keys := make([]string, 0)
	for i := range mapTest {
		keys = append(keys, i)
	}
	fmt.Printf("keys: %+v\n", keys)

	for k, v := range mapTest {
		for i := range v {
			fmt.Printf("key: %s, v: %+v\n", k, v[i])
		}
	}
}
