package main

import (
	"fmt"
	"time"
)

func work(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second * 1)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)
	for w := 1; w <= 3; w++ {
		go work(w, jobs, result)
	}
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	for j := 1; j <= 9; j++ {
		<-result
	}

}
