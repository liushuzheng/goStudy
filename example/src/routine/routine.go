package main

import (
	"fmt"
	"time"
)

func main() {
	// f("direct")
	// go f("goroutine")
	// var input string
	// fmt.Scanln(&input)
	// intChan := make(chan int)
	// go test2(intChan)
	// go test1(intChan)
	// var input string
	// fmt.Scanln(&input)
	cn := make(chan int, 10)
	go func(chans chan int) {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Second * 1)
			chans <- i
		}
		close(cn)
	}(cn)

	for {
		v, ok := <-cn
		if !ok {
			fmt.Println("退出")
			break
		}
		fmt.Println(v)
	}
	// var s string
	// fmt.Scanln(&s)

}

func f(from string) {

	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Millisecond * 100)
	}

}

func test1(intChan chan int) {

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		intChan <- i
		fmt.Println("发送数据了", i)
	}

}

func test2(intChan chan int) {

	for {

		i := <-intChan
		fmt.Println("接收到数据了", i)
	}

}
