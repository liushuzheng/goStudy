package structs

import "testing"

func TestAddUser(t *testing.T) {
	user := USer{
		Username: "张三",
		Password: "123456",
		Email:    "lsz@ml.com",
	}
	user.AddUser()
}
