package main

import "fmt"

func main() {
	// initDateMap()
	// noDataMap()
	// crudMap()
	complexMap()
}

//初始化map 有数据
func initDateMap() {
	var x = map[string]string{
		"A": "Apple",
		"B": "Banana"}
	for key, value := range x {
		fmt.Println("key:", key, "value:", value)

	}
}

//初始化map 无数据 需要用make函数
func noDataMap() {

	var x map[string]string
	// x["A"] = "Apple" 未进行初始化 x 为nil 报错
	fmt.Println(x == nil)
	x = make(map[string]string)
	x["A"] = "Apple"
	fmt.Println(x == nil)
	for key, value := range x {
		fmt.Println("key:", key, "value:", value)

	}
}

func crudMap() {
	//create
	x := make(map[string]int)
	x["A"] = 5
	x["B"] = 10
	// fmt.Println(x["B"])
	//判断是否有值
	/* if value, ok := x["C"]; ok {
		fmt.Println(value)
	} */
	fmt.Println("Before Delete")
	fmt.Println("Length:", len(x))
	//reader
	fmt.Println(x)
	//delete
	delete(x, "A")

	fmt.Println("After Delete")
	fmt.Println("Length:", len(x))
	fmt.Println(x)
	//update
	x["B"] = 15
	fmt.Println(x)

}

//复杂map
func complexMap() {
	//1 先声明变量，后赋值
	// var student = make(map[string]map[string]int)
	// student["1"] = map[string]int{"小明": 15, "小刚": 16}
	//2声明变量时就赋值
	var student = map[string]map[string]int{"1": {"小明": 15, "小刚": 16}}
	for stuNo, stuInfo := range student {
		fmt.Println("Student:", stuNo)
		for name, age := range stuInfo {
			fmt.Println("name", name, "age", age)
		}
	}

}
