package main

import (
	"fmt"
	"time"
)

func saySomething(something string) {
	fmt.Println("I'm about to do a thing")
	fmt.Println(something)
	fmt.Println("I did it")
}

func loopForNoReasonThisManyTimes(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("looped %d times\n", i)
	}
	callMe()
}

func callMe() {
	myslc := make([]int, 1, 10)
	myslc[0] = 1
	fmt.Println(myslc)
}

func launchABunchOfGoroutines() {
	for i := 0; i < 10; i++ {
		go doSomethingUseless(i)
	}
}

func doSomethingUseless(i int) {
	time.Sleep(time.Second)
	fmt.Println("uh..")
	time.Sleep(time.Second)
	fmt.Println("..hello..")
	time.Sleep(time.Second)
	fmt.Println("..there")
}

func main() {
	myarg := "something"
	saySomething(myarg)
	n := 10
	loopForNoReasonThisManyTimes(n)
	launchABunchOfGoroutines()
	time.Sleep(time.Second)
	fmt.Println("done")
}
