package main

import "fmt"

func main() {
	square := make(chan int)

	go calculateSqr(square)

	for val := range square {
		fmt.Println("val - ", val)
	}

	/**
	for {
		val, ok := <-square
		if ok == false {
			fmt.Println("second loop", val, ok)
			break
		} else {
			fmt.Println(val)
		}
	}
	*/
}

func calculateSqr(square chan int) {
	for i := 0; i < 10; i++ {
		square <- i * i
	}
	close(square)
}
