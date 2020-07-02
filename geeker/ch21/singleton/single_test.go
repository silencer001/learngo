package once_test

import (
	"fmt"
	"sync"
	"testing"
)

type Singleton struct {
}

var singleInstance *Singleton

var once sync.Once

func GetSingleTonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Ojb")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			obj := GetSingleTonObj()
			fmt.Printf("obj address:%p\n", obj)
			wg.Done()
		}(&wg)
	}
	wg.Wait()
}
