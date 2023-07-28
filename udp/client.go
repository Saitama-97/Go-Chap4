package main

import (
	"fmt"
	"net"
)

/**
 * @Time    : 2023.07.28 11:04
 * @File    : client.go
 * @Project : chapter4
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : UDP Client端
 */

func main() {
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 19999,
	})
	if err != nil {
		fmt.Println("dial udp failed,err:", err)
		return
	}
	defer udpConn.Close()
	// 准备要发送的数据
	sendData := []byte("hello world")
	// 发送数据
	n, err := udpConn.Write(sendData)
	if err != nil {
		fmt.Println("write to udp failed,err:", err)
		return
	}
	fmt.Printf("Write to UDP: count:%v data:%v\n", n, string(sendData))
	// 用来接收数据
	var buf [1024]byte
	n, err = udpConn.Read(buf[:])
	if err != nil {
		fmt.Println("read from udp failed,err:", err)
		return
	}
	fmt.Printf("Receive from UDP: count:%v data:%v\n", n, string(buf[:n]))

	//n1, addr1, err1 := udpConn.ReadFromUDP(buf[:])
	//if err1 != nil {
	//	fmt.Println("read from udp failed,err:", err1)
	//	return
	//}
	//fmt.Printf("Receive from UDP: count:%v addr:%v data:%v\n", n1, addr1, string(buf[:n1]))
}
