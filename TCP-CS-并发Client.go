package main

/*

代码走读:

	客户端通过net.Dial创建了一个和服务器之间的连接

	它通过无限循环中的os.Stdin接收来自键盘的输入直到输入了“exit”。注意使用\r和\n换行符分割字符串（在windows平台下使用\r\n）。
接下来分割后的输入通过connection的Write方法被发送到服务器。

	当然，服务器必须先启动好，如果服务器并未开始监听，客户端是无法成功连接的。

	如果在服务器没有开始监听的情况下运行客户端程序，客户端会停止并打印出以下错误信息：
对tcp 127.0.0.1:8889发起连接时产生错误：由于目标计算机的积极拒绝而无法创建连接。

	打开控制台并转到服务器和客户端可执行程序所在的目录，
Windows系统下输入TCP-CS-并发Server.exe（或者只输入 TCP-CS-并发Server），Linux系统下输入 ./TCP-CS-并发Server

	接下来控制台出现以下信息：Starting the server ...

	在Windows和linux系统中，我们可以通过CTRL/C停止程序。

	然后开启2个或者3个独立的控制台窗口，然后分别输入TCP-CS-并发Client回车启动客户端程序

以下是服务器的输出：
	Starting the server ...
	[ 127.0.0.1:9727 ] 客户端连接成功...
	[ 127.0.0.1:9728 ] 客户端连接成功...
	Received Client 127.0.0.1:9728 data: hello
	Received Client 127.0.0.1:9727 data: world

*/

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// 创建通信socket
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		// 由于目标计算机积极拒绝而无法创建连接
		fmt.Println("net.Dial error", err)
		return
	}
	defer conn.Close()

	strbuf := make([]byte, 4096)

	// 在子go程中，读stdin
	go func() {
		// 循环从Stdin中读键盘输入
		for {
			fmt.Println("What to send to the server? Type exit to quit.")
			n, err := os.Stdin.Read(strbuf) //不能使用fmt.Scan,因为fmt.Scan遇见空格/回车 自动终止提取
			if err != nil {
				fmt.Println("=os.Stdin.Read error", err)
				return
			}
			// 将键盘读到的数据，通过conn写到服务器
			conn.Write(strbuf[:n])
		}
	}()

	buf := make([]byte, 4096)

	// 循环在主go程中，与服务器进行通信
	for {
		n, err := conn.Read(buf)

		if n == 0 {
			fmt.Println("客户端检测到服务器关闭。关闭本端！")
			return
		}

		if err != nil && err != io.EOF {
			fmt.Println("=conn.Read error", err)
			return
		}
		fmt.Println("服务器回发数据:", string(buf[:n]))
	}
}
