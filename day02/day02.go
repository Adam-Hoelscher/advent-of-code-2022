package day02

import (
	_ "embed"
	"strings"
)

//go:embed adventofcode.com_2022_day_2_input.txt
var Question string

func Answer() []int {
	return []int{solve(movesDirect(Question)), solve(movesOutcome(Question))}
}

func movesDirect(q string) [][]int {
	var moves [][]int
	for _, l := range strings.Split(q, "\n") {
		op, me := l[0]-'A', l[2]-'X'
		moves = append(moves, []int{int(me), int(op)})
	}
	return moves
}

func movesOutcome(q string) [][]int {
	var moves [][]int
	for _, l := range strings.Split(q, "\n") {
		op, me := l[0]-'A', (l[0]+l[2]-'A'-'X'+2)%3
		moves = append(moves, []int{int(me), int(op)})
	}
	return moves
}

func solve(moves [][]int) int {
	var total int
	for _, m := range moves {
		total += m[0] + 1 + (4+m[0]-m[1])%3*3
	}
	return total
}
