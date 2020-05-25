package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func HttpGet(url string, client *http.Client) (result []byte, err error) {
	fmt.Println("debug:enter HttpGet")
	resp, err := client.Get(url)
	if err != nil {
		fmt.Errorf("HttpGet(%s):%w", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	buf := make([]byte, 4*1024)
	result = []byte{}
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			if err == io.EOF {
				break
			}
			fmt.Errorf("rep.Body.Read : %w", err)
			return nil, err
		}
		result = append(result, buf[:n]...)
	}
	return result, nil
}
func spiderJoy(url string, client *http.Client) (title, content string, err error) {
	result, err := HttpGet(url, client)
	if err != nil {
		return "", "", fmt.Errorf("HttpGet(%s):%w", url, err)
	}
	result1 := strings.ReplaceAll(string(result), "\n", "")
	//fmt.Println(string(result))
	//<h1 class="f18"><a href="https://m.pengfu.com/content/1857807/" title="
	reg := regexp.MustCompile(`<h1 class="f18"><a href="(?:(.*?))">(?:(.*?))</a></h1>`)
	tiles := reg.FindAllStringSubmatch(string(result1), 1)
	for _, data := range tiles {
		//fmt.Println("tile[]:#", data, "#")
		title = data[2]
		break
	}
	//<div class="con-txt">
	reg1 := regexp.MustCompile(`<div class="con-txt">(?:(.*?))</div>`)

	contents := reg1.FindAllStringSubmatch(string(result1), 1)
	//fmt.Println(contents)
	for _, data := range contents {
		content = data[1]
		break
	}
	content = strings.TrimSpace(content)
	return title, content, nil
}
func SpiderPage(i int, client *http.Client, joys chan<- string, wg *sync.WaitGroup) {
	//明确要爬取的url
	//https://m.pengfu.com/index_1.html
	url := "https://m.pengfu.com/index_" + strconv.Itoa(i) + ".html"
	fmt.Printf("current url:%s\n", url)

	//开始爬取页面
	result, err := HttpGet(url, client)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(result))
	//<h1 class="f18"><a href="https://m.pengfu.com/content/1857793/" title="拜祖先">拜祖先</a></h1>
	//取<h1 class="dp-b"><a href="URL"
	reg := regexp.MustCompile(`<h1 class="f18"><a href="(?s:(.*?))"`)
	//reg := regexp.MustCompile(`div class="(?s:(.*?))"`)
	joyUrls := reg.FindAllStringSubmatch(string(result), -1)
	//fmt.Println("joyurls", joyUrls)
	if joyUrls != nil {
		for _, data := range joyUrls {
			title, content, err := spiderJoy(data[1], client)
			if err != nil {
				fmt.Println(err)
				continue
			}
			joys <- title + "\n" + content + "\n"
		}
	}
	wg.Done()
}
func saveToFile(joys <-chan string) {
	fp, err := os.Create("./joy.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	for joy := range joys {
		fp.Write([]byte(joy))
	}
}
func DoWork(start, end int, client *http.Client) {
	fmt.Printf("准备爬取第%d页到%d页的网址\n", start, end)
	wg := sync.WaitGroup{}
	joys := make(chan string)
	for i := start; i <= end; i++ {
		wg.Add(1)
		//定义要给函数，爬主页面
		go SpiderPage(i, client, joys, &wg)
	}
	go func() {
		wg.Wait()
		close(joys)
	}()
	saveToFile(joys)
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页>")
	fmt.Scan(&start)
	fmt.Printf("请输入中止页>")
	fmt.Scan(&end)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	DoWork(start, end, client)
}
