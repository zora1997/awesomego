package main

import (
	"fmt"
	"sync"
)

func main() {
	uids := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		tmp := uids
		fmt.Printf("tmp: %+v, uids: %+v\n", tmp, uids)
	}()
	for i := range uids {
		wg.Add(1)
		temp := uids[i]
		go func() {
			defer wg.Done()
			fmt.Printf("uid: %d, temp: %d\n", uids[i], temp)
		}()
	}
	wg.Wait()
}
