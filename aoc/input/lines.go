package input

import (
	"bytes"
	"iter"
)

type Lines [][]byte

// Height returns the height of input block in O(1)
func (lines Lines) Height() int {
	return len(lines)
}

// Height returns the width of the first line in O(1)
func (lines Lines) Width() int {
	return len(lines)
}

// WidthMax returns the minimum width of input block in O(n)
func (lines Lines) WidthMax() int {
	var maxLen int
	for _, l := range lines {
		if lineLen := len(l); lineLen > maxLen {
			maxLen = lineLen
		}
	}
	return maxLen
}

// WidthMin returns the minimum width of input block in O(n)
func (lines Lines) WidthMin() int {
	if len(lines) == 0 {
		return 0
	}
	minLen := len(lines[0])
	for _, l := range lines[1:] {
		if lineLen := len(l); lineLen < minLen {
			minLen = lineLen
		}
	}
	return minLen
}

func (lines Lines) HasConstantWidth() bool {
	if len(lines) == 0 {
		return true
	}
	fstLen := len(lines[0])
	for _, l := range lines[1:] {
		if len(l) != fstLen {
			return false
		}
	}
	return true
}

func (lines Lines) Contains(x, y int) bool {
	if y < 0 || y >= len(lines) {
		return false
	}
	if x < 0 || x >= len(lines[y]) {
		return false
	}
	return true
}

func (lines Lines) At(x, y int) byte {
	return lines[y][x]
}

func (lines Lines) Set(x, y int, v byte) {
	lines[y][x] = v
}

func (lines Lines) Transposed() Lines {
	if !lines.HasConstantWidth() {
		panic("(input.Lines).Transpose: Cannot call method if lines have different lengths")
	}
	transposed := make(Lines, lines.Width())
	for y := range lines.Width() {
		transposed[y] = make([]byte, lines.Height())
	}
	for x, y := range lines.Points() {
		transposed[y][x] = lines.At(x, y)
	}
	return transposed
}

func (lines Lines) Points() iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for y := range lines {
			for x := range lines[y] {
				if !yield(x, y) {
					return
				}
			}
		}
	}
}

func (lines Lines) WithOffset(skip int) iter.Seq2[int, int] {
	if !lines.HasConstantWidth() {
		panic("(input.Lines).WithOffset: Cannot call method if lines have different lengths")
	}
	return func(yield func(int, int) bool) {
		for y, line := range lines[skip : len(lines)-skip] {
			for x := range line[skip : len(line)-skip] {
				if !yield(x+skip, y+skip) {
					break
				}
			}
		}
	}
}

func (lines Lines) Flatten() []byte {
	var totalLen int
	for _, line := range lines {
		totalLen += len(line)
	}
	flattened := make([]byte, 0, totalLen)
	for x, y := range lines.Points() {
		flattened = append(flattened, lines.At(x, y))
	}
	return flattened
}

func (lines Lines) String() string {
	buf := new(bytes.Buffer)
	for _, line := range lines {
		buf.Write(line)
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (lines Lines) Copy() Lines {
	out := make(Lines, lines.Height())
	for i := range out {
		out[i] = make([]byte, lines.Width())
		copy(out[i], lines[i])
	}
	return out
}
