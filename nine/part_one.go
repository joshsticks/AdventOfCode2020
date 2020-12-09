package nine

import (
	"strconv"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func SolvePart1() int {
	numbers := inputhelper.GetStrings("inputs/input9.txt")

	previousSums := make(map[int]bool)
	var found int
	for i := 1; i < len(numbers); i++ {
		num, _ := strconv.Atoi(numbers[i])
		if i > 24 {
			value, _ := previousSums[num]
			if !value {
				found = num
				break
			}
		}
		for j := i - 1; j >= 0; j-- {
			num2, _ := strconv.Atoi(numbers[j])
			previousSums[num+num2] = true
		}
	}

	return found
}
