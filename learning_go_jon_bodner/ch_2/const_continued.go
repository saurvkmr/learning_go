package main

import "fmt"

const x = 10
const y int = 10

func main() {
	var a int = x
	var b float64 = x
	var c byte = x

	//var d float32 = y
	// cannot use y (constant 10 of type int) as float32 value in variable declarationcompilerIncompatibleAssign

	fmt.Println(x, a, b, c)
	fmt.Println(y)
}
