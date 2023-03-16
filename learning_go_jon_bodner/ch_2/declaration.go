package main

import "fmt"

func main() {
	var x int = 10
	var y = 10
	var z int // z is assigned zero value of int
	var a, b int = 10, 20
	var d, c = 10, 20
	var name, age = "Albus", 0

	var (
		my_name, my_age     = "Albus", 0
		wife_name, wife_age = "magonigal", 0
		house               string
	)

	//elders_wand_owner string = "Ablus Dumbledore" without var you can provide type
	elders_wand_owner := "Malfoy"
	elders_wand_owner, stupid_owner := "Harry", "you know who"

	fmt.Println(x, y, z, a, b, c, d, name, age)
	fmt.Println(my_name, my_age, wife_name, wife_age, house)
	fmt.Println(elders_wand_owner, stupid_owner)

}
