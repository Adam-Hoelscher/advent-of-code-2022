package day01

import (
	"container/heap"
	_ "embed"
	"strconv"
	"strings"

	"github.com/Adam-Hoelscher/advent-of-code-2022/util"
)

//go:embed adventofcode.com_2022_day_1_input.txt
var Question string

func Answer() []int {
	return []int{solve(Question, 1), solve(Question, 3)}
}

func solve(q string, max int) int {

	h := new(util.MinIntHeap)

	var curr int
	for _, line := range strings.Split(q, "\n") {
		val, err := strconv.Atoi(line)
		if err != nil {
			heap.Push(h, curr)
			if h.Len() > max {
				heap.Pop(h)
			}
			curr = 0
		} else {
			curr += val
		}
	}

	heap.Push(h, curr)
	if h.Len() > max {
		heap.Pop(h)
	}

	var total int
	for _, x := range h.IntHeap {
		total += x
	}

	return total
}
