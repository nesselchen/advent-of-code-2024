package day6

import (
	"bytes"

	"github.com/nesselchen/aoc-2024/aoc/input"
)

type Solver struct{}

type direction struct {
	x, y int
}

var dirs = []direction{
	{0, -1}, // north
	{+1, 0}, // east
	{0, +1}, // south
	{-1, 0}, // west
}

func (d direction) String() string {
	switch d {
	case direction{0, -1}:
		return "^"
	case direction{+1, 0}:
		return ">"
	case direction{0, +1}:
		return "v"
	case direction{-1, 0}:
		return "<"
	}
	panic("unreachable")
}

func (Solver) First(grid input.Lines) int {
	var (
		marked  = 1
		x, y    = 0, 0
		currDir = 0
	)
	for i, row := range grid {
		idx := bytes.IndexByte(row, '^')
		if idx == -1 {
			continue
		}
		x, y = idx, i
		grid.Set(x, y, 'X')
		break
	}

	for {
		dir := dirs[currDir]
		nextX, nextY := x+dir.x, y+dir.y
		if !grid.Contains(nextX, nextY) {
			break
		}
		switch grid.At(nextX, nextY) {
		case '.':
			marked++
			grid.Set(nextX, nextY, 'X')
			x, y = nextX, nextY
		case 'X':
			x, y = nextX, nextY
		case '#':
			currDir = (currDir + 1) % 4
			dir = dirs[currDir]
		}
	}
	return marked
}

func (Solver) Second(grid input.Lines) int {
	var (
		loops          = 0
		x, y           = 0, 0
		startX, startY = 0, 0
		currDir        = 0
	)
	for i, row := range grid {
		idx := bytes.IndexByte(row, '^')
		if idx == -1 {
			continue
		}
		startX, startY = idx, i
		break
	}

	x, y = startX, startY
	for {
		dir := dirs[currDir]
		nextX, nextY := x+dir.x, y+dir.y
		if !grid.Contains(nextX, nextY) {
			break
		}
		switch grid.At(nextX, nextY) {
		case '.':
			if causesLoop(grid, x, y, currDir) {
				loops++
			}
			x, y = nextX, nextY
		case '^':
			x, y = nextX, nextY
		case '#':
			currDir = (currDir + 1) % 4
			dir = dirs[currDir]
		}
	}
	return loops
}

func causesLoop(grid input.Lines, x, y int, currDir int) bool {
	d := dirs[currDir]
	nextX, nextY := x+d.x, y+d.y
	original := grid.At(nextX, nextY)

	if grid.At(nextX, nextY) == '^' {
		return false
	}

	grid.Set(nextX, nextY, '#')
	defer grid.Set(nextX, nextY, original)

	type triple [3]int
	visited := make(map[triple]bool)

	for {
		nextX, nextY := x+d.x, y+d.y
		if !grid.Contains(nextX, nextY) {
			return false
		}
		switch grid.At(nextX, nextY) {
		case '.', '^':
			x, y = nextX, nextY
		case '#':
			t := triple{x, y, currDir}
			if found := visited[t]; found {
				return true
			}
			visited[t] = true
			currDir = (currDir + 1) % 4
			d = dirs[currDir]
		default:
			panic("You should never get here")
		}
	}
}
