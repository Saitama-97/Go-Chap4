package main

import (
	"bufio"
	"fmt"
	"net"
)

/**
 * @Time    : 2023/7/26 16:33
 * @File    : server1.go
 * @Project : chapter4
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : TCP Server端
			  1.监听端口
			  2.接收客户端请求建立连接
			  3.创建goroutine处理连接
*/

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		// 用于存放读取到的数据
		var buf [128]byte
		// 读取数据到buf中，返回读取到的字节数
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		// 将读取到的数据转换为字符串
		recStr := string(buf[:n])
		fmt.Println("收到client端发来的数据:", recStr)
		// 将收到的数据发回给客户端
		conn.Write([]byte(recStr))
	}
}

func main() {
	// 1.监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}
	// 2.接收客户端请求，并创建goroutine处理连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}
		go process(conn)
	}
}
