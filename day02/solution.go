package day02

import (
	"strconv"
	"strings"

	"github.com/WHYjun/advent-of-code-2021/utils"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	var position, depth int
	for _, line := range lines {
		commands := strings.Split(line, " ")
		if commands[0] == "forward" {
			n, err := strconv.Atoi(commands[1])
			utils.PanicError(err)
			position += n
		} else if commands[0] == "down" {
			n, err := strconv.Atoi(commands[1])
			utils.PanicError(err)
			depth += n
		} else {
			n, err := strconv.Atoi(commands[1])
			utils.PanicError(err)
			depth -= n
		}
	}
	return position * depth
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	var position, depth, aim int
	for _, line := range lines {
		commands := strings.Split(line, " ")
		if commands[0] == "forward" {
			n, err := strconv.Atoi(commands[1])
			utils.PanicError(err)
			position += n
			depth = depth + (n * aim)
		} else if commands[0] == "down" {
			n, err := strconv.Atoi(commands[1])
			utils.PanicError(err)
			aim += n
		} else {
			n, err := strconv.Atoi(commands[1])
			utils.PanicError(err)
			aim -= n
		}
	}
	return position * depth
}
