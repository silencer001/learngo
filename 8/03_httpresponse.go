package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial", err)
		return
	}
	defer conn.Close()

	reqheader := "GET / HTTP/1.1\r\nCache-Control: max-age=0\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Language: zh-Hans-CN,zh-Hans;q=0.5\r\nUpgrade-Insecure-Requests: 1\r\nUser-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.18363\r\nAccept-Encoding: gzip, deflate\r\nHost: 127.0.0.1:8000\r\nConnection: Keep-Alive\r\n\r\n"
	//先发请求包，服务器才会回相应包
	conn.Write([]byte(reqheader))
	fmt.Printf("start receive")
	//接受服务器回复的相应包
	buf := make([]byte, 1024*4)
	n, err := conn.Read(buf)
	if n == 0 {
		fmt.Println("read err =", err)
		return
	}
	fmt.Printf("#%v#", string(buf[:n]))
}

/*
PS C:\Users\juden\go\src\learngo\8> go run .\03_httpresponse.go
start receive#HTTP/1.1 200 OK
Date: Sat, 16 May 2020 02:54:39 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

hello world
#
PS C:\Users\juden\go\src\learngo\8> go run .\03_httpresponse.go
start receive#HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Sat, 16 May 2020 02:55:06 GMT
Content-Length: 19

404 page not found
#
*/
