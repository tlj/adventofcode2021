package main

import (
	"adventofcode2021/utils"
	"fmt"
)

func increaseCounter(data []int) int {
	res := 0
	for x := 1; x < len(data); x++ {
		if data[x] > data[x-1] {
			res++
		}
	}
	return res
}

func sumSlice(data []int) int {
	res := 0
	for _, i := range data {
		res += i
	}
	return res
}

func increaseCounterOnSlidingWindow(data []int) int {
	res := 0
	for x := 1; x < len(data) - 2; x++ {
		nextWindowSum := sumSlice(data[x:x+3])
		previousWindowSum := sumSlice(data[x-1:x+2])
		if nextWindowSum > previousWindowSum {
			res++
		}
	}
	return res
}

func part1(input []int) int {
	return increaseCounter(input)
}

func part2(input []int) int {
	return increaseCounterOnSlidingWindow(input)
}

func main() {
	input, err := utils.LoadInputIntoInts("day01/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}
