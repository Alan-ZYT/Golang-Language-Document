package main

/*

	输出结果:

	成功输出:
	HTTP/1.1 200 OK  //应答行
	//应答头
	Date: Wed, 10 Jul 2019 16:57:27 GMT
	Content-Length: 14
	Content-Type: text/plain; charset=utf-8
					//空行
	Hello,userinfo  //应答包体


	失败输出:

	HTTP/1.1 404 Not Found
	Content-Type: text/plain; charset=utf-8
	X-Content-Type-Options: nosniff
	Date: Wed, 10 Jul 2019 17:00:23 GMT
	Content-Length: 19

	404 page not found

*/
import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("net.Dial error", err.Error())
		return
	}
	defer conn.Close()

	//伪装浏览器发送HTTP请求协议
	httpRequest := "GET /userinfo HTTP/1.1\r\nHost:127.0.0.1:8080\r\n\r\n"
	conn.Write([]byte(httpRequest))

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal("conn.Read error", err.Error())
		return
	}
	fmt.Printf("\n%s\n", string(buf[:n]))
}
