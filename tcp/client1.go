package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/**
 * @Time    : 2023/7/26 16:33
 * @File    : client1.go
 * @Project : chapter4
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : TCP Client端
			  1.建立与Server端的连接
			  2.发送数据
			  3.接收数据
			  4.关闭连接
*/

func main() {
	// 1.建立与Server端的连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Dial failed,err:", err)
		return
	}
	defer conn.Close()
	// 2.发送数据
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 读取用户输入
		input, _ := inputReader.ReadString('\n')
		// 整理输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}
		// 3.接收数据
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read from server failed,err:", err)
		}
		recvData := string(buf[:n])
		fmt.Println("收到server端发来的数据:", recvData)
	}
}
