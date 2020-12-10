package ten

import (
	"sort"
	"strconv"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func SolvePart1() int {
	numberStrings := inputhelper.GetStrings("inputs/input10.txt")
	var numbers []int

	for _, x := range numberStrings {
		num, _ := strconv.Atoi(x)
		numbers = append(numbers, num)
	}

	sort.Ints(numbers)

	oneJolt := 0
	twoJolt := 0
	threeJolt := 0

	for i, x := range numbers {
		diff := 0
		if i == 0 {
			diff = x - 0
		} else {
			diff = x - numbers[i-1]
		}

		switch diff {
		case 1:
			oneJolt++
		case 2:
			twoJolt++
		case 3:
			threeJolt++
		}

		if i == len(numbers)-1 {
			threeJolt++
		}
	}

	return oneJolt * threeJolt
}
