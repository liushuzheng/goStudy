package structs

import (
	"fmt"
	"sqlweb/util"
)

// USer   用户结构体
type USer struct {
	ID       int
	Username string
	Password string
	Email    string
}

//AddUser 预编译添加用户
func (user *USer) AddUser() error {

	sqlStr := "insert into users (Username,password,email) values (?,?,?)"
	// 预编译
	stmt, err := util.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常:", err)
		return err
	}
	_, err1 := stmt.Exec(user.Username, user.Password, user.Email)
	if err1 != nil {
		fmt.Println("执行出现异常", err1)
		return err1
	}
	// fmt.Printf("结果是%v", result)
	return nil

}
