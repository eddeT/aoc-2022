package main

import (
	"fmt"

	"github.com/eddeT/aoc-2022/d01"
	"github.com/eddeT/aoc-2022/d02"
	"github.com/eddeT/aoc-2022/d03"
	"github.com/eddeT/aoc-2022/d04"
	"github.com/eddeT/aoc-2022/d05"
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
		d03.P1()
		d03.P2()
	case 4:
		d04.P0()
		d04.P1()
		d04.P2()
	case 5:
		d05.P1("0")
		d05.P1("1")
		d05.P2("0")
		d05.P2("1")
	default:
		fmt.Print("Input out of bounds")
	}

}
