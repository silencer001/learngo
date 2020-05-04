package main

import (
	"bytes"
	"fmt"
	"net"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()
	//获取客户端的网络地址
	addr := conn.RemoteAddr().String()
	fmt.Println("addr connect sucessful")

	//读取用户数据
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		//bytes.Trim(s []byte, cutset string)

		str := bytes.TrimFunc(buf[:n], func(r rune) bool {
			if r == '\r' || r == '\n' {
				return true
			} else {
				return false
			}
		})
		//windows下会自动嵌入\r\n
		if "exit" == string(str) {
			return
		}
		fmt.Println(addr, " read: ", string(str))
		//把数据转换为大写
		//conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
		conn.Write(bytes.ToUpper(str))
	}
}
func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			continue
		}
		go HandleConn(conn)
	}
}
