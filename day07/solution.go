package day07

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/WHYjun/advent-of-code-2021/utils"
)

func Part1(input string) int {
	positions := strings.Split(input, ",")
	var intPos []int
	for _, position := range positions {
		p, err := strconv.Atoi(position)
		utils.PanicError(err)
		intPos = append(intPos, p)
	}

	if intPos == nil {
		utils.PanicError(fmt.Errorf("invalid inputs"))
	}

	sort.Ints(intPos)
	mid := intPos[(len(intPos) / 2)]

	ans := 0
	for _, pos := range intPos {
		ans += utils.AbsInts(mid - pos)
	}
	return ans
}

func Part2(input string) int {
	var maxStarts, minStarts = 0, math.MaxInt64
	positions := strings.Split(input, ",")
	var intPos []int
	for _, position := range positions {
		p, err := strconv.Atoi(position)
		utils.PanicError(err)
		intPos = append(intPos, p)
	}

	for _, v := range intPos {
		maxStarts = utils.MaxInts(maxStarts, v)
		minStarts = utils.MinInts(minStarts, v)
	}
	moveCost := func(dist int) (cost int) {
		return dist * (dist + 1) / 2
	}
	var minCost = math.MaxInt64
	for i := minStarts; i < maxStarts; i++ {
		var cost int
		for _, v := range intPos {
			cost += moveCost(utils.AbsInts(i - v))
		}
		minCost = utils.MinInts(minCost, cost)
	}
	return minCost
}
