package day01

import (
	"strings"

	"github.com/WHYjun/advent-of-code-2021/utils"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	numbers:= utils.MustConvertLinesToIntegerList(lines)

	count := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i - 1] < numbers[i] {
			count++
		}
	}
	return count
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	numbers:= utils.MustConvertLinesToIntegerList(lines)

	count := 0
	for i := 0; i < len(numbers) - 3; i++ {
		if numbers[i] < numbers[i+3] {
			count++
		}
	}
	return count
}
