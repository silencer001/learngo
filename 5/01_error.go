package main

import (
	"errors"
	"fmt"
)

func main() {
	//err是*errors.errorString类型，但这个类型是非导出的，怎么理解
	//err := fmt.Errorf("this is a error：%s", "test test test")
	var err error = fmt.Errorf("this is a error：%s", "test test test")
	fmt.Printf("err type is %T, err string:%s\n", err, err)

	err2 := errors.New("this is normal err2")
	fmt.Printf("err2 type:%T, err2 string:%s", err2, err2)
}
