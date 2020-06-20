package main

import "fmt"

func main() {
	var name = "Saurav "
	lastName := "kumar"

	var age int = 30
	fmt.Println("My name is %s and my age is %d", name+lastName, age)

	name += lastName
	fmt.Println(name)
}
