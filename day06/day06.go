package day06

import (
	_ "embed"
)

//go:embed adventofcode.com_2022_day_6_input.txt
var Question string

func Answer() []int {
	return []int{solve(Question, 4), solve(Question, 14)}
}

func solve(q string, n int) int {

	var count [256]int
	var uniq int

	for i := 0; i < len(q); i++ {

		if uniq == n {
			return i
		}

		count[q[i]]++
		if count[q[i]] == 1 {
			uniq++
		} else if count[q[i]] == 2 {
			uniq--
		}

		if i >= n {

			count[q[i-n]]--
			if count[q[i-n]] == 1 {
				uniq++
			} else if count[q[i-n]] == 0 {
				uniq--
			}

		}

	}

	panic("Not found!")
}
