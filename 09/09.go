package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum = 0
	//mutex sync.Mutex
	mutex sync.RWMutex
)

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	sum += i
	//mutex.Unlock()
}

func readSum() int {
	mutex.RLock()
	defer mutex.RUnlock()
	b := sum
	return b
}

func main()  {
	race()
}

func run()  {
	var wg sync.WaitGroup
	wg.Add(110)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			add(10)
		}()
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("和为:", readSum())
		}()
	}
	wg.Wait()
}

func doOnce()  {
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

func race() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号已经就位")
			cond.L.Lock()
			cond.Wait()
			fmt.Println(num, "号开始跑......")
			cond.L.Unlock()
		}(i)
	}

	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast()
	}()
	wg.Wait()
}