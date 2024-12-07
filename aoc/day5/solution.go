package day5

import (
	"slices"
	"strconv"
	"strings"

	"github.com/nesselchen/aoc-2024/aoc/input"
)

type Solver struct{}

// Part 1

func (Solver) SolveFirst(lines input.Lines) int {
	rules, pages := parseInput(lines)
	return solveAll(rules, pages)
}

func solveAll(rules map[int64][]int64, pages [][]int64) int {
	var sum int
	for _, p := range pages {
		if valid(rules, p) {
			sum += int(p[len(p)/2])
		}
	}
	return sum
}

func valid(rules map[int64][]int64, page []int64) bool {
	for i, n := range page {
		for _, bad := range rules[n] {
			if slices.Contains(page[:i], bad) {
				return false
			}
		}
	}
	return true
}

// Part 2

func (Solver) SolveSecond(lines input.Lines) int {
	rules, pages := parseInput(lines)
	return checkAndFixAll(rules, pages)
}

func checkAndFixAll(rules map[int64][]int64, pages [][]int64) int {
	var sum int
	for _, p := range pages {
		if !valid(rules, p) {
			fix(rules, p)
			sum += int(p[len(p)/2])
		}
	}
	return sum
}

var first = false

func fix(rules map[int64][]int64, page []int64) []int64 {
	for i := range page {
		n := page[0]
		for j := 0; j < len(page)-i-1; j++ {
			if slices.Contains(rules[n], page[j+1]) {
				n = page[j+1]
				continue
			}
			page[j], page[j+1] = page[j+1], page[j]
		}
	}
	return page
}

// Used by both

func parseInput(lines input.Lines) (map[int64][]int64, [][]int64) {
	rules := make(map[int64][]int64)
	pagesStart := 0
	for i, l := range lines {
		if len(l) == 0 {
			pagesStart = i + 1
			break
		}
		splits := strings.SplitN(string(l), "|", 2)
		fst, snd := splits[0], splits[1]
		before, err := strconv.ParseInt(fst, 10, 64)
		if err != nil {
			panic("Day 5: parsing error")
		}
		after, err := strconv.ParseInt(snd, 10, 64)
		if err != nil {
			panic("Day 5: parsing error")
		}
		rules[before] = append(rules[before], after)
	}
	pages := make([][]int64, 0)
	for i, l := range lines[pagesStart:] {
		splits := strings.Split(string(l), ",")
		pages = append(pages, make([]int64, 0))
		for _, s := range splits {
			n, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				panic("Day 5: parsing error")
			}
			pages[i] = append(pages[i], n)
		}
	}
	return rules, pages
}
