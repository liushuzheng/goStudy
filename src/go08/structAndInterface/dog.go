package structandinterface

type Dog struct {
	name string
	age  int
}

//NewDog   构造体
func NewDog() Dog {
	var d Dog
	return d
}

func (dog *Dog) Name(name string) {
	dog.name = name
}

func (dog *Dog) Age(age int) {
	dog.age = age
}
