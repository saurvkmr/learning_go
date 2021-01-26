package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
}

func (p *Person) setAge(age int) {
	p.age = age
}

func (e Employee) isMajor() bool {
	return e.age > 18
}

func main() {
	p := Person{name: "Lucifer"}
	p.setAge(25)
	fmt.Println(p)

	emp := Employee{p}
	fmt.Println(emp.isMajor())
}
