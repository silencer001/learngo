package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	//打开文件
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("this is data test file.\n")
	for i := 0; i < 10; i++ {

		fmt.Fprintf(file, "i = %d\n", i)
	}
}
func ReadFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	buf := make([]byte, 1024*2)
	n, err := f.Read(buf)
	//panic: EOF
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Println("buf = ", string(buf[:n]))
}

func ReadFileLine(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//新建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(f)
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("read file end")
				break
			} else {
				panic(err)
			}
		}
		fmt.Printf("buf = #%s#\n", string(buf))
	}
}
func main() {
	path := os.Args[1]
	WriteFile(path)
	ReadFileLine(path)
}
