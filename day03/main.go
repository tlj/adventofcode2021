package main

import (
	"adventofcode2021/utils"
	"fmt"
	"strconv"
	"strings"
)

type rateType string

const (
	gamma   = "gamma"
	epsilon = "epsilon"
)

type ratingType string

const (
	co2ScrubberRating     = "co2scrubber"
	oxygenGeneratorRating = "oxygengenerator"
)

func findEpsilonBit(input string) string {
	if strings.Count(input, "0") > strings.Count(input, "1") {
		return "1"
	} else {
		return "0"
	}
}

func findGammaBit(input string) string {
	if strings.Count(input, "1") > strings.Count(input, "0") {
		return "1"
	} else {
		return "0"
	}
}

func findMostCommonValueInBitPosition(input []string, position int) string {
	occurrences := ""
	for _, in := range input {
		occurrences += string(in[position])
	}
	// if there are an equal number, return 1
	if strings.Count(occurrences, "0") > strings.Count(occurrences, "1") {
		return "0"
	} else {
		return "1"
	}
}

func mostCommonBitsString(input []string) string {
	var ret string
	for x := 0; x < len(input[0]); x++ {
		ret += findMostCommonValueInBitPosition(input, x)
	}
	return ret
}

func filterOnBit(input []string, bit int, ratingType2 ratingType) []string {
	var filtered []string
	mostCommonBits := mostCommonBitsString(input)
	for _, in := range input {
		if (ratingType2 == oxygenGeneratorRating && in[bit] == mostCommonBits[bit]) ||
			(ratingType2 == co2ScrubberRating && in[bit] != mostCommonBits[bit]) {
			filtered = append(filtered, in)
		}
	}
	return filtered
}

func findRate(input []string, rateType rateType) string {
	ret := ""
	if len(input) == 0 {
		return ret
	}
	size := len(input[0])
	for i := 0; i < size; i++ {
		bits := ""
		for _, in := range input {
			if len(in) < size {
				continue
			}
			bits += string(in[i])
		}
		var bit string
		switch rateType {
		case gamma:
			bit = findGammaBit(bits)
		case epsilon:
			bit = findEpsilonBit(bits)
		}
		ret += bit
	}
	return ret
}

func findRating(input []string, ratingType ratingType) string {
	i := 0
	for len(input) > 1 {
		input = filterOnBit(input, i, ratingType)
		i++
	}
	return input[0]
}

func part1(input []string) int64 {
	gammaValue, _ := strconv.ParseInt(findRate(input, gamma), 2, 64)
	epsilonValue, _ := strconv.ParseInt(findRate(input, epsilon), 2, 64)
	return gammaValue * epsilonValue
}

func part2(input []string) int64 {
	oxygenGeneratorValue, _ := strconv.ParseInt(findRating(input, oxygenGeneratorRating), 2, 64)
	co2ScrubberValue, _ := strconv.ParseInt(findRating(input, co2ScrubberRating), 2, 64)
	return oxygenGeneratorValue * co2ScrubberValue
}

func main() {
	input, err := utils.LoadInput("day03/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}
