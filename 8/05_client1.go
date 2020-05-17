package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000/go")
	if err != nil {
		fmt.Println("http.Get:", resp)
		return
	}
	defer resp.Body.Close()
	//fmt.Println(resp.Body)

	fmt.Println("Status = ", resp.Status)
	fmt.Println("StatusCode = ", resp.StatusCode)
	fmt.Println("Header = ", resp.Header)

	buf := make([]byte, 4*1024)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 { //如果判断err！= nil，Body最后一次的数据无法打印
			fmt.Println("read err = ", err)
			break
		}
		fmt.Println("body:", string(buf[:n]))
	}
}
