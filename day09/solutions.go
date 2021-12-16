package day09

import (
	"github.com/WHYjun/advent-of-code-2021/utils"
)

func Part1(input string) int {
	intSlices := utils.ConvertInputToIntSlices(input)
	return findLowPointsAndCalculateRiskLevels(intSlices)
}

func Part2(input string) int {
	return 0
}

func findLowPointsAndCalculateRiskLevels(intSlices [][]int) int {
	riskLevels := 0
	for y, row := range intSlices {
		for x, val := range row {
			above := 9
			below := 9
			left := 9
			right := 9
			if y > 0 {
				y2 := y - 1
				x2 := x
				above = intSlices[y2][x2]
			}
			if y < len(intSlices)-1 {
				y2 := y + 1
				x2 := x
				below = intSlices[y2][x2]
			}
			if x > 0 {
				y2 := y
				x2 := x - 1
				left = intSlices[y2][x2]
			}
			if x < len(row)-1 {
				y2 := y
				x2 := x + 1
				right = intSlices[y2][x2]
			}
			if val < above && val < below && val < left && val < right {
				riskLevels += val + 1
			}
		}
	}
	return riskLevels
}