package day05

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/Adam-Hoelscher/advent-of-code-2022/util"
)

//go:embed adventofcode.com_2022_day_5_input.txt
var Question string

func Answer() []string {
	return []string{solveSingle(Question), solveMulti(Question)}
}

func parseMoves(lines []string) [][3]int {

	var moves [][3]int
	seps := []string{"move ", "from ", "to "}

	for _, ln := range lines {

		for _, s := range seps {
			ln = strings.ReplaceAll(ln, s, "")
		}

		var mv [3]int
		for i, num := range strings.Split(ln, " ") {
			val, _ := strconv.Atoi(num)
			mv[i] = val
		}

		moves = append(moves, mv)
	}

	return moves
}

func parseProblem(q string) ([]*util.ByteStack, [][3]int) {

	lines := strings.Split(q, "\n")
	var i int
	for lines[i][1] != '1' {
		i++
	}

	stacks := parseStart(lines[:i])
	moves := parseMoves(lines[i+2:])

	return stacks, moves
}

func parseStart(lines []string) []*util.ByteStack {

	n := len(lines[0])/4 + 1
	stacks := make([]*util.ByteStack, n)
	for i := range stacks {
		stacks[i] = new(util.ByteStack)
	}

	for i := len(lines) - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if lines[i][4*j+1] != ' ' {
				stacks[j].Push(lines[i][4*j+1])
			}
		}
	}

	return stacks
}

func readStacks(stacks []*util.ByteStack) string {

	ans := new(strings.Builder)
	for _, s := range stacks {
		ans.WriteByte(s.Peek())
	}

	return ans.String()
}

func solveMulti(q string) string {

	stacks, moves := parseProblem(q)

	temp := new(util.ByteStack)

	for _, mv := range moves {

		for i := 0; i < mv[0]; i++ {
			temp.Push(stacks[mv[1]-1].Pop())
		}

		for temp.Len() > 0 {
			stacks[mv[2]-1].Push(temp.Pop())
		}

	}

	return readStacks(stacks)
}

func solveSingle(q string) string {

	stacks, moves := parseProblem(q)

	for _, mv := range moves {

		for i := 0; i < mv[0]; i++ {
			stacks[mv[2]-1].Push(stacks[mv[1]-1].Pop())
		}

	}

	return readStacks(stacks)
}
