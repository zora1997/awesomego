package main

import "sync"

var lock sync.Mutex
var count = 1

// loop count++
func work() {
	for {
		lock.Lock()
		count++
		lock.Unlock()
	}
}

func main() {

}
