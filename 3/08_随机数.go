package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(int64(os.Getpid()) & time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
}
