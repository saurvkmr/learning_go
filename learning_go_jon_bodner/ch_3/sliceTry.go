package main

import "fmt"

func main() {
	var x []int
	var y = []int{10, 20}
	fmt.Println(x, y, len(x), len(y))
	y = append(y, 50)
	fmt.Println(x, y, len(x), len(y))
	y = append(y, 60, 70)
	fmt.Println(x, y, len(x), len(y))
	z := []int{101, 102}
	y = append(y, z...)
	fmt.Println(y, len(y), cap(y))

	a := make([]int, 5)
	fmt.Println(a, len(a), cap(a))

	b := make([]int, 5, 10)
	fmt.Println(a, len(b), cap(b))
	b = append(b, 5, 6, 7)
	fmt.Println(b, len(b), cap(b))
	c := make([]int, 0, 10)
	c = append(c, 5, 6, 7)
	fmt.Println(c, len(c), cap(c))
}
