package d04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/eddeT/aoc-2022/input"
)

type elfCleaner struct {
	low  int
	high int
}

func unsafeAtoi(number string) int {

	convertedInt, _ := strconv.Atoi(number)

	return convertedInt

}

func parsePair(pair string) (elfCleaner, elfCleaner) {
	minusCommaPair := strings.Replace(pair, ",", "-", -1)
	parsedPair := strings.Split(minusCommaPair, "-")

	elf1 := elfCleaner{
		low:  unsafeAtoi(parsedPair[0]),
		high: unsafeAtoi(parsedPair[1]),
	}
	elf2 := elfCleaner{
		low:  unsafeAtoi(parsedPair[2]),
		high: unsafeAtoi(parsedPair[3]),
	}
	return elf1, elf2
}

func compareCleaningOverlapFull(elf1, elf2 elfCleaner) int {

	overlap := 0
	if (elf1.low <= elf2.low && elf1.high >= elf2.high) || (elf2.low <= elf1.low && elf2.high >= elf1.high) {
		overlap = 1
	}
	return overlap
}

func compareCleaningOverlapAny(elf1, elf2 elfCleaner) int {
	overlap := 0
	match := false

	for i := elf1.low; i <= elf1.high; i++ {
		for j := elf2.low; j <= elf2.high; j++ {
			if j == i {
				match = true
				break
			}
		}
		if match {
			break
		}
	}

	if match {
		overlap = 1
	}
	return overlap
}

func calculateFullOverlap(pairs []string) {
	overlaps := 0
	for _, pair := range pairs {

		elf1, elf2 := parsePair(pair)

		overlaps += compareCleaningOverlapFull(elf1, elf2)
	}

	fmt.Println(overlaps)
}

func calculateAnyOverlap(pairs []string) {
	overlaps := 0

	for _, pair := range pairs {

		elf1, elf2 := parsePair(pair)
		overlaps += compareCleaningOverlapAny(elf1, elf2)
	}

	fmt.Println(overlaps)
}

func P0() {
	pairs := input.ReadLines("d04/0.input")
	calculateFullOverlap(pairs)
}

func P1() {
	pairs := input.ReadLines("d04/1.input")
	calculateFullOverlap(pairs)
}

func P2() {
	pairs := input.ReadLines("d04/0.input")
	calculateAnyOverlap(pairs)

	pairs = input.ReadLines("d04/1.input")
	calculateAnyOverlap(pairs)
}
