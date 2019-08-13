package main

import (
	"fmt"
	"time"
)

func fixedShooting3(msgChannel chan string) {
	var times = 3
	var t = 1
	for {
		if t <= times {
			msgChannel <- "fixed shooting"
		}
		t++
		time.Sleep(time.Second * 1)
		// fmt.Println("continue fixed shooting ...")
	}

}

func threePointShooting3(msgChannel chan string) {
	var times = 5
	var t = 1
	for {
		if t <= times {
			msgChannel <- "three point shooting"
		}
		t++
		time.Sleep(time.Second * 1)
	}
}

func main4() {

	cFixed := make(chan string)
	c3Point := make(chan string)
	go fixedShooting3(cFixed)
	go threePointShooting3(c3Point)
	go func() {
		for {
			select {
			case msg1 := <-cFixed:
				fmt.Println(msg1)
			case msg2 := <-c3Point:
				fmt.Println(msg2)
			case <-time.After(time.Second * 5):
				fmt.Println("timeout check again ...")
			}
		}
	}()

	var input string
	fmt.Scan(&input)
}
