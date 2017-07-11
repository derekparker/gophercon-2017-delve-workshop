package main

import "fmt"

func main() {
	ch := make(chan int)
	go uhhhhWAT(doSomething(ch))
	ch <- 5

}

func uhhhhWAT(i int) {
	fmt.Println(i)
}

func doSomething(ch chan int) int {
	return <-ch
}
