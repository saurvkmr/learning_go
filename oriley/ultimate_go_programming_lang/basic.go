package main

import "fmt"

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
  // declearing anything with var sets zero value the variable for the data type
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var a bool \t %T [%v]\n", d, d)

	var e1 example

	fmt.Printf("var e1 example \t %T [%v]\n", e1, e1)

}
