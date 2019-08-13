package main

import (
	"fmt"
	"time"
)

func fixedShooting2(msgChannel chan string) {
	for {
		msgChannel <- "fixed shooting"
		time.Sleep(time.Second * 1)
		// fmt.Println("continue fixed shooting ...")
	}
}

func threePointShooting2(msgChannel chan string) {
	for {
		msgChannel <- "three point shooting"
		time.Sleep(time.Second * 1)
	}
}

func mains() {

	cFixed := make(chan string)
	c3Point := make(chan string)
	go fixedShooting2(cFixed)
	go threePointShooting2(c3Point)
	go func() {
		for {
			select {
			case msg1 := <-cFixed:
				fmt.Println(msg1)
			case msg2 := <-c3Point:
				fmt.Println(msg2)
			}
		}
	}()

	var input string
	fmt.Scan(&input)
}
