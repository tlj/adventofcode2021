package utils

import (
	"strconv"
	"strings"
)

var BingoFields map[int]bool

type BingoBoard struct {
	Fields         [][]int
	AlreadyBingoed bool
}

func NewBingoBoard() *BingoBoard {
	b := BingoBoard{}
	if BingoFields == nil {
		BingoFields = make(map[int]bool)
	}
	return &b
}

func (bb *BingoBoard) CreateFromLines(lines []string) error {
	y := -1
	for _, l := range lines {
		cols := strings.Fields(l)
		if len(cols) == 0 {
			continue
		}
		y++
		bb.Fields = append(bb.Fields, []int{})
		for _, c := range cols {
			c, err := strconv.Atoi(c)
			if err != nil {
				return err
			}
			if _, ok := BingoFields[c]; !ok {
				BingoFields[c] = false
			}
			bb.Fields[y] = append(bb.Fields[y], c)
		}
	}
	return nil
}

func (bb *BingoBoard) HasBingo() bool {
	if bb.AlreadyBingoed {
		return false
	}
	// vertical
	for x := range bb.Fields {
		found := 0
		for y := range bb.Fields[x] {
			if BingoFields[bb.Fields[x][y]] {
				found++
			} else {
				break
			}
		}
		if found == len(bb.Fields[x]) {
			bb.AlreadyBingoed = true
			return true
		}
	}
	// horizontal
	for y := 0; y < len(bb.Fields[0]); y++ {
		found := 0
		for x := range bb.Fields {
			if BingoFields[bb.Fields[x][y]] {
				found++
			} else {
				break
			}
		}
		if found == len(bb.Fields) {
			bb.AlreadyBingoed = true
			return true
		}
	}
	return false
}

func (bb *BingoBoard) SumUndrawn() int {
	sum := 0
	for _, l := range bb.Fields {
		for _, v := range l {
			if !BingoFields[v] {
				sum += v
			}
		}
	}
	return sum
}
