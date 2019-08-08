package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func first() {
	fmt.Println("first func run")
}

func second() {
	fmt.Println("second func run")
}

//int 方法在每个go文件运行时会先运行
// func init() {
// 	fmt.Println("初始化方法")
// }

// defer释放文件资源
func fileDemo() {
	fName := "D:\\abc\\aa\\11.txt"
	f, err := os.Open(fName)
	defer f.Close()
	if err != nil {
		os.Exit(1)
	}
	// var breader *(bufio.Reader)
	breader := bufio.NewReader(f)
	for {
		line, ok := breader.ReadString('\n')
		if ok != nil {
			break
		}
		fmt.Print(strings.Trim(line, "\r\n"))

		// fmt.Print(line)
	}
}

//错误代码 不会走panic 下的代码
func panicExample() {
	fmt.Println("I am walking and singing ...")
	panic("It starts to rain cats and dogs")
	msg := recover()
	fmt.Println(msg)

}

func panics() {
	var a, b int = 2, 0
	a = (a / b)
	fmt.Println(a)
}

func panicCorrect() {
	defer func() {
		msg := recover()
		fmt.Println(msg)

	}()
	fmt.Println("I am walking and singing ...")
	// panic("It starts to rain cats and dogs")
	panics()

}

func main() {
	// defer 在函数运行结束时候运行一段代码或者调用一个清理函数
	// defer second()
	// first()
	// fileDemo()
	// panicExample()
	panicCorrect()
}
