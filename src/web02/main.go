package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	myHandler := MyHandler{}
	// http.Handle("/myHander", &myHandler)
	// http.ListenAndServe("", nil)

	// server 结构体

	server := http.Server{

		Addr:        ":8181",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second,
	}
	fmt.Println("自定义的server")
	server.ListenAndServe()
}

// MyHandler 自定义handler
type MyHandler struct{}

// serverHTTP 实现的接口
func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "自定义的处理器结构体", r.URL.Path)
}
