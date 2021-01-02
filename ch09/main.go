package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum   int
	mutex sync.RWMutex
)

func main() {
	run()
	doOnce()
	race()
	syncMap()
}

func run() {
	var wg sync.WaitGroup
	//因为要监控110个协程，所以设置计数器为110
	wg.Add(110)
	for i := 0; i < 100; i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			add(10)
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			fmt.Println("和为:", readSum())
		}()
	}
	//一直等待，只要计数器值为0
	wg.Wait()
}

func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

//10个人赛跑,1个裁判发号施令
func race() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号已经就位")
			cond.L.Lock()
			cond.Wait() //等待发令枪响
			fmt.Println(num, "号开始跑……")
			cond.L.Unlock()
		}(i)
	}
	//等待所有goroutine都进入wait状态
	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast() //发令枪响
	}()
	wg.Wait()
}

func syncMap() {
	syncMap := sync.Map{}
	syncMap.Store(1, 1)
	syncMap.Store(1, 2)
	fmt.Println(syncMap.Load(1))
}

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
}

func readSum() int {
	mutex.RLock()
	defer mutex.RUnlock()
	b := sum
	return b
}
