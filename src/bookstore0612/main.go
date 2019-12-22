package main

import (
	"net/http"
	"text/template"
)

func main() {
	// 设置静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))
	http.HandleFunc("/main", IndexHandler)
	http.ListenAndServe(":8080", nil)
}

//	indexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//	解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, "")
}
