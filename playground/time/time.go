package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		timee := time.Now()
		fmt.Printf("timee: %d\n", timee.UnixMicro()/1e3)
		var updateTime int64
		var expire int64
		updateTime = 60
		expire = 20
		if time.Now().UnixNano()/int64(time.Millisecond)-updateTime < expire {
			fmt.Println("expire")
		} else {
			fmt.Println("not expire")
		}
		fmt.Printf("now: %+v\n", time.Now().UnixMicro()/int64(time.Millisecond))

		fmt.Printf("time: %d\n", time.Since(timee).Microseconds())
	*/
	// 秒级时间戳，eg: 1683633713
	fmt.Printf("time: %+v\n", time.Now().Unix())

}
