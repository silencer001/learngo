package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func numToSlice(num int) []int {
	s := make([]int, 4)
	num *= 10
	for i := 0; num > 10; i++ {
		s[i] = num / int(math.Pow10(4-i))
		num = num % int(math.Pow10(4-i))
	}
	fmt.Println("s = ", s)
	return s
}

func scanNum() int {
	var num int
	for {
		fmt.Println("请输入四位数：")
		fmt.Scan(&num)
		if num > 10000 || num < 1000 {
			fmt.Println("数值不合法，请重新输入！")
		} else {
			break
		}
	}
	return num
}

func createRandnum() int {
	rand.Seed(time.Now().UnixNano())
	for {
		num:=rand.Intn(10000)
		if num >= 1000 {
			return num
		}
	}
}

func compareNum(src []int) {
	for {
		n:=0
		num := scanNum()
		s:= numToSlice(num)
		for i:=0;i<len(s);i++ {
			if s[i] > src[i] {
				fmt.Printf("第%d位大了!\n",i)
			}else if s[i] < src[i] {
				fmt.Printf("第%d位小了!\n",i)
			}else {
				fmt.Printf("第%d位猜对了!\n",i)
				n++
			}
		}
		if n == len(src) {
			break
		}
	}
}
func main() {
	randnum := createRandnum()
	fmt.Println("randnum:",randnum)
	randslice := numToSlice((randnum))	
	compareNum(randslice)
}
