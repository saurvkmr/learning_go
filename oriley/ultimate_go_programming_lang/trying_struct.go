
package main

import "fmt"

type emp struct {
	name string
	age  int
	dept string
}

func main() {
	alice := emp{"alice", 30, "sales"}
	fmt.Println(alice)
  // anonymous struct
	june := struct {
		name string
		age  int
		dept string
	}{"june", 10, "dept"}

	fmt.Println(june)
}
