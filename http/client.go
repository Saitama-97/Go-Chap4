package main

import (
	"fmt"
	"io"
	"net/http"
)

/**
 * @Time    : 2023.08.01 13:50
 * @File    : client.go
 * @Project : chapter4
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : http 编程 - 客户端
 */

func main() {
	resp, _ := http.Get("http://127.0.0.1:8000/go/")
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)

	buf := make([]byte, 1024)

	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read failed, err:", err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}
}
