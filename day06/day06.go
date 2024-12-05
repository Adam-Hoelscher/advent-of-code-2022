package day06

import (
	_ "embed"
	"math/bits"
)

//go:embed adventofcode.com_2022_day_6_input.txt
var Question string

func Answer() []int {
	return []int{solve(Question, 4), solve(Question, 14)}
}

func solve(q string, n int) int {

	var set uint

	for _, r := range q[:n] {
		set ^= 1 << (r % 26)
	}

	for i, r := range q[n:] {
		set ^= 1 << (r % 26)
		set ^= 1 << (q[i] % 26)
		if bits.OnesCount(set) == n {
			return i + n + 1
		}
	}

	panic("Not found!")
}
