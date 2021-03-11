package main

import (
	"fmt"
	"sync"
	"time"
)

//var sum = 0

var (
	sum int
	mutex sync.Mutex
)

func main() {
	for i := 0; i < 100; i++ {
		go add(10)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("和为:", sum)
}

func add(i int) {
	mutex.Lock()
	sum += i
	mutex.Unlock()
}
