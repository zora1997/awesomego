package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//设置五秒超时，可使用WithDeadline(在某个时间点超时)或者WithTimeout(在多少时间后超时)
	//ctx , cancel := context.WithDeadline(context.Background(),time.Now().Add(time.Second * 5))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	go doTimeOut(ctx)

	time.Sleep(time.Second * 10)
	return
}

//任务处理函数
func doTimeOut(ctx context.Context) {
	for {
		//判断超时
		if deadline, ok := ctx.Deadline(); ok {
			if time.Now().After(deadline) {
				fmt.Println(ctx.Err().Error(), "timeout...")
				return
			}
		}
		select {
		//未超时 任务完成
		case <-ctx.Done():
			fmt.Println("done")
			return
		//继续完成任务
		default:
			fmt.Println("work")
		}
		time.Sleep(time.Second * 1)
	}
}
