package main

import (
	"adventofcode2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func createBoards(input []string) []*utils.BingoBoard {
	input = input[2:]

	var boards []*utils.BingoBoard
	var boardInput []string
	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			continue
		}

		boardInput = append(boardInput, input[i])
		if len(boardInput) == 5 {
			board := utils.NewBingoBoard()
			if err := board.CreateFromLines(boardInput); err != nil {
				panic(err)
			}
			boardInput = []string{}
			boards = append(boards, board)
		}
	}

	return boards
}

func part1(input []string) int {
	drawn := strings.Split(input[0], ",")
	boards := createBoards(input)

	for _, d := range drawn {
		dv, err := strconv.Atoi(d)
		if err != nil {
			panic(err)
		}
		utils.BingoFields[dv] = true

		for _, b := range boards {
			if b.HasBingo() {
				return b.SumUndrawn() * dv
			}
		}
	}

	return 0
}

func part2(input []string) int {
	drawn := strings.Split(input[0], ",")
	boards := createBoards(input)
	lastBingoSum := 0

	for _, d := range drawn {
		dv, err := strconv.Atoi(d)
		if err != nil {
			panic(err)
		}
		utils.BingoFields[dv] = true

		for _, b := range boards {
			if b.HasBingo() {
				lastBingoSum = b.SumUndrawn() * dv
			}
		}
	}

	return lastBingoSum
}

func main() {
	input, err := utils.LoadInput("day04/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}
