package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 10; i++ {
		var num = rand.Intn(1000)
		fmt.Println(num)
	}
}
