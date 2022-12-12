// --- Day 5: Supply Stacks ---
package d05

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/eddeT/aoc-2022/input"
)

func splitInput(input []string) (stacks, instructions []string) {

	blankLine := false
	for _, line := range input {
		if line == "" {
			blankLine = true
			continue
		}
		if blankLine {
			instructions = append(instructions, line)
			continue

		}
		stacks = append(stacks, line)

	}

	return stacks, instructions
}

func parseStacks(rawStacks []string) [][]string {
	for i, j := 0, len(rawStacks)-1; i < j; i, j = i+1, j-1 {
		rawStacks[i], rawStacks[j] = rawStacks[j], rawStacks[i]
	}
	invertedStacks := rawStacks
	invertedStacks = invertedStacks[1:]

	re := regexp.MustCompile("[A-Z]|    ")
	var reorganized [][]string
	for _, invertedStack := range invertedStacks {
		matches := re.FindAllString(invertedStack, -1)
		for i, match := range matches {
			matches[i] = strings.Replace(match, "    ", "-", -1)
		}

		reorganized = append(reorganized, matches)

	}
	maxHeight := len(reorganized[0])
	maxWidth := len(reorganized)

	var finalStack [][]string
	for range reorganized[0] {
		finalStack = append(finalStack, nil)
	}

	for i := 0; i < maxHeight; i++ {
		for j := 0; j < maxWidth; j++ {
			if reorganized[j][i] != "-" {
				finalStack[i] = append(finalStack[i], reorganized[j][i])
			}
		}
	}

	return finalStack
}

func parseInstructions(rawInstructions []string) (parsedInstructions [][]int) {

	re := regexp.MustCompile("[0-9]{1,2}")
	for i, rawInstruction := range rawInstructions {
		parsedInstructions = append(parsedInstructions, nil)

		matches := re.FindAllString(rawInstruction, -1)
		var intMatches []int
		for _, y := range matches {
			k, err := strconv.Atoi(y)
			if err != nil {
				panic(err)
			}
			intMatches = append(intMatches, k)

		}

		parsedInstructions[i] = append(parsedInstructions[i], intMatches...)
	}
	return parsedInstructions

}

func moveCrates9000(supplyStacks [][]string, instructions [][]int) [][]string {
	for _, instruction := range instructions {

		source := instruction[1] - 1
		destination := instruction[2] - 1
		amount := instruction[0]
		for moves := 0; moves < amount; moves++ {
			sourceLen := len(supplyStacks[source]) - 1
			move := supplyStacks[source][sourceLen]
			supplyStacks[destination] = append(supplyStacks[destination], move)
			supplyStacks[source] = supplyStacks[source][:sourceLen]

		}

	}
	return supplyStacks

}

func moveCrates9001(supplyStacks [][]string, instructions [][]int) [][]string {
	for _, instruction := range instructions {

		source := instruction[1] - 1
		destination := instruction[2] - 1
		amount := instruction[0]
		max := len(supplyStacks[source]) - amount

		supplyStacks[destination] = append(supplyStacks[destination], supplyStacks[source][max:]...)
		supplyStacks[source] = supplyStacks[source][:len(supplyStacks[source])-amount]

	}
	return supplyStacks
}

func printSupplyTop(supplyStacks [][]string) {
	for i := 0; i < len(supplyStacks); i++ {
		lenght := len(supplyStacks[i]) - 1
		if lenght < 0 {
			continue
		}
		fmt.Print(supplyStacks[i][len(supplyStacks[i])-1])

	}
	fmt.Println("\n-----")
}

func solve(file string) ([][]string, [][]int) {
	raw := input.ReadLines("d05/" + file + ".input")
	rStacks, rInstructions := splitInput(raw)
	pStacks := parseStacks(rStacks)
	pInstructions := parseInstructions(rInstructions)

	return pStacks, pInstructions
}

func P1(file string) {
	printSupplyTop(moveCrates9000(solve(file)))
}
func P2(file string) {
	printSupplyTop(moveCrates9001(solve(file)))
}
