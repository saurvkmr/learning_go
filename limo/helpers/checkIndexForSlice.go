package helpers

import (
	"errors"
	"fmt"
)

func checkForSlice(size, part int) {
	parts, err := Split(size, part)
	if err != nil {
		return
	}
	fmt.Println(parts)
	/**
	for i := 0; i < 9; i++ {
		fmt.Println(len(dummy[splitStart:splitTill]))
		splitStart += parts[0]
		splitTill += parts[0]
	}
	fmt.Println(len(dummy[splitStart : splitTill+parts[1]]))
	**/
}

func Split(num, n int) ([2]int, error) {
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
