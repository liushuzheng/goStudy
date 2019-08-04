package main

import "fmt"

/**
闭包代码测试
*/

func intSeq() func() int {

	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
