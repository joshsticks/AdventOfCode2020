package three

import (
	"fmt"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func GetTrees(txt []string, right int, down int) int {
	count := 0

	y := right
	for x := down; x < len(txt); x += down {
		if y >= len(txt[x]) {
			y -= len(txt[x])
		}

		runes := []rune(txt[x])
		letterAt := string(runes[y])

		if letterAt == "#" {
			count++
		}

		y += right
	}

	return count
}

func SolvePart2() int {
	txt := inputhelper.GetStrings("inputs/input3.txt")

	count1 := GetTrees(txt, 1, 1)
	count2 := GetTrees(txt, 3, 1)
	count3 := GetTrees(txt, 5, 1)
	count4 := GetTrees(txt, 7, 1)
	count5 := GetTrees(txt, 1, 2)

	fmt.Println(count1)
	fmt.Println(count2)
	fmt.Println(count3)
	fmt.Println(count4)
	fmt.Println(count5)

	return count1 * count2 * count3 * count4 * count5
}
