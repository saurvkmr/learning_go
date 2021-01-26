package main

import (
	"fmt"
)

type Speaker interface {
	speak() string
}

type Dog struct {
	sound string
}

func (d Dog) speak() string {
	return d.sound
}

type Cat struct {
	sound string
}

func (c Cat) speak() string {
	return c.sound
}

type Human struct {
	sound string
}

func (h Human) speak() string {
	return h.sound
}

type Alien struct {
	sound string
}

func (a Alien) speak() string {
	return a.sound
}

func main() {
	d := Dog{"bow"}
	c := Cat{"meow"}
	h := Human{"hi"}
	a := Alien{"aiiee"}

	species := [...]Speaker{d, c, h, a}
	for _, s := range species {
		fmt.Println(s)
	}

}
