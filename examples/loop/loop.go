package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	words := []string{"Well", ",", "that's", "just", "like", "your", "opinion", ",", "man"}
	wg.Add(len(words))
	for _, w := range words {
		go func() {
			fmt.Println(w)
			wg.Done()
		}()
	}
	wg.Wait()
}
