package day08

import (
	"strings"
)

func Part1(input string) int {
	var one, four, seven, eight int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		entries := strings.Split(line, "|")
		output := entries[1]
		for _, value := range strings.Split(output, " ") {
			if len(value) == 2 {
				one += 1
			} else if len(value) == 4 {
				four += 1
			} else if len(value) == 3 {
				seven += 1
			} else if len(value) == 7 {
				eight += 1
			}
		}
	}

	return one + four + seven + eight
}

func Part2(input string) int {
	return 0
}
