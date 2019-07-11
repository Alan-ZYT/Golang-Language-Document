package main

/*

	虽然使用Go开发的web应用大多已经处于前后端分离的模式，这里还是需要提及一下静态文件的管理：

	在本案例中，使用http.FileServer(http.Dir("D:\\GoCode"))创建了一个文件服务器，当访问地址以/static/开头时，
	路由器mux将移除该字符串，并在public目录中查找被请求的文件，
	如下代码所示：

*/

import (
	"fmt"
	"net/http"
)

func main() {
	// 创建一个Go默认的路由
	mux := http.NewServeMux()
	// 静态文件管理：注意生成二进制文件时，public的路径，推荐使用绝对路径
	files := http.FileServer(http.Dir("D:\\GoCode"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
