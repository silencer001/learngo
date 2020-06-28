package ch_broadcast

import (
	"fmt"
	"sync"
	"testing"
)

func producer(ch chan int, wg *sync.WaitGroup) {
	go func() {
	for i:=0;i<10;i++ {
		ch<-i
	}

	close(ch)
	wg.Done()
	}()
}

func receiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for data := range ch {
			fmt.Println(data)
		}
		wg.Done()
	}()
}

func TestCloseChanel(t *testing.T) {
	ch := make(chan int, 1)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	receiver(ch, wg)
	wg.Add(1)
	producer(ch, wg)
	wg.Wait()
}