package main

import (
	"adventofcode2021/utils"
	"testing"
)

func Test_part1(t *testing.T) {
	input, err := utils.LoadInput("input.txt")
	if err != nil {
		panic(err)
	}
	want := 1524750
	if got := part1(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	input, err := utils.LoadInput("input.txt")
	if err != nil {
		panic(err)
	}
	want := 1592426537
	if got := part2(input); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}
