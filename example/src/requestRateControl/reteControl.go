package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)
	for rep := range requests {
		<-limiter
		fmt.Println("request", rep, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

	// myChan := make(chan int, 5)
	// myChan <- 10
	// close(myChan)
	// <-myChan
	// _, ok := <-myChan
	// if !ok {
	// 	fmt.Println("获取了数据")
	// }
	// _, oks := <-myChan
	// if !oks {
	// 	fmt.Println("获取了数据")
	// }

}
