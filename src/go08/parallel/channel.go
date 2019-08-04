package main

import (
	"fmt"
	"time"
)

func fixedShooting(msgChannel chan string) {
	for {
		msgChannel <- "fixed shooting"
		fmt.Println("continue fixed shooting ...")
	}
}

func count(msgChannel chan string) {
	for {
		msg := <-msgChannel
		fmt.Println(msg)
		time.Sleep(time.Second)
	}
}

func main() {
	var c chan string
	c = make(chan string)
	go fixedShooting(c)
	go count(c)
	var input string
	fmt.Scan(&input)
}
