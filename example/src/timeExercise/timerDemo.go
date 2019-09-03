package main

import (
	"fmt"
	"time"
)

func main() {
	tickerDemo()
}

func timerDemo() {
	// Timer 代表了未来的一个事件 告诉Timer 要等待多久
	// 然后计时器提供了一个通道 这个通道在等待的时间结束后得到通知
	// 这里的timer 将等待2秒
	timer1 := time.NewTimer(time.Second * 2)

	// 这里的'<-timer1.C' 在timer的通道'C' 上面阻塞等待 直到有个值发送给通道 通知通道计时器已经等待完成
	// time.NewTimer方法获取的timer1的结构体定义为
	// type Timer struct {
	// 	C <-chan Time
	// 	r runtimeTimer
	// }
	<-timer1.C
	fmt.Println("Timer 1 expired")
	// 如果仅仅需要等待的话 可以使用'time.Sleep'
	// 而timer的独特之处在于你可以在你可以在timer等待完成之前取消等待
	timer2 := time.NewTimer(time.Second * 1)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

func tickerDemo() {

	// Ticker 使用和 Timer 相似的机制 同样使用一个通道来发送数据
	// 这里我们使用range 函数遍历通道的数据 这些数据每个500毫秒被发送一次
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Ticker at", t)
		}
	}()

	// Ticker 和 Timer 一样可以被停止 一旦Ticker
	// 停止后 通道将不再接收数据 这里我们将在 1500 毫秒之后停止
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopper")
}
