package day6

import (
	"strings"

	"github.com/nesselchen/aoc-2024/aoc/input"
)

type Solver struct{}

type direction struct {
	X, Y int
}

var dirs = []direction{
	{0, -1}, // north
	{1, 0},  // east
	{0, 1},  // south
	{-1, 0}, // west
}

func (Solver) First(grid input.Lines) int {
	var (
		marked  = 1
		x, y    = 0, 0
		currDir = 0
	)
	for i, row := range grid {
		idx := strings.IndexByte(string(row), '^')
		if idx == -1 {
			continue
		}
		x, y = idx, i
		grid.Set(x, y, 'X')
	}

	for {
		dir := dirs[currDir]
		nextX, nextY := x+dir.X, y+dir.Y
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
