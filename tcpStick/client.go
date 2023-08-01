package main

import (
	"fmt"
	"net"
)

/**
 * @Time    : 2023.08.01 09:17
 * @File    : client.go
 * @Project : chapter4
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : TCP 粘包现象 - 客户端
 */

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12000")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msgStr := "Hello World"
		conn.Write([]byte(msgStr))
	}
}
