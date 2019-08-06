/*
panic:
	停止当前函数执行
	一直向上返回，执行每一层的defer
	如果没有遇见recover，程序退出

recover：
	仅在defer调用中使用
	获取oanic的值
	如果无法处理，可以重新panic

panic error使用场景：
	意料之中的：使用error，例如：文件打不开
	意料之外的：使用panic。例如：数组越界

错误处理示例：
	defer+panic+recover
	Type Assertion
	函数式编程的灵活应用
	
http服务器的性能分析
	import _ "net/http/pprof"
	访问 /debug/pprof example：http://127.0.0.1:8888/debug/pprof/
	使用 go tool pprof 分析性能
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"资源管理与出错处理/filelistingServer/filelisting"
)

// 统一出错处理
type appHandler func(writer http.ResponseWriter, request *http.Request) error

//函数式编程；参数和返回值都是一个函数
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//panic
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)

		if err != nil {
			log.Printf("Error occurred handling erequest: %s", err.Error())

			// user Error
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			// system Error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	//Error() string
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.Handlerfilelist))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("http.ListenAndServe Error:", err)
		return
	}
}
