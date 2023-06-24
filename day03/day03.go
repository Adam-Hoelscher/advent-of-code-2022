package day03

import (
	_ "embed"
	"strings"
)

//go:embed adventofcode.com_2022_day_3_input.txt
var Question string

func Answer() []int {
	return solve(Question)
}

func solve(q string) []int {

	rucksacks := strings.Split(q, "\n")

	var itemTotal, tokenTotal int
	var item, token rune
	var group map[rune]int

	for i, sack := range rucksacks {

		g := i % 3
		if g == 0 {
			group = map[rune]int{}
		}

		set := map[rune]int{}
		m := len(sack) / 2

		for _, r := range sack[:m] {
			set[r] |= 1
			group[r] |= 1 << g
		}

		for i, r := range sack {
			set[r] |= 1 << (i / m)
			if set[r] == 3 {
				item = r
			}
			group[r] |= 1 << g
			if group[r] == 7 {
				token = r
			}
		}

		// fmt.Println("item", string(item), value(item))
		itemTotal += value(item)
		if g == 2 {
			// fmt.Println("token", string(token))
			tokenTotal += value(token)
		}

	}

	return []int{itemTotal, tokenTotal}
}

func value(r rune) int {
	if 'a' <= r && r <= 'z' {
		return int(r - 'a' + 1)
	} else {
		return int(r - 'A' + 27)
	}
}
