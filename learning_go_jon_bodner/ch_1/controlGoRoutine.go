package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	elem := make(chan int, 20)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			elem <- i
			time.Sleep(500 * time.Millisecond)
			fmt.Print(<-elem)
		}(i)
	}

	wg.Wait()
}
