package six

import (
	"strings"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func countUnanymous(s string) int {
	count := 0
	people := strings.Count(s, "\n") + 1
	found := make(map[rune]int)
	for _, x := range s {
		if x != 10 {
			found[x]++
		}
	}

	for _, x := range found {
		if x == people {
			count++
		}
	}

	return count
}

func SolvePart2() int {

	answers := inputhelper.GetStringsFromFileByDelim("inputs/input6.txt")
	count := 0

	for _, x := range answers {
		count += countUnanymous(x)
	}

	return count
}
