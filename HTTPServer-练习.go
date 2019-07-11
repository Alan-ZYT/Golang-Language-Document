package main

/*

练习一:

	编写一个网页服务器监听端口9999，有如下处理函数：

当请求http://localhost:9999/hello/Name时，响应：hello Name(Name需是一个合法的姓，比如Chris或者Madeleine)

当请求http://localhost:9999/shouthello/Name时，响应：hello NAME


	Example:

	http://127.0.0.1:8889/hello/alan

	Hello alan!

	http://127.0.0.1:8889/shouthello/alan

	Hello ALAN!

练习二:

	创建一个空结构hello并使它实现http.Handler。

Example:

	http://127.0.0.1:5000/

	Hello!

*/

import (
	"fmt"
	"net/http"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/hello/"):] //get everything after the /hello/ part of the URL
	fmt.Fprintf(w, "Hello %s!", remPartOfURL)
}

func shouthelloHandler(w http.ResponseWriter, r *http.Request) {
	remPartOfURL := r.URL.Path[len("/shouthello/"):] //get everything after the /shouthello/ part of the URL
	fmt.Fprintf(w, "Hello %s!", strings.ToUpper(remPartOfURL))
}

func main001() {
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/shouthello/", shouthelloHandler)
	http.ListenAndServe("127.0.0.1:9999", nil)
}

type Hello struct{}

func main() {
	var h Hello
	http.ListenAndServe("127.0.0.1:5000", h)
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}
