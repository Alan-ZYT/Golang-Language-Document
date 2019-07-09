package main

/*

代码走读:

	我们在main()创建了一个net.Listener的变量，他是一个服务器的基本函数：用来监听和接收来自客户端的请求
	（来自localhost即IP地址为127.0.0.1端口为8889基于TCP协议）。这个Listen()函数可以返回一个error类型的错误变量。
	用一个无限for循环的listener.Accept()来等待客户端的请求。客户端的请求将产生一个net.Conn类型的连接变量。
	然后一个独立的协程使用这个连接执行HandlerConnect(),开始使用一个4096字节的缓冲data来读取客户端发送来的数据并且把它们打印到服务器的终端，
	n获取客户端发送的数据字节数；当客户端发送的所有数据都被读取完成时，协程就结束了。这段程序会为每一个客户端连接创建一个独立的协程。

必须先运行服务器代码，再运行客户端代码。

*/
import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {

	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	defer listener.Close()

	// 循环监听并接收来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			continue
		}
		// 创建go程,调用 HandlerConnect 函数 处理客户端数据
		go HandlerConnect(conn)
	}
}

func HandlerConnect(conn net.Conn) {
	// 在读取结束时,关闭conn
	defer conn.Close()

	fmt.Println("[", conn.RemoteAddr().String(), "]", "客户端连接成功...")

	buf := make([]byte, 4096)
	// 循环读客户端发送的数据
	for {
		n, err := conn.Read(buf)

		if string(buf[:n-1]) == "exit" {
			fmt.Println("服务器检测到", conn.RemoteAddr().String(), "客户端退出...")
			return
		}

		if n == 0 {
			fmt.Println("服务器检测到", conn.RemoteAddr().String(), "客户端关闭,务器关闭本端!!!")
			return
		}

		if err != nil && err != io.EOF {
			fmt.Println("conn.Read error:", err)
			return
		}
		//处理数据 (小写转大写为例)
		retStr := strings.ToUpper(string(buf[:n]))
		//打印显示客户端数据
		fmt.Printf("Received Client %s data: %v", conn.RemoteAddr().String(), string(buf[:n]))

		// 将转化后的 大写字符串写回
		conn.Write([]byte(retStr))
	}
}
