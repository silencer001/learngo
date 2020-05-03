package main

import "fmt"

func main() {
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	//对于map只有len，没有cap
	fmt.Println("len = ", len(m1))
	//m1 =  map[]
	//len =  0
	//panic: assignment to entry in nil map
	//	m1[1]="judeng"

	//可以通过make创建
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len = ", len(m2))
	//invalid operation: m1 == m2 (map can only be compared to nil)
	//fmt.Println("m1 == m2:", m1==m2)
	m2[1] = "王爽爽"
	//通过make创建，可以指定长度（实际是容量），len为空
	m3 := make(map[int]string, 10)
	fmt.Println("m3 = ", m3)
	fmt.Println("len = ", len(m3))

	m4 := map[int]string{1:"judeng",3:"王爽",2:"jututu"}
	fmt.Println("m4:",m4)

}
