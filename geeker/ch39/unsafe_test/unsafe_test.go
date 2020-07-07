package unsafe_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

func TestAtomicGetSet(t *testing.T) {
	var shared_p unsafe.Pointer
	writeDataFunc := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shared_p, unsafe.Pointer(&data))
	}
	readDataFunc := func() {
		data := atomic.LoadPointer(&shared_p)
		//t.Log(*(*[]int)data)
		fmt.Println(data, *(*[]int)(data))
	}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFunc()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFunc()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

/*panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x8 pc=0x50c564]

goroutine 27 [running]:
learngo/geeker/ch39/unsafe_test_test.TestNotAtomicGetSet.func2()
        C:/Users/17102083/go/src/learngo/geeker/ch39/unsafe_test/unsafe_test.go:60 +0x34
learngo/geeker/ch39/unsafe_test_test.TestNotAtomicGetSet.func4(0xc0000324c0, 0xc0000121c0)
        C:/Users/17102083/go/src/learngo/geeker/ch39/unsafe_test/unsafe_test.go:75 +0x37
created by learngo/geeker/ch39/unsafe_test_test.TestNotAtomicGetSet
        C:/Users/17102083/go/src/learngo/geeker/ch39/unsafe_test/unsafe_test.go:73 +0x15d */
// func TestNotAtomicGetSet(t *testing.T) {
// 	var shared_p unsafe.Pointer
// 	writeDataFunc := func() {
// 		data := []int{}
// 		for i := 0; i < 100; i++ {
// 			data = append(data, i)
// 		}
// 		//atomic.StorePointer(&shared_p, unsafe.Pointer(&data))
// 		shared_p = unsafe.Pointer(&data)
// 	}
// 	readDataFunc := func() {
// 		//data := atomic.LoadPointer(&shared_p)
// 		data := shared_p
// 		//t.Log(*(*[]int)data)
// 		fmt.Println(data, *(*[]int)(data))
// 	}
// 	var wg sync.WaitGroup
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func() {
// 			for i := 0; i < 10; i++ {
// 				writeDataFunc()
// 				time.Sleep(time.Microsecond * 100)
// 			}
// 			wg.Done()
// 		}()
// 		wg.Add(1)
// 		go func() {
// 			for i := 0; i < 10; i++ {
// 				readDataFunc()
// 				time.Sleep(time.Microsecond * 100)
// 			}
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }
