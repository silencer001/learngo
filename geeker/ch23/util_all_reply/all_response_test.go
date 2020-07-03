
package all_response

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runtask(i int) string {
	//	rand.Seed(time.Now().UnixNano())
	//	randi := rand.Intn(1000)
	time.Sleep(time.Millisecond * 10)
	//ret := "rand millisecond " + strconv.Itoa(randi)
	return fmt.Sprintf("return task from %d", i)
}
func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runtask(i)
			ch <- ret
		}(i)
	}
	//return <-ch
	var ret string
	for i:=0;i<numOfRunner;i++ {
		ret += <-ch + "\n"
	}
	return ret
}

func TestFirstResponse(t *testing.T) {
	t.Log("before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second)
	t.Log("after:", runtime.NumGoroutine())
}

// func TestRandInt(t *testing.T) {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 10; i++ {
// 		t.Log("i value: ", rand.Intn(1000))
// 	}
// }
