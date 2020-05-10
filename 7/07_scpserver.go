package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func receiveFile(conn net.Conn, filename string, filesize int) {
	fp, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create ", err)
		return
	}
	receivesize := 0
	buf := make([]byte, 1024*16)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("file receive over.")
				break
			} else {
				fmt.Println("conn.Read ", err)
				os.Remove(filename)
				return
			}
		}
		receivesize += n
		fp.Write(buf[:n])
	}
	if receivesize != filesize {
		fmt.Println("file checksum error, file size:%d, but received size:%d\n", filesize, receivesize)
		os.Remove(filename)
		return
	}
}

func main() {
	//更改进程运行目录
	_, err := os.Stat("E:\\download")
	if err != nil {
		if !os.IsExist(err) {
			fmt.Println("path not exist, to create...\n")
			err := os.Mkdir("E:\\download", 0700)
			if err != nil {
				fmt.Println("os.Mkdir ", err)
				return
			}
		} else {
			fmt.Println("os.Stat ", err)
			return
		}
	}

	err = os.Chdir("E:\\download")
	if err != nil {
		//os.Mkdir("~/download", 0700)
		fmt.Println("os.Chdir ", err)
		return
	}
	//创建服务器
	ln, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println("net.Listen ", err)
		return
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("ln.Accept ", err)
		return
	}
	defer conn.Close()

	//获得文件名和文件大小
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read ", err)
		return
	}
	if buf[n-1] != '\n' && buf[n-2] != '\r' {
		fmt.Println("bad format:", string(buf[:n]))
		return
	}
	info := strings.Split(string(buf[:n-2]), ":")
	filename := info[0]
	filesize, _ := strconv.Atoi(info[1])
	fmt.Printf("receive file name: %s, file size: %d\n", filename, filesize)
	conn.Write([]byte("OK"))

	//接收文件
	receiveFile(conn, filename, filesize)

}
