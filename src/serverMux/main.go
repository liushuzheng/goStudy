package main

import (
	"fmt"
	"net/http"
)

func main() {

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", handler)
	// 创建路由
	http.ListenAndServe(":8080", serverMux)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建的多路复用器处理请求", r.URL.Path)
}
