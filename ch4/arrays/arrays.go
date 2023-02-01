package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	a := [...]int{1,2,3,4,5,6,7}
	fmt.Printf("value of a: %v\n", a)
	fmt.Printf("lenght of a: %v\n", len(a))
	fmt.Printf("type of a: %T\n", a)

	fmt.Printf("value of USD: %v\n", USD)
	fmt.Printf("value of GBP: %v\n", GBP)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Printf("value of symbol: %v\n", symbol)
	fmt.Printf("type of symbol: %T\n", symbol)
	fmt.Println(RMB, symbol[RMB])

	// comparing arrays
	a1 := [2]int{1,2}
	b1 := [...]int{1,2}
	c1 := [2]int{1,3}
	fmt.Println(a1 == b1, a1 == c1, b1 == c1)
	d1 := [3]int{1,2,3}
	// we cannot compare d1 with the others because they are of different type
	fmt.Printf("value of d1: %v\n", d1)
	fmt.Printf("type of d1: %T\n", d1)
}

