package main

import ("fmt")

func main() {
	sum, diff, product := mathOps(10, 20)
	fmt.Printf("sum = %v, difference = %d, product = %d", sum, diff, product)
}

func mathOps(first, second int) (int, int, int) {
	return first + second, first - second, first * second
}