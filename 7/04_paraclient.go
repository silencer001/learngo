package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Dial :", err)
		return
	}
	defer conn.Close()
	ch := make(chan struct{}, 0)
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					fmt.Println("remote close connetion")
				} else {
					fmt.Println("conn.Read: ", err)
				}
				ch <- struct{}{}
				return
			}
			fmt.Println("read frome server: ", string(buf[:n]))
		}
	}()
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				fmt.Println("os.Stdin.Read :", err)
				return
			}
			conn.Write(buf[:n])
		}
	}()
	<-ch
}
