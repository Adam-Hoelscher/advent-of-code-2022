package day07

import (
	_ "embed"
	"math"
	"strconv"
	"strings"
)

//go:embed adventofcode.com_2022_day_7_input.txt
var Question string

type dir struct {
	name     string
	parent   *dir
	children map[string]*dir
	size     int
}

func newDir(name string, parent *dir, isDir bool) *dir {
	d := new(dir)
	d.name = name
	if parent == nil {
		d.parent = d
	} else {
		d.parent = parent
	}
	if isDir {
		d.children = map[string]*dir{}
	}
	return d
}

func (d *dir) cd(path string) *dir {
	if path == ".." {
		return d.parent
	}
	return d.children[path]
}

func (d *dir) getSize() int {

	if d.size > 0 {
		return d.size
	}

	for _, ch := range d.children {
		d.size += ch.getSize()
	}

	return d.size
}

func (d *dir) ls(lines []string) {
	for _, l := range lines {
		words := strings.Split(l, " ")
		if words[0] == "dir" {
			d.children[words[1]] = newDir(words[1], d, true)
		} else {
			file := newDir(words[1], d, false)
			size, _ := strconv.Atoi(words[0])
			file.size = size
			d.children[words[1]] = file
		}
	}
}

func (d *dir) pwd() string {

	var names []string

	for d.parent != d {
		names = append(names, d.name)
		d = d.parent
	}

	path := new(strings.Builder)
	path.WriteByte('/')

	for i := len(names) - 1; i >= 0; i-- {
		path.WriteString(names[i])
		path.WriteByte('/')
	}

	return path.String()
}

func Answer() []int {
	return solve(Question, 100_000, 70_000_000-30_000_000)
}

func solve(q string, maxSize, maxUsed int) []int {

	root := newDir("", nil, true)
	d := root

	lines := strings.Split(q, "\n")[1:]
	n := len(lines)

	for i := 0; i < n; i++ {
		l := lines[i]

		cmd := l[2:]

		if cmd[:2] == "cd" {
			d = d.cd(cmd[3:])
		} else {
			j := i + 1
			for j < n {
				if lines[j][0] == '$' {
					break
				}
				j++
			}
			d.ls(lines[i+1 : j])
			i = j - 1
		}

	}

	needed := root.getSize() - maxUsed

	var sum int
	del := math.MaxInt

	queue := []*dir{root}
	for i := 0; i < len(queue); i++ {

		d := queue[i]

		if d.children == nil {
			continue
		}

		dSize := d.getSize()

		if dSize < maxSize {
			sum += dSize
		}

		if del > dSize && dSize > needed {
			del = dSize
		}

		for _, ch := range d.children {
			queue = append(queue, ch)
		}
	}

	return []int{sum, del}
}
