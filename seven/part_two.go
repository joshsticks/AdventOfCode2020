package seven

import (
	"regexp"
	"strconv"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func goDownBags(bags []string, bag string) int {
	count := 0
	for _, x := range bags {
		re1 := regexp.MustCompile(`([0-9])\s([a-z]+\s[a-z]+)\sbag`)
		matches := re1.FindAllStringSubmatch(x, -1)

		re2 := regexp.MustCompile(`^([a-z]+\s[a-z]+)\sbags`)
		parent := re2.FindStringSubmatch(x)

		if parent[1] == bag {
			for _, y := range matches {
				num, _ := strconv.Atoi(y[1])
				count += num
				count += num * goDownBags(bags, y[2])
			}
		}
	}

	return count
}

func SolvePart2() int {
	bags := inputhelper.GetStrings("inputs/input7.txt")
	return goDownBags(bags, "shiny gold")
}
