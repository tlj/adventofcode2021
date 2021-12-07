package main

import (
	"adventofcode2021/utils"
	"fmt"
	"math"
	"sort"
)

func findMedian(input []int) int {
	sort.Ints(input)
	mNumber := len(input) / 2
	median := input[mNumber]
	return median
}

func findAverage(input []int) int {
	total := 0.0
	for _, i := range input {
		total += float64(i)
	}
	return int(math.Round(total / float64(len(input))))
}

func part1(input []int) int {
	fuel := 0
	median := findMedian(input)
	for _, i := range input {
		fuel += int(math.Abs(float64(i) - float64(median)))
	}
	return fuel
}

func part2(input []int) int {
	fuel := 0
	average := findAverage(input)
	smallest := 999999999
	// Since the average isn't guaranteed to be the best spot, but it should be close enough, we'll check
	// the values around it to find the best one
	for closeValue := average - 1; closeValue < average+1; closeValue++ {
		for _, i := range input {
			distance := int(math.Abs(float64(i) - float64(closeValue)))
			inc := 1
			for x := 0; x < distance; x++ {
				fuel += inc
				inc++
			}
		}
		if fuel < smallest {
			smallest = fuel
		}
		// Starting to increase again so we know we found the smallest value
		if fuel > smallest {
			break
		}
		fuel = 0
	}

	return smallest
}

func main() {
	input, err := utils.LoadLineIntoInts("day07/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}
