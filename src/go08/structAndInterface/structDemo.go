//Package structandinterface package如果和所在的文件夹不同 引用的时候需要 加上前缀package名  然后是路径
package structandinterface

//Rect 定义一个矩形
type Rect struct {
	Width, Length float64
}

//Area 计算矩形的面积 方法只能定义在结构体的文件中 定义在结构体外会报错
func (react *Rect) Area() float64 {
	// 结构体遇到指针的时候 不需要使用*去访问结构体的成员
	return react.Length * react.Width
}
