package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main2() {

	go listElem(10, "go_a")
	go listElem(10, "go_b")
	var input string
	fmt.Scanln(&input)

}

func listElem(n int, tag string) {
	for i := 0; i < n; i++ {
		fmt.Println(tag, i)
		tick := time.Duration(rand.Intn(100))
		time.Sleep(time.Millisecond * tick)
	}

}
