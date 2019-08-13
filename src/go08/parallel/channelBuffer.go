package main

import (
	"fmt"
	"strconv"
	"time"
)

func shooting(msgChan chan string) {
	var group = 1
	for {
		if group <= 10 {
			for i := 1; i <= 10; i++ {
				msgChan <- strconv.Itoa(group) + ":" + strconv.Itoa(i)
			}
		}
		group++
		time.Sleep(time.Second * 2)
	}
}

func counts(msgChan chan string) {
	for {
		fmt.Println(<-msgChan)
	}
}

func main() {
	var c = make(chan string)
	go shooting(c)
	go counts(c)
	var input string
	fmt.Scan(&input)

}
