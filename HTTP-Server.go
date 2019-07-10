package main

/*

	http协议：
	应用层协议。广泛应用于 万维网的网络数据通信

	http: 超文本传输协议 (明文)

	https: 超文本传输协议 (加密)  -- SSL协议  TLS 协议 (网景公司)

URL:
	统一资源定位符: 在网络环境中唯一确定一个 数据资源.

	语法:  协议  服务器名(IP) 路径名 文件名

http 请求协议格式：【重点】

	1. 请求行：	请求方法 “空格” URL“空格”协议版本 \r\n		GET /itcast.txt HTTP/1.1

	2. 请求头：	key：value		Host：127.0.0.1:8000

	3. 空行：		\r\n		代表 http请求头 结束。

	4. 请求包包体：	受请求方法影响。GET方法没有包体。	POST 含有包体。

http 应答协议格式：【重点】

	1. 应答行：	协议版本	  状态码	状态描述

	2. 应答头：	key：value

	3. 空行：		\r\n		代表 http应答头 结束。

	4. 应答包包体。	成功：浏览器请求的数据		200 ok

			失败：失败信息。			404 Not Found


	HTTP是一个比TCP更高级的协议，它描述了客户端浏览器如何与网页服务器进行通信。Go语言有自己的net/http包

	以下示例 代码引入了http包并启动了网页服务器，和net.Listen("tcp", "127.0.0.1:8889")函数的tcp服务器是类似的，
	使用http.ListenAndServe("127.0.0.1:8080", nil)函数，
	如果成功会返回空，否则会返回一个错误（可以指定127.0.0.1为其它地址，8080是指定的端口号）


	创建 HTTP 服务器：

	1. 注册回调函数	http.HandleFunc("服务器提供的URL"， 回调函数名)

	2. 实现回调函数

		func myHandler(w http.ResponseWriter, r *http.Request) {

			w：写回给客户端的数据

			r：从客户端读到的数据

	3. 创建监听 接收客户端（浏览器）请求

		http.ListenAndServe("IP+port", nil)

	func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

		参1：服务器提供 url

		参2：回调函数名

回调函数：func(w http.ResponseWriter, r *http.Request)

type ResponseWriter interface {

    Header() Header

    WriteHeader(int)

    Write([]byte) (int, error)  【重点】
}

type Request struct {

    Method string

    URL *url.URL

    Proto      string // "HTTP/1.0"

    Header Header

    Body io.ReadCloser
    .....
}

*/

import (
	"log"
	"net/http"
)

//实现回调函数
func handler(w http.ResponseWriter, r *http.Request) {
	//w: 写回给客户端(浏览器)数据
	//r: 从浏览器读到的数据

	//io.WriteString(w, "Hello,HttpServer...\n")
	//fmt.Fprintf(w, "Hello,"+r.URL.Path[1:])
	w.Write([]byte("Hello," + r.URL.Path[1:]))
}

//注册回调函数
func main() {
	//注册回调函数 http.HandleFunc("服务器提供的URL"， 回调函数名)
	//第一个参数是请求的路径，第二个参数是处理这个路径请求的函数的引用
	http.HandleFunc("/", handler)

	//创建监听 接收客户端（浏览器）请求
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err.Error())
	}
}
