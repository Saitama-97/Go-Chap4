package main

import (
	"fmt"
	"net"
)

/**
 * @Time    : 2023.07.28 10:42
 * @File    : server.go
 * @Project : chapter4
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : UDP Server端，UDP是无连接的，不需要建立连接，直接发送数据
			  UDP属于无连接、不可靠、面向报文的传输协议，但是实时性好，适合用于音视频传输
*/

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 19999,
	})
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}
	defer listener.Close()
	for {
		// 存储读取到的数据
		var buf [1024]byte
		// 接收客户端发送的数据
		n, addr, err := listener.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("read from udp failed,err:", err)
			continue
		}
		// 打印接收的数据
		fmt.Printf("Receive from UDP: count:%v addr:%v data:%v\n", n, addr, string(buf[:n]))
		// 将接收到的数据返回给客户端
		dataToWrite := []byte(string(buf[:n]) + "123")
		n, arr := listener.WriteToUDP(dataToWrite, addr)
		if arr != nil {
			fmt.Println("write to udp failed,err:", arr)
			continue
		}
		// 打印发送的数据
		fmt.Printf("Write to UDP: count:%v addr:%v data:%v\n", n, addr, string(dataToWrite))
	}
}
