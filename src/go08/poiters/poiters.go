package main

import "fmt"

func easyPointer() {
	var x int
	var xptr *int
	x = 10
	//将x 的地址赋给指针xptr
	xptr = &x
	fmt.Println(x)
	fmt.Println(xptr)
	// 获取xptr 地址
	fmt.Println(&xptr)
	fmt.Println(&xptr)
	fmt.Println(*xptr)
}

// 支持直接换值
func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	// easyPointer()
	var x, y = 10, 20
	fmt.Println(x)
	fmt.Println(y)
	swap(&x, &y)
	fmt.Println(x)
	fmt.Println(y)
}
