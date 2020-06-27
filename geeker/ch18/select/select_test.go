
package select_csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 1000)
	return "done"
}

func otherService() {
	fmt.Println("otherService done")
	time.Sleep(time.Millisecond * 500)
}

func async_service() chan string {
	ch := make(chan string)
	go func() {
		fmt.Println("debug: service start!")
		ch <- service()
		fmt.Println("debug: service end!")
	}()
	return ch
}

func TestSyncService(t *testing.T) {
	fmt.Println(service())
	otherService()
}

func TestAsyncService(t *testing.T) {
	ch := async_service()
	otherService()
	fmt.Println(<-ch)
	fmt.Println("TestAsync end!")
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-async_service():
		fmt.Println(ret)
	case <- time.After(time.Second):
		fmt.Println("time out")
	}
}