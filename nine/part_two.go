package nine

import (
	"strconv"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func SolvePart2() int {
	numbers := inputhelper.GetStrings("inputs/input9.txt")
	total := 32321523

	for i, x := range numbers {
		start, _ := strconv.Atoi(x)
		sum := start
		var keep []int
		keep = append(keep, start)
		for j := i + 1; j < len(numbers); j++ {
			curr, _ := strconv.Atoi(numbers[j])
			sum += curr
			keep = append(keep, curr)
			if total == sum {
				min, max := MinMax(keep)
				return min + max
			} else if sum > total {
				break
			}
		}
	}

	return 0
}
