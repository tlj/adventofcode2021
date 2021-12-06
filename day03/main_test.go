package main

import (
	"reflect"
	"strings"
	"testing"
)

var (
	exampleInput = strings.Split(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`, "\n")
)

func Test_findGammaBit(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "most of 1",
			args: args{
				input: "11100",
			},
			want: "1",
		},
		{
			name: "most of 0",
			args: args{
				input: "00100",
			},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGammaBit(tt.args.input); got != tt.want {
				t.Errorf("findGammaBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findEpsilonBit(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "most of 1",
			args: args{
				input: "11100",
			},
			want: "0",
		},
		{
			name: "most of 0",
			args: args{
				input: "00100",
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findEpsilonBit(tt.args.input); got != tt.want {
				t.Errorf("findGammaBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRate(t *testing.T) {
	type args struct {
		input    []string
		rateType rateType
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example, gamma",
			args: args{
				input:    exampleInput,
				rateType: gamma,
			},
			want: "10110",
		},
		{
			name: "example, epsilon",
			args: args{
				input:    exampleInput,
				rateType: epsilon,
			},
			want: "01001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRate(tt.args.input, tt.args.rateType); got != tt.want {
				t.Errorf("findRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	want := int64(198)
	if got := part1(exampleInput); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	want := int64(230)
	if got := part2(exampleInput); got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}

func Test_findMostCommonValueInBitPosition(t *testing.T) {
	type args struct {
		input    []string
		position int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example, first bit",
			args: args{
				input:    exampleInput,
				position: 0,
			},
			want: "1",
		},
		{
			name: "example, second bit",
			args: args{
				input:    exampleInput,
				position: 1,
			},
			want: "0",
		},
		{
			name: "equal = 1",
			args: args{
				input:    []string{"1", "0"},
				position: 0,
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMostCommonValueInBitPosition(tt.args.input, tt.args.position); got != tt.want {
				t.Errorf("findMostCommonValueInBitPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mostCommonBitsString(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example",
			args: args{
				input: exampleInput,
			},
			want: "10110",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostCommonBitsString(tt.args.input); got != tt.want {
				t.Errorf("mostCommonBitsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterOnBit(t *testing.T) {
	type args struct {
		input      []string
		bit        int
		ratingType ratingType
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example, first bit",
			args: args{
				input: exampleInput,
				bit:   0,
				ratingType: oxygenGeneratorRating,
			},
			want: []string{
				"11110",
				"10110",
				"10111",
				"10101",
				"11100",
				"10000",
				"11001",
			},
		},
		{
			name: "example, second bit",
			args: args{
				input: []string{
					"11110",
					"10110",
					"10111",
					"10101",
					"11100",
					"10000",
					"11001",
				},
				bit: 1,
				ratingType: oxygenGeneratorRating,
			},
			want: []string{
				"10110",
				"10111",
				"10101",
				"10000",
			},
		},
		{
			name: "example co2scrub, first bit",
			args: args{
				input: exampleInput,
				bit:   0,
				ratingType: co2ScrubberRating,
			},
			want: []string{
				"00100",
				"01111",
				"00111",
				"00010",
				"01010",
			},
		},
		{
			name: "example co2scrub, second bit",
			args: args{
				input: []string{
					"00100",
					"01111",
					"00111",
					"00010",
					"01010",
				},
				bit: 1,
				ratingType: co2ScrubberRating,
			},
			want: []string{
				"01111",
				"01010",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterOnBit(tt.args.input, tt.args.bit, tt.args.ratingType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterOnBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRating(t *testing.T) {
	type args struct {
		input []string
		ratingType ratingType
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example",
			args: args{
				input: exampleInput,
				ratingType: oxygenGeneratorRating,
			},
			want: "10111",
		},
		{
			name: "example",
			args: args{
				input: exampleInput,
				ratingType: co2ScrubberRating,
			},
			want: "01010",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRating(tt.args.input, tt.args.ratingType); got != tt.want {
				t.Errorf("findRating() = %v, want %v", got, tt.want)
			}
		})
	}
}
