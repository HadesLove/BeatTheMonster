package main

import (
	"fmt"
	"time"
)

func main()  {
	/*go fmt.Println("golang")
	fmt.Println("我是 main goroutine")
	time.Sleep(time.Second)*/

	/*ch := make(chan string)

	go func() {
		fmt.Println("golang")
		ch <- "goroutine 完成"
	}()

	fmt.Println("我是 main goroutine")

	v := <-ch
	fmt.Println("接收到的chan中的值为:", v)

	cacheCh := make(chan int, 5)

	cacheCh <- 2
	cacheCh <- 3
	fmt.Println("cacheCh容量为:", cap(cacheCh), ", 元素个数为:", len(cacheCh))
	close(cacheCh)*/

	//onlySend := make(chan<- int)
	//onlyReceive := make(<-chan int)

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
	time.Sleep(time.Second)
	return chanName + ":filePath"
}
