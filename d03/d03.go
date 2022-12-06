// --- Day 3: Rucksack Reorganization ---
package d03

import (
	"fmt"
	"strings"

	"github.com/eddeT/aoc-2022/input"
)

func splitRucksack(rucksack string) (string, string) {
	rucksackLen := len(rucksack)
	rucksackHalfLen := (rucksackLen + 1) / 2
	firstCompartment := rucksack[0:rucksackHalfLen]
	secondCompartment := rucksack[rucksackHalfLen:rucksackLen]

	return firstCompartment, secondCompartment
}

func findCommonItem(first, second string) string {
	common := ""
	for _, c := range first {
		if strings.Contains(second, string(c)) {
			common = string(c)
		}
	}

	return common
}

func findCommonTrio(first, second, third string) byte {
	var common byte

	for _, c := range first {
		if strings.Contains(second, string(c)) && strings.Contains(third, string(c)) {
			common = byte(c)
		}
	}

	return common

}

func calculateSinglePriorty(item byte) int {
	if item > 90 {
		item -= 96
	} else {
		item -= 38
	}
	return int(item)
}

func P0() {
	rucksacks := input.ReadLines("d03/0.input")
	score := 0
	for _, rucksack := range rucksacks {
		commonItem := findCommonItem(splitRucksack(rucksack))
		score += calculateSinglePriorty(byte(commonItem[0]))
	}

	fmt.Println(score)

}

func P1() {
	rucksacks := input.ReadLines("d03/1.input")
	score := 0
	for _, rucksack := range rucksacks {
		commonItem := findCommonItem(splitRucksack(rucksack))
		score += calculateSinglePriorty(byte(commonItem[0]))
	}

	fmt.Println(score)
}

func P2() {
	rucksacks := input.ReadLines("d03/1.input")
	score := 0

	for i := 0; i < len(rucksacks); i += 3 {
		commonItem := findCommonTrio(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
		score += calculateSinglePriorty(byte(commonItem))
	}
	fmt.Println(score)
}
