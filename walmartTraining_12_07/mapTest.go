package main

import (
	"fmt"
)

func main() {
	simpleMap()
	nestedMap()
}

func simpleMap() {
	countryCapital := map[string]string{
		"AFGHANISTAN": "KABUL",
		"BELGIUM":     "BRUSSELS",
		"INDIA":       "NEW DELHI",
	}

	for country, capital := range countryCapital {
		fmt.Println("Country :", country, ", Capital: ", capital)
	}

	_, ok := countryCapital["Bangalesh"]
	if !ok {
		fmt.Println("No such country")
	}
}

func nestedMap() {
	countryStateCapital := map[string]map[string]string{
		"India": {
			"Bihar":       "Patna",
			"West Bengal": "Kolkata"},
		"US": {
			"Alabama":  "Montgomery",
			"Arkansas": "Little Rock"},
	}
	for country, stateCapital := range countryStateCapital {
		fmt.Println("Country :", country)
		for state, capital := range stateCapital {
			fmt.Println("State: ", state, ", Capital: ", capital)
		}
	}
}
