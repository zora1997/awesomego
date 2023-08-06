package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	attempts := int32(3)
	backoff := time.Duration(300) * time.Millisecond
	err := Retry(attempts, backoff, func() error {
		if attempts >= 2 {
			fmt.Printf("failed, attempts: %d\n", attempts)
			attempts--
			return errors.New("attempts >= 2")
		}
		fmt.Printf("success, attempts: %d\n", attempts)
		return nil
	})
	if err != nil {
		fmt.Printf("all failed")
	}

	fmt.Printf("success, attempts final: %d\n", attempts)
}

// Retry 失败简易重试方法 提供 2^n 次方退避等待
// attempts 尝试重试次数
// backoff 重试间隔时间 会随着重试次数以 2^n 次方递增
// fn 重试函数 可将需要重试的函数放入其中 并返回 err 返回值
func Retry(attempts int32, backoff time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(Stop); ok {
			return s.error
		}

		if attempts--; attempts > 0 {
			time.Sleep(backoff)
			return Retry(attempts, 2*backoff, fn)
		}
		fmt.Printf("attempts #%d after %s.", attempts, backoff)
		return err
	}
	return nil
}

type Stop struct {
	error
}
