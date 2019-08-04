package main

import (
	"fmt"
	"time"
)

func fixedShooting(msgChannel chan string) {
	for {
		msgChannel <- "fixed shooting"
		time.Sleep(time.Second)
		// fmt.Println("continue fixed shooting ...")
	}
}

func count(msgChannel chan string) {
	for {
		msg := <-msgChannel
		fmt.Println(msg)
	}
}

func threePointShooting(msgChannel chan string) {
	for {
		msgChannel <- "three point shooting"
	}
}

func main3() {
	var c chan string
	c = make(chan string)
	go fixedShooting(c)
	go threePointShooting(c)
	go count(c)
	var input string
	fmt.Scan(&input)
}
