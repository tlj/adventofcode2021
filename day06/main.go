package main

import (
	"adventofcode2021/utils"
	"fmt"
)

func incrementDays(fishes map[int]int, days int) map[int]int {
	for i := 0; i < days; i++ {
		fishes = increment(fishes)
	}

	return fishes
}

func increment(fishes map[int]int) map[int]int {
	newFishes := make(map[int]int, 10)

	for i := 1; i <= 8; i++ {
		if fishes[i] > 0 {
			newFishes[i-1] = fishes[i]
		}
	}

	if fishes[0] > 0 {
		newFishes[8] = fishes[0]
		newFishes[6] += fishes[0]
	}

	return newFishes
}

func part1(input map[int]int) int {
	fishes := incrementDays(input, 80)
	sum := 0
	for i := 0; i <= 8; i++ {
		sum += fishes[i]
	}
	return sum
}

func part2(input map[int]int) int {
	fishes := incrementDays(input, 256)
	sum := 0
	for i := 0; i <= 8; i++ {
		sum += fishes[i]
	}
	return sum
}

func main() {
	input, err := utils.LoadLineIntoInts("day06/input.txt")
	if err != nil {
		panic(err)
	}

	var daysLeft map[int]int
	daysLeft = make(map[int]int, 10)
	for _, v := range input {
		daysLeft[v]++
	}

	fmt.Printf("Part1: %d\n", part1(daysLeft))
	fmt.Printf("Part2: %d\n", part2(daysLeft))
}
