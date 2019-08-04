package main

import (
	"fmt"
	"runtime"
	"time"
)

// getSumDivisible 未知
func getSumDivisible(num int, divider int, resultChan chan int) {
	sum := 0
	for value := 0; value < num; value++ {
		if value%divider == 0 {
			sum += value
		}
	}
	resultChan <- sum
}

func main1() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	LIMIT := 1000
	//用于被15除结果
	job1 := make(chan int, 1)
	//用于3 5 除结果
	job2 := make(chan int, 2)
	start := time.Now()
	go getSumDivisible(LIMIT, 15, job1)
	go getSumDivisible(LIMIT, 3, job2)
	go getSumDivisible(LIMIT, 5, job2)
	// sum15 :=
	sum15, sum3, sum5 := <-job1, <-job2, <-job2
	sum := sum3 + sum5 - sum15
	end := time.Now()
	fmt.Println(sum)
	fmt.Println(end.Sub(start))

}
