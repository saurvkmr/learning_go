package main

import "fmt"

func main() {
	var i int = 0
	for i < 10 {
		fmt.Print(i, " ")
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Print(j, " ")
	}

	for j := range 10 {
		fmt.Print(j, " ")
	}
	for {
		fmt.Println("infinite loop")
		break
	}
}
