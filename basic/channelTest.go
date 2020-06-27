package main

import "fmt"

func main() {
	var chname chan string
	fmt.Println(chname)

	anotherChan := make(chan string)
	fmt.Println(anotherChan)

	anotherChan <- "Saurav"
	fmt.Println(<-anotherChan)
	var name = <-anotherChan
	fmt.Println(name)
}
