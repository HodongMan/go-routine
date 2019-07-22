package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	helloWaitGroup := func() {
		defer wg.Done()
		fmt.Println("hello")
	}

	wg.Add(1)
	go helloWaitGroup()
	wg.Wait()
}
