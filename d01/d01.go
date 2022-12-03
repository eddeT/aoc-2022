package d01

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/eddeT/aoc-2022/input"
)

func parseElfCalories(inputFile string) []int {
	lines := input.ReadLines(inputFile)

	var elfCalories []int
	calories := 0
	for _, l := range lines {
		if l == "" {
			elfCalories = append(elfCalories, calories)

			calories = 0
		} else {
			caloriePart, _ := strconv.ParseInt(l, 10, 32)
			calories += int(caloriePart)
		}

	}

	return elfCalories

}

func mostCalories(elfCalories []int) int {
	total := len(elfCalories)
	for j := 1; j < total; j++ {
		if elfCalories[0] < elfCalories[j] {
			elfCalories[0] = elfCalories[j]
		}

	}
	return elfCalories[0]
}

func topThreeCalories(elfCalories []int) int {
	sort.Ints(elfCalories[:])
	return elfCalories[len(elfCalories)-1] +
		elfCalories[len(elfCalories)-2] +
		elfCalories[len(elfCalories)-3]
}

func puzzleSolution1(inputFile string) {
	inputFile = "d01/" + inputFile + ".input"
	fmt.Println(mostCalories(parseElfCalories(inputFile)))
}

func P0() {
	puzzleSolution1("0")
}

func P1() {
	puzzleSolution1("1")
}

func P2() {
	fmt.Println(topThreeCalories(parseElfCalories("d01/1.input")))
}
