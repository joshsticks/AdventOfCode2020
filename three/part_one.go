package three

import "github.com/joshsticks/AdventOfCode2020/inputhelper"

func SolvePart1() int {
	txt := inputhelper.GetStrings("inputs/input3.txt")
	count := 0

	y := 3
	for x := 1; x < len(txt); x++ {
		if y >= len(txt[x]) {
			y -= len(txt[x])
		}

		runes := []rune(txt[x])
		letterAt := string(runes[y])

		if letterAt == "#" {
			count++
		}

		y += 3
	}

	return count
}
