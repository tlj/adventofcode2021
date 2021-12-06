package main

import (
	"adventofcode2021/utils"
	"fmt"
	"strconv"
	"strings"
)

func splitCoordinates(input string) (int, int) {
	parts := strings.Split(input, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])

	return x, y
}

func generateFloorMap(lines []string, allowDiagonal bool) [1000][1000]int {
	floorMap := [1000][1000]int{}

	for _, line := range lines {
		fromTo := strings.Split(line, " -> ")

		fromX, fromY := splitCoordinates(fromTo[0])
		toX, toY := splitCoordinates(fromTo[1])

		yModifier := 0
		if fromY > toY {
			yModifier = -1
		} else if fromY < toY {
			yModifier = 1
		}
		xModifier := 0
		if fromX > toX {
			xModifier = -1
		} else if fromX < toX {
			xModifier = 1
		}

		y := fromY
		x := fromX
		if fromX == toX {
			for y != toY {
				floorMap[y][fromX]++
				y += yModifier
			}
			floorMap[y][fromX]++
		} else if fromY == toY {
			for x != toX {
				floorMap[fromY][x]++
				x += xModifier
			}
			floorMap[fromY][x]++
		} else if allowDiagonal {
			for y != toY && x != toX {
				floorMap[y][x]++
				y += yModifier
				x += xModifier
			}
			floorMap[y][x]++
		}
	}

	return floorMap
}

func countCrossingLines(floorMap [1000][1000]int) int {
	count := 0
	for _, rows := range floorMap {
		for _, col := range rows {
			if col >= 2 {
				count++
			}
		}
	}
	return count
}

func part1(input []string) int {
	return countCrossingLines(generateFloorMap(input, false))
}

func part2(input []string) int {
	return countCrossingLines(generateFloorMap(input, true))
}

func main() {
	input, err := utils.LoadInput("day05/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}
