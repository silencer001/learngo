package main

import (
	"fmt"

	"errors"
)

func MyPanic(a, b int) (result int, err error) {
	if b == 0 {
		err := errors.New("Error:b is zero!")
		return 0, err
	}
	return a / b, nil
}

func main() {
	result, err := MyPanic(10, 0)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("result = ", result)
	}
}
