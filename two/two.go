package two

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func GetRegexVariables(x string) (string, string, string, string) {
	re := regexp.MustCompile(`([0-9]|[1-9][0-9])-([0-9]|[1-9][0-9])\s([a-z]):\s(.*)$`)
	match := re.FindStringSubmatch(x)
	return match[1], match[2], match[3], match[4]
}

func GenerateRegex(x string) string {
	min, max, letter, _ := GetRegexVariables(x)

	return fmt.Sprintf(`:\s([^%[3]v]*[%[3]v]){%[1]v,%[2]v}[^%[3]v]*$`, min, max, letter)
}

func PasswordMatchesPolicy(x string) int {
	regex := GenerateRegex(x)
	re := regexp.MustCompile(regex)
	match := re.FindString(x)

	if len(match) > 0 {
		return 1
	}

	return 0
}

func PasswordMatchesPolicy2(x string) int {
	pos1, pos2, letter, password := GetRegexVariables(x)

	p1, _ := strconv.Atoi(pos1)
	p2, _ := strconv.Atoi(pos2)
	runes := []rune(password)
	letterAtP1 := string(runes[p1-1])
	letterAtP2 := string(runes[p2-1])

	if letterAtP1 == letter && letterAtP2 == letter {
		return 0
	} else if letterAtP1 == letter {
		return 1
	} else if letterAtP2 == letter {
		return 1
	}

	return 0
}

func SolvePart1() int {
	txt := inputhelper.GetStrings("inputs/input2.txt")
	count := 0

	for _, x := range txt {
		count += PasswordMatchesPolicy(x)
	}

	return count
}

func SolvePart2() int {
	txt := inputhelper.GetStrings("inputs/input2.txt")
	count := 0

	for _, x := range txt {
		count += PasswordMatchesPolicy2(x)
	}

	return count
}
