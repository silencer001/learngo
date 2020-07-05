package syncpool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	sp := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Object.")
			return 100
		},
	}

	v := sp.Get().(int)
	fmt.Println(v)
	sp.Put(3)
	runtime.GC()            //此时会清除sync.pool中缓存的对象
	v1, _ := sp.Get().(int) //get 100
	fmt.Println(v1)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	sp := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new object.")
			return 10
		},
	}
	sp.Put(100)
	sp.Put(100)
	sp.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(sp.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
