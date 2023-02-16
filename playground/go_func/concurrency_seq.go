/*
package main

import (
	"fmt"
	"sync"
)

func main() {
	// 用来阻塞 bbb
	a := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ccc()
	}()

	wg.Add(1)
	go func() {
		// aaa 执行完，close channel a..
		defer close(a)
		defer wg.Done()
		aaa()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bbb(a)
	}()

	wg.Wait()
}

func aaa() {
	fmt.Printf("do aaa\n")
}

func bbb(a chan struct{}) {
	<-a
	fmt.Printf("do bbb\n")
}

func ccc() {
	fmt.Printf("do ccc\n")
}
*/

package main

import (
	"fmt"
	"time"
)

// A首先被a阻塞，A()结束后关闭b，使b可读
func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

// B首先被a阻塞，B()结束后关闭b，使b可读
func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	close(b)
}

// C首先被a阻塞
func C1(a chan struct{}) {
	<-a
	fmt.Println("C1()!")
}

// C首先被a阻塞
func C2(a chan struct{}) {
	<-a
	fmt.Println("C2()!")
}

// C首先被a阻塞
func C3(a chan struct{}) {
	<-a
	fmt.Println("C3()!")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go C1(z)
	go A(x, y) // -> 第 99 行
	go C2(z)
	go B(y, z)
	go C3(z)

	// 关闭x，让x可读
	close(x) // -> 第 97 行
	time.Sleep(3 * time.Second)
}
