package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		newerr := fmt.Errorf("HttpGet %w", err)
		return "", newerr
	}
	defer resp.Body.Close()

	//读取网页body内容
	buf := make([]byte, 4*1024)
	var result string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			if err == io.EOF {
				break
			} else {
				newerr := fmt.Errorf("resp.Body %w", err)
				return "", newerr
			}
		}
		result += string(buf[:n])
	}
	return result, nil
}
func DoWork(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 页面\n", start, end)
	//	var sw sync.WaitGroup

	sw := sync.WaitGroup{}
	//sw.Add(end - start + 1)
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=dota&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		fmt.Println("debug url = ", url)
		//并发爬取
		sw.Add(1)
		go func(i int, url string) {
			defer sw.Done()
			res, err := HttpGet(url)
			if err != nil {
				fmt.Println("HttpGet ", err)
				return
			}
			fileName := strconv.Itoa(i) + ".html"
			f, err := os.Create(fileName)
			if err != nil {
				fmt.Println("os.Create err ", err)
				return
			}
			defer f.Close()

			f.Write([]byte(res))
		}(i, url) //必须传入i和url，否则传入的i固定为最后一个值
	}
	sw.Wait()
}
func main() {
	var start, end int
	fmt.Printf("请输入起始页>")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页")
	fmt.Scan(&end)

	DoWork(start, end)
}
