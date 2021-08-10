package main

import "fmt"

func main() {
	array()
	slice()
}

func array() {
	var arr [3]string
	arr[0] = "India"
	arr[1] = "America"
	arr[2] = "Thailand"

	arr_2 := [3]string{"India", "America", "Thailand"}
	arr_3 := [...]string{"India", "America", "Thailand"}

	fmt.Println(arr)
	fmt.Println(arr_2)
	fmt.Println(arr_3)
}

func slice() {
	countries := []string{"India", "America", "Thailand"}
	countries = append(countries, "Germany", "Poland", "Greece")

	europe := make([]string, 5)
	europe = append(europe, "Germany", "Poland", "Sweden", "Norway")

	fmt.Println(countries)
	fmt.Println(europe)

	for i, country := range europe {
		fmt.Println(i, country)
	}
}
