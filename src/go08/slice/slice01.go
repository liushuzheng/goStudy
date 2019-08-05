package main

import "fmt"

func main() {
	// make 创建切片
	var x = make([]float64, 5)
	fmt.Println("Capcity", cap(x), "Length", len(x))
	var y = make([]float64, 5, 10)
	fmt.Println("Capcity", cap(y), "Length", len(y))
	for i := 0; i < len(x); i++ {
		//数据类型转换
		x[i] = float64(i)
	}
	fmt.Println(x)
	for i := 0; i < len(y); i++ {
		//数据类型转换
		y[i] = float64(i)
	}
	fmt.Println(y)

}
