package six

import (
	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func countUnique(s string) int {
	found := make(map[rune]bool)
	for _, x := range s {
		if x != 10 {
			found[x] = true
		}
	}

	return len(found)
}

func SolvePart1() int {

	answers := inputhelper.GetStringsFromFileByDelim("inputs/input6.txt")
	count := 0

	for _, x := range answers {
		count += countUnique(x)
	}

	return count
}
