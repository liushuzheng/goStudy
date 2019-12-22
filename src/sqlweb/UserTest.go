package main

import (
	"fmt"
	"sqlweb/structs"
	"sqlweb/util"
)

func main() {
	user := GetUser()
	fmt.Println(user)
}

func GetUser() *structs.USer {
	user := &structs.USer{}
	row := util.DB.QueryRow("select * from users")
	row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	return user
}
