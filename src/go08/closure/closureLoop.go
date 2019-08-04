package go08

import "fmt"

//ClosuLoop 循环闭包
func ClosureLoop() {
	var fs []func() int
	for i := 0; i < 3; i++ {
		fs = append(fs, func() int {
			// fmt.Println("i is ", i)
			return i
		})
	}
	for _, f := range fs {
		fmt.Printf("%p = %v\n", f, f())
	}
	//输出的都是3 闭包是在调用的时候才会调用函数中的变量
}
