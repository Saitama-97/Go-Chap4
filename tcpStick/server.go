package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

/**
 * @Time    : 2023.08.01 09:03
 * @File    : server.go
 * @Project : chapter4
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : TCP 粘包现象 - 服务端
 */

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
		}
		if err == io.EOF {
			break
		}
		recString := string(buf[:n])
		fmt.Println("收到数据:", recString)
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:12000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
