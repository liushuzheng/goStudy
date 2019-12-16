package main

import (
	"encoding/json"
	"net/http"
	"sqlweb/structs"
)

func testJSonRes(w http.ResponseWriter, r *http.Request) {

	// 设置相应的类型
	w.Header().Set("content-type", "application-json")
	// 创建用户
	user := structs.USer{
		ID:       1,
		Username: "admin",
		Password: "123",
		Email:    "admin@qq.com",
	}

	marshal, _ := json.Marshal(user)
	w.Write(marshal)

}

func testRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(302)
}

func main() {

	http.HandleFunc("/testJson", testJSonRes)
	http.HandleFunc("/redirect", testRedirect)
	http.ListenAndServe(":8080", nil)

}
