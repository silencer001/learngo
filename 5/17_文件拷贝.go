package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	src := os.Args[1]
	dst := os.Args[2]
	if src == dst {
		fmt.Println("dest file same to source file!")
		return
	}
	sf, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer df.Close()
	var total int
	buf := make([]byte, 4*1024)
	for {
		sn, err := sf.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("read end n =", sn)
				break
			}
			panic(err)
		}
		dn, err := df.Write(buf[:sn])
		if err != nil {
			panic(err)
		}
		if dn != sn {
			fmt.Printf("wite to dest err :read %d, but write %d\n", sn, dn)
			return
		}
		total += dn
	}
	fmt.Println("copy success! total byte :", total)
}