package main

import "fmt"

func main() {
	var x [3]int
	var y = [3]int{}
	z := [4]int{1: 10} // sparse arrray
	fmt.Println(x, y, z)
}
