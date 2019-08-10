package main

import (
	"fmt"
	"go08/structandinterface"
)

func main() {
	//未使用key 警告
	// var rect = structandinterface.Rect{5, 6}
	var rect structandinterface.Rect
	rect.Width = 10.2
	rect.Length = 2
	fmt.Println((&rect).Area())
	// 结构体遇到指针的时候 不需要使用传入地址
	fmt.Println(rect.Area())
}
