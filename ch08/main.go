package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("飞雪无情")
		ch <- "goroutine 完成"
	}()

	fmt.Println("我是 main goroutine")

	v := <-ch
	fmt.Println("接收到的chan中的值为：", v)

	cacheCh := make(chan int, 5)
	cacheCh <- 2
	cacheCh <- 3
	fmt.Println("cacheCh容量为:", cap(cacheCh), ",元素个数为：", len(cacheCh))

	firstCh := make(chan string)
	secondCh := make(chan string)
	threeCh := make(chan string)

	go func() {
		firstCh <- downloadFile("firstCh")
	}()

	go func() {
		secondCh <- downloadFile("secondCh")
	}()

	go func() {
		threeCh <- downloadFile("threeCh")
	}()

	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-threeCh:
		fmt.Println(filePath)
	}
}

func downloadFile(chanName string) string {
	//模拟下载文件,可以自己随机time.Sleep点时间试试
	time.Sleep(time.Second)
	return chanName + ":filePath"
}
