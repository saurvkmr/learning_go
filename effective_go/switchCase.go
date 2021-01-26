package main

import (
	"flag"
	"fmt"
)

func main() {
	var day = flag.String("days", "Sunday", "which day is it? ")
	flag.Parse()
	fmt.Printf(weekday(*day))

}

func weekday(day string) string {
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return "Weekday"
	case "Saturday", "Sunday":
		return "Weekend"
	}
	return "Learn days of a week"
}
