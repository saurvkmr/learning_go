package main

import "fmt"

const name = "saurav"

// name = "saurav" a non constant variable cannont be declared outside

func main() {
	myPrint(name)
	// name = "Kumar" // gives error as it supposed to so little bit similar to final in java or val of kotlin
	const salary int = 300 / 10
	fmt.Println(salary)
	// myPrint(salary)

}

func myPrint(anything string) {
	fmt.Println(anything)
}
