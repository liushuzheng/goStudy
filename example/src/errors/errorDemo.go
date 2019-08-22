package main

import (
	"errors"
	"fmt"
)

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 faild", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 faild", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}
	// 如果需要使用自定义错误类型返回的错误数据 需要使用类型断言
	// 来获取一个自定义错误类型的实例才行
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}

//约定错误代码是函数的最后一个返回值
//并且类型是error, 这是一个内置接口
func f1(arg int) (int, error) {

	if arg == 42 {
		// errors.New 使用错误信息作为参数，构建一个基本的错误
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// 可以通过error接口的方法Error() string 来定义错误
// 自定义一个错误类型表示上面例子中的参数错误

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d-%s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
