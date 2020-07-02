package first_response

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func runtask() string {
	randi := rand.Intn(1000)
	time.Sleep(time.Millisecond * time.Duration(randi))
	ret := "rand millisecond " + strconv.Itoa(randi)
	return ret
}
func FirstResponse() string {
	rand.Seed(time.Now().UnixNano())
	numOfRunner := 10
	ch := make(chan string)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runtask()
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log(FirstResponse())
}
