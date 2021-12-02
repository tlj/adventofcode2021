package main

import (
	"adventofcode2021/utils"
	"fmt"
)

func part1(input []string) int {
	submarine := utils.NewSubmarine(0, 0, 0)
	if err := submarine.LoadInputStrings(input); err != nil {
		panic(err)
	}
	if err := submarine.RunFaulty(); err != nil {
		panic(err)
	}
	return submarine.HorizontalPosition * submarine.Depth
}

func part2(input []string) int {
	submarine := utils.NewSubmarine(0, 0, 0)
	if err := submarine.LoadInputStrings(input); err != nil {
		panic(err)
	}
	if err := submarine.Run(); err != nil {
		panic(err)
	}
	return submarine.HorizontalPosition * submarine.Depth
}

func main() {
	input, err := utils.LoadInput("day02/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}
