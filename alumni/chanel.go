package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup()
	wg.Add(1)
	c := make(chan int)
	i := 0
	go func() {
		c <- i
		fmt.Println("go()")
		wg.Done()
	}()
	wg.Wait()
}
