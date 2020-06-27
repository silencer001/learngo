package groutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		//		go func(i int) {
		//fmt.Println(i)
		//}(i)
		fmt.Println("main thread i: ", i)
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}
