// --- Day 2: Rock Paper Scissors ---

package d02

import (
	"fmt"

	"github.com/eddeT/aoc-2022/input"
)

/*
65 || 88 = Rock 1 point
66 || 89 = Paper 2 points
67 || 90 = Scissor 3 points
Loss = 0
Tie = 3
Win = 6

New rules:

*/

const (
	Rock     byte = 65
	Paper    byte = 66
	Scissors byte = 67
	Lose     byte = 88
	Draw     byte = 89
	Win      byte = 90
)

func outcomePoints(points *int, me, opponent byte) {

	if (me == Rock && opponent == Scissors) ||
		(me == Paper && opponent == Rock) ||
		(me == Scissors && opponent == Paper) {
		*points += 6
	} else if me == opponent {
		*points += 3
	}

	switch rpsPoints := me; rpsPoints {
	case Rock:
		*points += 1
	case Paper:
		*points += 2
	case Scissors:
		*points += 3
	}

}

func parseRPS1(singleStrategy string) int {
	opponent := byte(singleStrategy[0])
	me := byte(singleStrategy[2])
	me -= 23

	var points = 0

	outcomePoints(&points, me, opponent)

	return points

}

func parseRPS2(singleStrategy string) int {

	opponent := byte(singleStrategy[0])
	me := byte(singleStrategy[2])

	switch strategy := me; strategy {
	case Win:
		if opponent == Rock {
			me = Paper
		} else if opponent == Paper {
			me = Scissors
		} else if opponent == Scissors {
			me = Rock
		}
	case Draw:
		me = opponent
	case Lose:
		if opponent == Rock {
			me = Scissors
		} else if opponent == Paper {
			me = Rock
		} else if opponent == Scissors {
			me = Paper
		}
	}

	var points = 0

	outcomePoints(&points, me, opponent)
	return points

}

func calculateScore(strategyFile string, correctType bool) int {
	strategy := input.ReadLines("d02/" + strategyFile + ".input")

	sum := 0

	for _, l := range strategy {
		if correctType {
			sum += parseRPS2(l)
		} else {
			sum += parseRPS1(l)
		}

	}

	return sum
}

func P0() {
	fmt.Println(calculateScore("0", false))
}

func P1() {
	fmt.Println(calculateScore("1", false))
}

func P2() {
	fmt.Println(calculateScore("0", true))
	fmt.Println(calculateScore("1", true))
}
