package main

import (
	"fmt"
	"net/http"
)

/**
 * @Time    : 2023.08.01 10:52
 * @File    : server.go
 * @Project : chapter4
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : http 编程 - 服务端
 */

func main() {
	// 单独写回调函数
	http.HandleFunc("/go", myHandler)
}

func myHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.RemoteAddr, "连接成功")
	fmt.Println("method:", request.Method)
	fmt.Println("url:", request.URL)
	fmt.Println("header:", request.Header)
	fmt.Println("body:", request.Body)
	writer.Write([]byte("hello"))
}
