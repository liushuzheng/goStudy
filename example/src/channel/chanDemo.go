package main

import "fmt"

func main1() {

	messages := make(chan string)

	go func(message chan string) {
		for i := 0; i < 10; i++ {

			messages <- fmt.Sprintf("ping+%d", i)
			fmt.Println(i)
		}

	}(messages)

	go func(message chan string) {
		for {
			msg := <-message
			fmt.Println(msg)
		}

	}(messages)

	var input string
	fmt.Scanln(&input)

}
