package main

import (
	"fmt"

	"github.com/eddeT/aoc-2022/d01"
	"github.com/eddeT/aoc-2022/d02"
	"github.com/eddeT/aoc-2022/d03"
)

func main() {

	var input int
	fmt.Printf("Enter the day you want to test [1-24]: ")
	fmt.Scan(&input)

	switch day := input; day {
	case 1:
		d01.P0()
		d01.P1()
		d01.P2()
	case 2:
		d02.P0()
		d02.P1()
		d02.P2()
	case 3:
		d03.P0()
		d03.P2()
	default:
		fmt.Print("Input out of bounds")
	}

}
