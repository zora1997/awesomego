package main

import (
	"fmt"
	"reflect"
)

type abcc struct {
	name string
	id   string
	age  int64
}

type abc struct {
	id  string
	age int64
}

func main() {
	abcd := make([]*abcc, 0, 10)
	abd := make([]*abc, 0, 10)

	abcd = append(abcd, &abcc{
		name: "abcd1",
		id:   "abcd1",
		age:  1,
	})
	abcd = append(abcd, &abcc{
		name: "abcd2",
		id:   "abcd2",
		age:  2,
	})
	abcd = append(abcd, &abcc{
		name: "abcd3",
		id:   "abcd3",
		age:  3,
	})
	abcd = append(abcd, &abcc{
		name: "abcd2",
		id:   "abcd2",
		age:  2,
	})

	abd = append(abd, &abc{
		id:  "abc1",
		age: 11,
	})
	abd = append(abd, &abc{
		id:  "abc2",
		age: 22,
	})
	abd = append(abd, &abc{
		id:  "abc2",
		age: 22,
	})
	abd = append(abd, &abc{
		id:  "abc3",
		age: 33,
	})
	fmt.Printf("first type: %T\n", abd)

	unDupTmp, dupTmp := unDupItems(abd)

	unDup := unDupTmp.([]*abc)
	dup := dupTmp.(map[string]struct{})

	fmt.Printf("%d unDup: %+v, type: %+v\n", len(unDup), unDup, typeof(unDup))
	fmt.Printf(" %d dup: %+v, type: %+v\n", len(dup), dup, typeof(dup))
}

func unDupItems(items interface{}) (interface{}, interface{}) {
	dupHelper := make(map[string]struct{})
	var dupItems interface{}

	// 注意这里直接使用x.(T)，如果不匹配会导致 panic

	switch items.(type) {
	case []*abcc:
		items := items.([]*abcc)
		result := make([]*abcc, 0, len(items))
		for _, item := range items {
			// 避免panic
			if item == nil {
				continue
			}

			if !checkDup(dupHelper, item.id) {
				result = append(result, item)
			}
		}
		return result, dupItems
	case []*abc:
		items := items.([]*abc)
		result := make([]*abc, 0, len(items))
		dupItems := make(map[string]struct{}, 0)
		for _, item := range items {
			// 避免panic
			if item == nil {
				continue
			}

			if !checkDup(dupHelper, item.id) {
				result = append(result, item)
			} else {
				dupItems[item.id] = struct{}{}
			}
		}
		return result, dupItems
	default:
		fmt.Println("unrecognized type")
	}

	return nil, nil
}

func checkDup(tmp map[string]struct{}, id string) bool {
	if _, ok := tmp[id]; !ok {
		tmp[id] = struct{}{}
		return false
	}
	return true
}

func typeof(v interface{}) string {
	return fmt.Sprintln(reflect.TypeOf(v).String())
}
