package main

import "fmt"

func main() {
	// var num int
	// fmt.Scan(&num)

	// switch num {
	// case 1:
	// 	fmt.Println("按下的是1楼")
	// 	break //go语言保留了break关键字，跳出switch语句，默认已经包含了break
	// case 2:
	// 	fmt.Println("按下的是2楼")
	// 	break
	// case 3:
	// 	fmt.Println("按下的是3楼")
	// default:
	// 	fmt.Println("未识别的楼层")
	// }

	var score int
	fmt.Println("请输入分数")
	fmt.Scan(&score)
	//switch可以加一个初始化语句，或者为空
	switch {
	case score > 90:
		fmt.Println("优秀")
		//fallthrough
	case score > 80:
		fmt.Println("良好")
		//fallthrough
	case score > 60:
		fmt.Println("及格")
	default:
		fmt.Println("差")
	}
}
