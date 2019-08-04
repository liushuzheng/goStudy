package closure

import "fmt"

/* 闭包代码测试 */
func intSeq() func() int {

	i := 0
	return func() int {
		i++
		return i
	}
}

// Closure 简单闭包
func Closure() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

}
