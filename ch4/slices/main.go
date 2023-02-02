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

	appendToSlice()
}

func appendToSlice() {
	var integers []int

	fmt.Printf("lenght of integers: %v\n", len(integers))
	fmt.Printf("capacity of integers: %v\n", cap(integers))

	for i := 0; i < 101; i++ {
		integers = append(integers, i)
	}

	fmt.Printf("value of integers: %v\n", integers)
	fmt.Printf("lenght of integers: %v\n", len(integers))
	fmt.Printf("capacity of integers: %v\n", cap(integers))
}

