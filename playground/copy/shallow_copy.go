package main

import "fmt"

type zora struct {
	Name     string
	Age      int
	MapExtra map[string]string
}

func main() {
	tmp := make(map[string]string)
	tmp["1"] = "1"
	z1 := zora{
		Name:     "zora",
		Age:      1,
		MapExtra: tmp,
	}
	z2 := ShallowCopyZora(z1)
	fmt.Printf("z1: %+v, z2:%+v\n", z1, z2)
}

// 这里需要注意浅拷贝带来的对地址的修改，
func ShallowCopyZora(z1 zora) zora {
	z2 := z1
	tmp := make(map[string]string)
	for k, v := range z1.MapExtra {
		tmp[k] = v + "2"
	}
	z2.Name = "zzora" + "_" + "z"
	z2.MapExtra = tmp
	return z2
}
