package main

import "fmt"

type testStruct struct {
	uid  uint64
	name string
}

func main() {

}

func Data() (orign []*testStruct) {
	dataOne := &testStruct{
		uid:  111,
		name: "dataOne",
	}

	dataTwo := &testStruct{
		uid:  222,
		name: "dataTwo",
	}

	dataThree := &testStruct{
		uid:  333,
		name: "dataThree",
	}

	orign = append(orign, dataOne, dataTwo, dataThree)

	return orign
}

func DeepCopy() (result []*testStruct) {
	origin := Data()
	for i := range origin {
		result = append(result, &testStruct{
			uid:  origin[i].uid,
			name: origin[i].name,
		})
	}
	fmt.Printf("result: %v", result)

	return result
}

func ShallowCopy() (result []*testStruct) {
	origin := Data()
	result = origin
	fmt.Printf("result: %v", result)

	return result
}
