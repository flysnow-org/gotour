package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)

	ctx,stop:=context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		watchDog(ctx,"【监控狗1】")
	}()

	go func() {
		defer wg.Done()
		watchDog(ctx,"【监控狗2】")
	}()

	go func() {
		defer wg.Done()
		watchDog(ctx,"【监控狗3】")
	}()

	valCtx:=context.WithValue(ctx,"userId",2)
	go func() {
		defer wg.Done()
		getUser(valCtx)
	}()

	time.Sleep(5 * time.Second) //先让监控狗监控5秒
	stop() //发停止指令
	wg.Wait()
}

func watchDog(ctx context.Context,name string) {
	//开启for select循环，一直后台监控
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name,"停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name,"正在监控……")
		}
		time.Sleep(1 * time.Second)
	}
}

func getUser(ctx context.Context){
	for  {
		select {
		case <-ctx.Done():
			fmt.Println("【获取用户】","协程退出")
			return
		default:
			userId:=ctx.Value("userId")
			fmt.Println("【获取用户】","用户ID为：",userId)
			time.Sleep(1 * time.Second)
		}
	}
}
