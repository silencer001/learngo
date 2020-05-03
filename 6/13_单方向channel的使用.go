package main

import "fmt"

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)
	for {
	}
}
