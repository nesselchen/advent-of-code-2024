package day3

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/nesselchen/aoc-2024/aoc/input"
)

type Solver struct{}

func (Solver) First(lines input.Lines) int {
	var count int
	for _, line := range lines {
		count += sumMuls(line)
	}
	return count
}

func (Solver) Second(lines input.Lines) int {
	var count int
	enabledChunks := strings.Split(string(lines.Flatten()), "do()")
	for _, chunk := range enabledChunks {
		splits := strings.SplitN(chunk, "don't()", 2)
		enabledPrefix := splits[0]
		count += sumMuls([]byte(enabledPrefix))
	}
	return count
}

var re = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func sumMuls(line []byte) int {
	var count int
	matches := re.FindAllSubmatch(line, -1)
	for _, m := range matches {
		a, _ := strconv.ParseInt(string(m[1]), 10, 64)
		b, _ := strconv.ParseInt(string(m[2]), 10, 64)
		count += int(a * b)
	}
	return count
}
