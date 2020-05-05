package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func writeToServer(conn net.Conn, buf []byte) error {
	total := 0
	to := len(buf)
	for {
		n, err := conn.Write(buf)
		if err != nil {
			//str := "writeToServer" + err
			return fmt.Errorf("writeToserver : %w", err)
		}
		total += n
		if total == to {
			return nil
		}
		buf = buf[n:]
	}
}
func main() {
	//检查输入文件，获取filename、filesize
	list := os.Args
	if len(list) < 3 {
		fmt.Println(`useage: xxx 		file
        ip:port 	server ip&port`)
		return
	}
	info, err := os.Stat(list[1])
	if err != nil {
		fmt.Println("os.Stat :", err)
		return
	}
	fp, err := os.Open(list[1])
	if err != nil {
		fmt.Println("os.Open :", err)
		return
	}
	defer fp.Close()
	//连接server
	tcp, err := net.ResolveTCPAddr("tcp", list[2])
	if err != nil {
		fmt.Println("get remote addr error:", err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcp)
	//conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Dial :", err)
		return
	}
	defer conn.Close()
	fmt.Println("debug lipadrr:", conn.LocalAddr().String())
	fmt.Println("debug ripaddr:", conn.RemoteAddr().String())

	buf := make([]byte, 1024*16)
	//向server端发送filename和size，等待回复OK
	str := fmt.Sprintf("%s:%d\r\n", info.Name(), info.Size())
	wn, _ := conn.Write([]byte(str))
	if wn != len(str) {
		fmt.Println("write failed, string len:%d, write len:%d", len(str), wn)
		return
	}
	rn, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read :", err)
		return
	}
	if string(buf[:rn]) != "OK" {
		fmt.Println("bad format:", string(buf[:rn]))
		return
	}

	//向server端发送文件内容
	for {
		n, err := fp.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("send success!")
				return
			} else {
				fmt.Println("fp.Read :", err)
				return
			}
		}
		err = writeToServer(conn, buf[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
