package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/WHYjun/advent-of-code-2021/utils"
)

type bingo [][]string

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	var numbers []string
	var boards []bingo
	var markedList [][]map[string]int
	i := 0
	for {
		if i > len(lines)-1 {
			break
		}
		if i == 0 {
			for _, number := range strings.Split(lines[0], ",") {
				numbers = append(numbers, number)
			}
			i++
		} else if lines[i] == "" {
			i++
		} else {
			boards = append(boards, convertToBingoBoard(lines[i:i+5]))
			markedList = append(markedList, []map[string]int{})
			i += 5
		}
	}

	for _, number := range numbers {
		marked := markOnBoards(boards, number)
		for i := 0; i < len(marked); i++ {
			markedList[i] = append(markedList[i], marked[i])
		}

		winner, ok := getWinnerIfExists(markedList)
		if ok {
			return calculateScore(boards[*winner], markedList[*winner])
		}
	}

	utils.PanicError(fmt.Errorf("no answer"))
	return 0
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	var numbers []string
	var boards []bingo
	var markedList [][]map[string]int
	i := 0
	for {
		if i > len(lines)-1 {
			break
		}
		if i == 0 {
			for _, number := range strings.Split(lines[0], ",") {
				numbers = append(numbers, number)
			}
			i++
		} else if lines[i] == "" {
			i++
		} else {
			boards = append(boards, convertToBingoBoard(lines[i:i+5]))
			markedList = append(markedList, []map[string]int{})
			i += 5
		}
	}

	var winners []int
	var winnerScores []int
	for _, number := range numbers {
		marked := markOnBoards(boards, number)
		for i := 0; i < len(marked); i++ {
			markedList[i] = append(markedList[i], marked[i])
		}

		ws := getWinnersIfExist(markedList)
		for _, w := range ws {
			found := false
			for _, k := range winners {
				if k == w {
					found = true
				}
			}
			if !found {
				winners = append(winners, w)
				winnerScore := calculateScore(boards[w], markedList[w])
				winnerScores = append(winners, winnerScore)
			}
		}
	}
	return winnerScores[len(winnerScores)-1:][0]
}

func getWinnersIfExist(markedList [][]map[string]int) (winners []int) {
	for i, marked := range markedList {
		rowCounter := map[int]int{
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
		}
		colCounter := map[int]int{
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
		}
		if len(marked) >= 5 {
			for _, mark := range marked {
				if mark != nil {
					rowCounter[mark["row"]]++
					colCounter[mark["col"]]++
				}
			}
			for _, v := range rowCounter {
				if v == 5 {
					found := false
					for _, w := range winners {
						if w == i {
							found = true
						}
					}
					if !found {
						winners = append(winners, i)
					}
				}
			}
			for _, v := range colCounter {
				if v == 5 {
					found := false
					for _, w := range winners {
						if w == i {
							found = true
						}
					}
					if !found {
						winners = append(winners, i)
					}
				}
			}
		}
	}
	return winners
}

func convertToBingoBoard(lines []string) bingo {
	var board bingo
	for _, row := range lines {
		var tmpRow []string
		row = strings.TrimPrefix(row, " ")
		row = strings.ReplaceAll(row, "  ", " ")
		numbers := strings.Split(row, " ")
		for _, number := range numbers {
			tmpRow = append(tmpRow, number)
		}
		board = append(board, tmpRow)
	}
	return board
}

func markOnBoards(boards []bingo, drawnNumber string) []map[string]int {
	var marked []map[string]int
	for i := 0; i < len(boards); i++ {
		isMarked := false
		board := boards[i]
		for rowIdx, row := range board {
			for colIdx, number := range row {
				if number == drawnNumber {
					marked = append(marked, map[string]int{
						"row": rowIdx,
						"col": colIdx,
					})
					isMarked = true
				}
			}
		}
		if !isMarked {
			marked = append(marked, nil)
		}
	}
	return marked
}

func getWinnerIfExists(markedList [][]map[string]int) (idx *int, ok bool) {
	for i, marked := range markedList {
		rowCounter := map[int]int{
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
		}
		colCounter := map[int]int{
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
		}
		if len(marked) >= 5 {
			for _, mark := range marked {
				if mark != nil {
					rowCounter[mark["row"]]++
					colCounter[mark["col"]]++
				}
			}
			for _, v := range rowCounter {
				if v == 5 {
					return &i, true
				}
			}
			for _, v := range colCounter {
				if v == 5 {
					return &i, true
				}
			}
		}
	}
	return nil, false
}

func calculateScore(board bingo, marked []map[string]int) int {
	var boardTotal int
	var lastNumber int
	for _, row := range board {
		for _, number := range row {
			numberInt, err := strconv.Atoi(number)
			utils.PanicError(err)
			boardTotal += numberInt
		}
	}

	for _, mark := range marked {
		if mark != nil {
			numberInt, err := strconv.Atoi(board[mark["row"]][mark["col"]])
			utils.PanicError(err)
			boardTotal -= numberInt
			lastNumber = numberInt
		}
	}
	return boardTotal * lastNumber
}