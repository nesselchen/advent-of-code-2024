package day4

type Solver struct{}

const (
	xmas = "XMAS"
	samx = "SAMX"
)

func (Solver) SolveFirst(lines []string) int {
	var count int
	for line, l := range lines {
		for i := range l {
			count += checkPos(lines, line, i)
		}
	}
	return count
}

// determine whether to look for a word and for which one
func checkPos(lines []string, line, i int) (count int) {
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
func checkDirection(lines []string, target string, line, dl, i, di int) int {
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

func (Solver) SolveSecond(lines []string) int {
	var (
		count int
		diff  byte = 'S' - 'M'
	)
	for lineNum, line := range lines[1 : len(lines)-1] {
		for i, c := range line[1 : len(line)-1] {
			if c != 'A' {
				continue
			}
			// positions are offset because iteration starts at 1
			var (
				topLeft  = lines[lineNum][i]
				botRight = lines[lineNum+2][i+2]
				topRight = lines[lineNum][i+2]
				botLeft  = lines[lineNum+2][i]
			)
			if abs(topLeft, botRight) != diff {
				continue
			}
			if abs(topRight, botLeft) != diff {
				continue
			}
			count++
		}
	}

	return count
}

func abs(a, b byte) byte {
	if a < b {
		return b - a
	}
	return a - b
}
