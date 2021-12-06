package main

import (
	"adventofcode2021/utils"
	"testing"
)

var exampleInput = []string{
	"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
	"",
	"22 13 17 11  0",
	"8  2 23  4 24",
	"21  9 14 16  7",
	"6 10  3 18  5",
	"1 12 20 15 19",
	"",
	"3 15  0  2 22",
	"9 18 13 17  5",
	"19  8  7 25 23",
	"20 11 10 24  4",
	"14 21 16 12  6",
	"",
	"14 21 17 24  4",
	"10 16 15  9 19",
	"18  8 23 26 20",
	"22 11 13  6  5",
	"2  0 12  3  7",
}

func Test_part1(t *testing.T) {
	utils.BingoFields = make(map[int]bool)
	want := 4512
	if got := part1(exampleInput); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	utils.BingoFields = make(map[int]bool)
	want := 1924
	if got := part2(exampleInput); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_createBoards(t *testing.T) {
	utils.BingoFields = make(map[int]bool)
	expectedNumberOfBoards := 3

	boards := createBoards(exampleInput)
	if len(boards) != expectedNumberOfBoards {
		t.Errorf("createBoards() want %v boards, got %v", expectedNumberOfBoards, len(boards))
	}

	boardSums := []int{300, 324, 325}
	for i := range boards {
		sum := boards[i].SumUndrawn()
		if sum != boardSums[i] {
			t.Errorf("createBoards() board %d want %v sum, got %v", i, boardSums[i], sum)
		}
	}
}
