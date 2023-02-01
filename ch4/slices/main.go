package main

import (
	"fmt"
)

func main() {
	months := []string{
		"",
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	fmt.Printf("Months value: %v\n", months)

	q2 := months[4:7]
	fmt.Printf("Quater value: %v\n", q2)

	summer := months[6:9]
	fmt.Printf("Summer value: %v\n", summer)
}
