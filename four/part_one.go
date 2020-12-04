package four

import (
	"strings"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func SolvePart1() int {

	passports := inputhelper.GetStringsFromFileByDelim("inputs/input4.txt")

	count := 0

	for _, x := range passports {
		if strings.Contains(x, "byr") && strings.Contains(x, "iyr") && strings.Contains(x, "eyr") && strings.Contains(x, "hgt") && strings.Contains(x, "hcl") && strings.Contains(x, "ecl") && strings.Contains(x, "pid") {
			count++
		}
	}

	return count
}
