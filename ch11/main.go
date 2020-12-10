package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	result := make(chan string)
	go func() {
		//模拟网络访问
		time.Sleep(8 * time.Second)
		result <- "服务端结果"
	}()

	select {
	case v := <-result:
		fmt.Println(v)
	case <-time.After(5 * time.Second):
		fmt.Println("网络访问超时了")
	}

	coms := buy(100) //采购100套配件
	//三班人同时组装100部手机
	phones1 := build(coms)
	phones2 := build(coms)
	phones3 := build(coms)
	//汇聚三个channel成一个
	phones := merge(phones1, phones2, phones3)
	packs := pack(phones) //打包它们以便售卖

	//输出测试，看看效果
	for p := range packs {
		fmt.Println(p)
	}

	vegetablesCh := washVegetables() //洗菜
	waterCh := boilWater()           //烧水

	fmt.Println("已经安排洗菜和烧水了，我先眯一会")
	time.Sleep(2 * time.Second)
	fmt.Println("要做火锅了，看看菜和水好了吗")
	vegetables := <-vegetablesCh
	water := <-waterCh
	fmt.Println("准备好了,可以做火锅了:", vegetables, water)

}

//工序1采购
func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}

//工序2组装
func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装(" + c + ")"
		}
	}()
	return out
}

//工序3打包
func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包(" + c + ")"
		}
	}()
	return out
}

//扇入函数（组件）,把多个chanel中的数据发送到一个channel中
func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	//把一个channel中的数据发送到out中
	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}

	wg.Add(len(ins))

	//扇入，需要启动多个goroutine用于处于多个channel中的数据
	for _, cs := range ins {
		go p(cs)
	}

	//等待所有输入的数据ins处理完，再关闭输出out
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

//洗菜
func washVegetables() <-chan string {
	vegetables := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		vegetables <- "洗好的菜"
	}()
	return vegetables
}

//烧水
func boilWater() <-chan string {
	water := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		water <- "烧开的水"
	}()
	return water
}
