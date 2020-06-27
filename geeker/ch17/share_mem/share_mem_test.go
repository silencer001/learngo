package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	var mut sync.Mutex
	for i := 0; i < 5000; i++ {
		go func() {
			mut.Lock()
			defer func() { mut.Unlock() }()
			counter++
		}()
	}
	time.Sleep(time.Second * 1)
	t.Logf("counter = %d\n", counter)
}

func TestCounterThreadSafeWaitGroup(t *testing.T) {
	counter := 0
	var mut sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1) //必须在wait thread中先执行，在goroutine中执行，可能wait先执行，子协程add后执行
		go func() {
			defer wg.Done()
			mut.Lock()
			defer func() { mut.Unlock() }()
			counter++
		}()
	}
	//time.Sleep(time.Second * 1)
	wg.Wait()
	t.Logf("counter = %d\n", counter)
}
