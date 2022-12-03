package main

import (
	"fmt"

	"github.com/eddeT/aoc-2022/d01"
)

func main() {

	var input int
	fmt.Printf("Enter the day you want to test: ")
	fmt.Scan(&input)

	switch day := input; day {
	case 1:
		d01.P0()
		d01.P1()
		d01.P2()
	}
}
