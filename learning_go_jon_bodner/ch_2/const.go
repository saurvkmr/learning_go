package main

import "fmt"

const (
	name  = "who am i"
	genre = "hacker"
)

const platform string = "netflix"

// const status := "watched" constants cannot be inistalized via :=

func main() {
	const runningTime = 1.20
	const time = "eventing"

	fmt.Println(time, runningTime)

	//time = "morning" cannot be reasigned

	fmt.Println(name, genre, platform)
}
