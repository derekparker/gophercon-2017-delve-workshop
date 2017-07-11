package main

import (
	"fmt"
	"time"
)

const workCount = 10

func main() {
	ch := produceWork(workCount)
	go dosomethingIrrelevant()
	consumeWork(ch)
}

func dosomethingIrrelevant() {
	for {
		time.Sleep(time.Second)
		fmt.Println("still haven't solved it?")
	}
}

func produceWork(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
	}()
	return ch
}

func consumeWork(ch chan int) {
	for {
		n := <-ch
		fmt.Println(n)
	}
}
