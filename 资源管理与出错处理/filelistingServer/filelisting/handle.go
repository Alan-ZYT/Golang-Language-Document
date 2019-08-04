package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (u userError) Error() string {
	return u.Message()
}

func (u userError) Message() string {
	return string(u)
}

func Handlerfilelist(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(
		request.URL.Path, prefix) != 0 {
		return userError(fmt.Sprintf("PATH %s must start with %s", request.URL.Path, prefix))
	}
	path := request.URL.Path[len(prefix):] // list/fib.txt
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer, err.Error(), http.StatusNotFound)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//fmt.Println("Error:", err)
		return err
	}
	writer.Write(all)
	return nil
}
