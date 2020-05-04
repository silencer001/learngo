package main

import (
	"fmt"
	"io"
	"net"
)

/*
func createConn(ln net.Listener) (net.Conn, error) {
	for()
}*/
func readFunc(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()
	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("read buf: ", string(buf[:n]))
		}
	}

}
func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	conch := make(chan net.Conn)
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println(err)
				continue
			}
			conch <- conn
		}
	}()
	for c := range conch {
		go readFunc(c)
	}
	for {
	}
}
func main01() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1024)
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		for {
			n, err := conn.Read(buf)
			if err == io.EOF {
				conn.Close()
				break
			} else if err != nil {
				panic(err)
			} else {
				fmt.Println("reaf buf: ", buf[:n])
			}
		}
	}
}
