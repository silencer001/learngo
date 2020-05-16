package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err :", err)
		return 
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("ln.Accept err: ", err)
		return 
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err :", err)
			return 
		}
		fmt.Printf("#%v#", string(buf[:n]))
	}
}