package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)

	// 创建路由
	http.ListenAndServe(":8080", nil)

}

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World", r.URL.Path)
}
