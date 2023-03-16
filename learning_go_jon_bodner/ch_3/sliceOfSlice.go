package main

import (
	"fmt"
)

func main() {
	var x = make([]int, 0, 20) // nil slice
	fmt.Println(x, len(x), cap(x))
	for i := 0; i < cap(x); i++ {
		x = append(x, i)
		//x[i] = i * i
	}
	fmt.Println(x, len(x), cap(x))
}
