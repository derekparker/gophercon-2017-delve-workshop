package main

import "fmt"

type foo struct {
	a string
}

func main() {
	imFreakingOut()
}

func imFreakingOut() {
	var a *foo
	fmt.Println(a.a)
}
