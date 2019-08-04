package main

import "fmt"

/* 闭包代码测试 */
func intSeq() func() int {

	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// nextInt := intSeq()
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	var fs []func() int
	for i := 0; i < 3; i++ {
		fs = append(fs, func() int {
			fmt.Println("i is ", i)
			return i
		})
	}
	for _, f := range fs {
		fmt.Printf("%p = %v\n", f, f())
	}
	//输出的都是3 闭包是在调用的时候才会调用函数中的变量
}
