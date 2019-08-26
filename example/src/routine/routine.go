package main

import (
	"fmt"
	"time"
)

func main() {
	f("direct")
	go f("goroutine")
	var input string
	fmt.Scanln(&input)
}

func f(from string) {

	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Millisecond * 100)
	}

}
