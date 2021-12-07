package main

import "testing"

var exampleInput = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

func Test_findAverage(t *testing.T) {
	want := 5
	if got := findAverage(exampleInput); got != want {
		t.Errorf("findAverage() = %v, want %v", got, want)
	}
}

func Test_findMedian(t *testing.T) {
	want := 2
	if got := findMedian(exampleInput); got != want {
		t.Errorf("findMedian() = %v, want %v", got, want)
	}
}

func Test_part1(t *testing.T) {
	want := 37
	if got := part1(exampleInput); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	want := 168
	if got := part2(exampleInput); got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}
