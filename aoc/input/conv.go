package input

import (
	"strconv"
	"strings"
)

func ParseInt(s string) int {
	n, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic("input.ParseInt: Failed to convert string to number")
	}
	return int(n)
}

func SplitOnce(s, sep string) (before, after string, ok bool) {
	splits := strings.SplitN(s, sep, 2)
	if len(splits) != 2 {
		return "", "", false
	}
	return splits[0], splits[1], true
}
