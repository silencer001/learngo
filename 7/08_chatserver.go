package main

import (
	"fmt"
	"net"
)

type Client struct {
	C    chan string
	Name string
	//Addr string
	Addr net.Addr
}

var onlineMap map[string]Client
var message = make(chan string) //上线通知频道

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}
func formatmsg(cli Client, msg string) string {
	return "[" + cli.Addr.String() + "]" + cli.Name + ":" + msg
}
func HandleConn(conn net.Conn) {
	defer conn.Close()
	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr()
	//创建结构体
	cli := Client{make(chan string), cliAddr.String(), cliAddr}
	onlineMap[cliAddr.String()] = cli
	message <- formatmsg(cli, "login")
	go WriteMsgToClient(cli, conn)
	cli.C <- formatmsg(cli, "I'm here")
	//read from conn
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err", err)
			message <- formatmsg(cli, "leave")
			return
		}
		message <- formatmsg(cli, string(buf[:n])) //可能会阻塞？
	}
}
func BroadCast() {
	for {
		msg := <-message
		for _, client := range onlineMap {
			client.C <- msg
		}
	}
}
func main() {
	onlineMap = make(map[string]Client)
	ln, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Println("net.Listen err :", err)
		return
	}
	defer ln.Close()
	//新开一个协程，转发消息，只要消息来了，遍历map，给map每个成员chan发送消息
	go BroadCast()
	//主协程，等待用户连接
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
			continue
		}
		go HandleConn(conn) //处理用户连接
	}
}
