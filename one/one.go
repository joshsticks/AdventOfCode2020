package one

import (
	"fmt"
	"strconv"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func FindTwentyTwenty(x int, y int) int {
	if x+y == 2020 {
		return x * y
	} else {
		return 0
	}
}

func FindTwentyTwentyThree(x int, y int, z int) int {
	if x+y+z == 2020 {
		return x * y * z
	} else {
		return 0
	}
}

func Solve() int {
	txt := inputhelper.GetStrings("inputs/input1.txt")
	fmt.Println(txt)

	for i, x := range txt {
		number1, _ := strconv.Atoi(x)
		for j := i + 1; j < len(txt); j++ {
			number2, _ := strconv.Atoi(txt[j])
			for k := j + 1; k < len(txt); k++ {
				number3, _ := strconv.Atoi(txt[k])
				result := FindTwentyTwentyThree(number1, number2, number3)
				if result > 0 {
					return result
				}
			}
		}
	}

	return 0
}
