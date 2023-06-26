package day04

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed adventofcode.com_2022_day_4_input.txt
var Question string

func Answer() []int {
	return solve(Question)
}

func solve(q string) []int {

	var part, full int

	for _, line := range strings.Split(q, "\n") {

		pairs := convertIntervals(line)
		a, b := pairs[0], pairs[1]

		if a[0] == b[0] {
			part++
			full++
		} else if a[0] < b[0] && b[0] <= a[1] {
			if b[1] <= a[1] {
				full++
			}
			part++
		} else if b[0] < a[0] && a[0] <= b[1] {
			part++
			if a[1] <= b[1] {
				full++
			}
		}

	}

	return []int{full, part}
}

func convertIntervals(pair string) [2][2]int {

	strs := strings.FieldsFunc(pair, func(r rune) bool {
		return strings.ContainsRune(",-", r)
	})

	var vals [2][2]int
	for i, repr := range strs {
		v, _ := strconv.Atoi(repr)
		vals[i/2][i%2] = v
	}

	return vals
}
