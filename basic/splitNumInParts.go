package main

import "errors"
import "fmt"

func main() {
	var dummy [12345]int
	parts, err := split(12345, 10)
	if err != nil {
		return
	}
	splitStart := 0
	splitTill := parts[0]
	for i := 0; i < 9; i++ {
		fmt.Println(len(dummy[splitStart:splitTill]))
		splitStart += parts[0]
		splitTill += parts[0]
	}
	fmt.Println(len(dummy[splitStart:splitTill+parts[1]]))
}

func split(num, n int) ([2]int, error) {
	var parts [2]int
	if n == 0 || num < n {
		return parts, errors.New("math: incompatible numbers")
	}
	rem := num % n
	x := num - rem
	div := x / n
	parts[0] = div
	parts[1] = rem
	return parts, nil
}