package eight

import (
	"regexp"
	"strconv"

	"github.com/joshsticks/AdventOfCode2020/inputhelper"
)

func runProgram(program []string) int {
	accumulator := 0
	visited := make(map[int]bool)
	i := 0
	re := regexp.MustCompile(`([a-z]{3})\s(\+|\-)([0-9]+)`)

	for !visited[i] {
		visited[i] = true
		instruction := re.FindStringSubmatch(program[i])
		switch command := instruction[1]; command {
		case "acc":
			val, _ := strconv.Atoi(instruction[3])
			if instruction[2] == "+" {
				accumulator += val
			} else {
				accumulator -= val
			}
			i++
		case "jmp":
			val, _ := strconv.Atoi(instruction[3])
			if instruction[2] == "+" {
				i += val
			} else {
				i -= val
			}
		case "nop":
			i++
		}
	}

	return accumulator
}

func SolvePart1() int {
	program := inputhelper.GetStrings("inputs/input8.txt")
	return runProgram(program)
}
