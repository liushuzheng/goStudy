package main

import (
	"fmt"
	"reflect"
)

func main() {
	// sliceDemo1()
	// sliceDemo2()
	// sliceDemo3()
	// sliceDemo4()
	sliceCopy()
}

// make初始化切片
func sliceDemo1() {
	// make 创建切片
	//只设置长度时，容量和长度一致
	var x = make([]float64, 5)
	fmt.Println("Capcity", cap(x), "Length", len(x))
	//长度 5 容量 10
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

//从数组取值赋给切片
func sliceDemo2() {
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var s1 = arr1[2:3]
	var s2 = arr1[:3]
	var s3 = arr1[2:]
	var s4 = arr1[:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}

//给切片添加元素
func sliceDemo3() {
	var arr1 = make([]int, 5, 10)
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	fmt.Println(arr1)
	arr1 = append(arr1, 5, 6, 7, 8, 9, 10)
	fmt.Println("Capacity:", cap(arr1), "Length:", len(arr1))
	fmt.Println(arr1)
}

//从数组取值赋给切片
func sliceDemo4() {
	var arr1 = []int{1, 2, 3, 4, 5}
	fmt.Println("arr1 类型是", reflect.TypeOf(arr1))
	var s1 = arr1[2:3]
	var s2 = arr1[:3]
	var s3 = arr1[2:]
	var s4 = arr1[:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}

func sliceCopy() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := make([]int, 5, 10)
	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)

}
