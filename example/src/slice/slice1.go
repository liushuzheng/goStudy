package main

import "fmt"

func main() {

	s := make([]string, 3)
	fmt.Println("emp", s)
	//可以用和数组的一样的方法来设置元素值或获取值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	//用len获取切片长度
	fmt.Println("len:", len(s))
	//追加会生成新的
	s = append(s, "d", "e")
	s = append(s, "f")
	fmt.Println("append:", s)

}
