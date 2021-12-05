package day05

import (
	"strconv"
	"strings"
)

func Part1(input string) int {
	graph := make(map[int]map[int]int)
	answer := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		points := strings.Split(line, " -> ")
		startPoint := strings.Split(points[0], ",")
		endPoint := strings.Split(points[1], ",")
		x1, _ := strconv.Atoi(startPoint[0])
		y1, _ := strconv.Atoi(startPoint[1])
		x2, _ := strconv.Atoi(endPoint[0])
		y2, _ := strconv.Atoi(endPoint[1])
		deltaX, deltaY := 0, 0
		if x1 < x2 {
			deltaX = 1
		} else if x1 > x2{
			deltaX = -1
		}
		if y1 < y2 {
			deltaY = 1
		} else if y1 > y2{
			deltaY = -1
		}

		// skip if not horizontal or vertical
		if x1 != x2 && y1 != y2 {
			continue
		}

		for {
			if _, ok := graph[x1]; !ok {
				graph[x1] = make(map[int]int)
			}
			graph[x1][y1]++
			if x1 == x2 && y1 == y2 {
				break
			}
			x1 += deltaX
			y1 += deltaY
		}
	}

	for _, point := range graph {
		for _, count := range point {
			if count >= 2 {
				answer++
			}
		}
	}
	return answer
}

func Part2(input string) int {
	graph := make(map[int]map[int]int)
	answer := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		points := strings.Split(line, " -> ")
		startPoint := strings.Split(points[0], ",")
		endPoint := strings.Split(points[1], ",")
		x1, _ := strconv.Atoi(startPoint[0])
		y1, _ := strconv.Atoi(startPoint[1])
		x2, _ := strconv.Atoi(endPoint[0])
		y2, _ := strconv.Atoi(endPoint[1])
		deltaX, deltaY := 0, 0
		if x1 < x2 {
			deltaX = 1
		} else if x1 > x2{
			deltaX = -1
		}
		if y1 < y2 {
			deltaY = 1
		} else if y1 > y2{
			deltaY = -1
		}

		for {
			if _, ok := graph[x1]; !ok {
				graph[x1] = make(map[int]int)
			}
			graph[x1][y1]++
			if x1 == x2 && y1 == y2 {
				break
			}
			x1 += deltaX
			y1 += deltaY
		}
	}

	for _, point := range graph {
		for _, count := range point {
			if count >= 2 {
				answer++
			}
		}
	}
	return answer
}
