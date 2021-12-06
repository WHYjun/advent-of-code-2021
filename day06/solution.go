package day06

import (
	"strconv"
	"strings"

	"github.com/WHYjun/advent-of-code-2021/utils"
)

func Part1(input string) int {
	return solve(input, 80)
}


func Part2(input string) int {
	return solve(input, 256)
}

func solve(input string, days int) int {
	cycles := getCyclesFromInput(input)
	counts := make(map[int]int)
	for i := 0; i < 9; i++ {
		counts[i] = 0
	}
	for _, cycle := range cycles {
		counts[cycle] += 1
	}
	for i := 0; i < days; i++ {
		newCounts := make(map[int]int)
		for timer, number := range counts {
			if timer == 0 {
				newCounts[6] += number
				newCounts[8] += number
			} else {
				newCounts[timer - 1] += number
			}
		}
		counts = newCounts
	}

	total := 0
	for _, number := range counts {
		total += number
	}
	return total
}

func getCyclesFromInput(input string) (cycles []int) {
	strCycles := strings.Split(input, ",")
	for _, strCycle := range strCycles {
		cycle, err := strconv.Atoi(strCycle)
		utils.PanicError(err)
		cycles = append(cycles, cycle)
	}
	return cycles
}
