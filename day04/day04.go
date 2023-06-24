package day04

import (
	_ "embed"
	"strconv"

	_ "embed"
	"strings"
)

//go:embed adventofcode.com_2022_day_4_input.txt
var Question string

func Answer() []int {
	return []int{solve(Question)}
}

func solve(q string) int {

	var total int

	for _, pair := range strings.Split(q, "\n") {
		intervals := convertIntervals(pair)
		a, b := intervals[0], intervals[1]
		if a[0] >= b[0] && a[1] <= b[1] {
			total++
		} else if a[0] <= b[0] && a[1] >= b[1] {
			total++
		}

	}

	return total
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
