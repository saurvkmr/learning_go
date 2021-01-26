package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["name"] = "saurav"
	m["title"] = "kumar"

	for fname, lname := range m {
		fmt.Println(fname, " ", lname)
	}

	var presidents = map[string]string{
		"Baiden": "2021 - ",
		"Trump":  "2017 - 2020",
		"Obama":  "2009 - 2017",
	}

	for president, tenure := range presidents {
		fmt.Println(president, " ", tenure)
	}

}
