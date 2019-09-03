package main

import "fmt"

func main() {
	loopChannel()
}

func closeChannel() {

	jobs := make(chan int, 5)
	done := make(chan bool)
	//这里是数据接收端协程 他重复使用 ' j,more := <-jobs'来从通道jobs获取数据
	// 这里的more 在通道关闭且通道中不在有数据可以接收的时候为false 我们通过判断more 来决定
	// 所有的数据师傅已经接收完成 如果所有的数据接收完成 那么向done通道写入 true
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("receive all jobs")
				done <- true
				return
			}
		}

	}()

	//这里向通道写入三个数据 然后关闭通道
	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	// 我们知道done通道在接收数据的时候会阻塞 所以在所有的数据发送接收完成后
	// 写入done的数据将在这里被接收 然后程序结束
	<-done
}

func loopChannel() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	//range 函数遍历每个从通道接收到的数据
	// 因为queue 在发送完两个数据之后就关闭了通道 所以这里我们range函数在接收到两个数据之后就结束了
	// 如果上面的queue 通道不关闭 那么range函数就不会结束 从而在接收第三个数据的时候就阻塞了
	for elem := range queue {

		fmt.Println(elem)
	}

}
