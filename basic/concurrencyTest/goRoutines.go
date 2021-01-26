package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go justSleep(&wg)
	}
	fmt.Println(time.Since(start))
	wg.Wait()
}

func justSleep(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("sleeping")
	time.Sleep(100 * time.Millisecond)
}
