package day4

import "github.com/nesselchen/aoc-2024/aoc/input"

type Solver struct{}

const (
	xmas = "XMAS"
	samx = "SAMX"
)

func (Solver) First(lines input.Lines) int {
	var count int
	for line, l := range lines {
		for i := range l {
			count += checkPos(lines, line, i)
		}
	}
	return count
}

// determine whether to look for a word and for which one
func checkPos(lines input.Lines, line, i int) (count int) {
	var target string
	switch lines[line][i] {
	case xmas[0]:
		target = xmas
	case samx[0]:
		target = samx
	default:
		return 0
	}

	count += checkDirection(lines, target, line, 1, i, 0)  // horizontal
	count += checkDirection(lines, target, line, 0, i, 1)  // vertical
	count += checkDirection(lines, target, line, 1, i, 1)  // right diagonal
	count += checkDirection(lines, target, line, -1, i, 1) // left diagonal

	return count
}

// line, i determine position and dl, di the direction in which to look
func checkDirection(lines input.Lines, target string, line, dl, i, di int) int {
	w, h := len(lines[line]), len(lines)
	offset := len(target) - 1

	if maxPos := line + dl*offset; maxPos < 0 || maxPos >= h {
		return 0
	}
	if maxPos := i + di*offset; maxPos < 0 || maxPos >= w {
		return 0
	}
	for pos := range target {
		if target[pos] != lines[line+pos*dl][i+pos*di] {
			return 0
		}
	}

	return 1
}

func (Solver) Second(lines input.Lines) int {
	var (
		count int
		diff  byte = 'S' - 'M'
	)
	for x, y := range lines.WithOffset(1) {
		if lines.At(x, y) != 'A' {
			continue
		}
		var (
			topLeft  = lines.At(x-1, y-1)
			botRight = lines.At(x+1, y+1)
			topRight = lines.At(x+1, y-1)
			botLeft  = lines.At(x-1, y+1)
		)
		if abs(topLeft, botRight) != diff {
			continue
		}
		if abs(topRight, botLeft) != diff {
			continue
		}
		count++
	}

	return count
}

func abs(a, b byte) byte {
	if a < b {
		return b - a
	}
	return a - b
}
