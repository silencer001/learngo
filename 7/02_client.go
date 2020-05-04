package main

import "net"

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	conn.Write([]byte("are u ok?"))
}
