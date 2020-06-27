package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	num := make(chan int)
	go allPrimes(num)

	num <- 1000000
	for val := range num {
		fmt.Println("prime", val)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func allPrimes(num chan int) {
	number := <-num
	for i := 3; i <= number; i++ {
		if isPrime(i) {
			num <- i
		}
	}
	close(num)
}

func isPrime(num int) bool {
	isPrime := true
	sqrt := int(math.Sqrt(float64(num))) + 1
	for i := 2; i <= sqrt; i++ {
		if (num % i) == 0 {
			isPrime = false
			break
		}
	}

	return isPrime
}
