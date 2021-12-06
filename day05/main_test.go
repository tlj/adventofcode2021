package main

import (
	"testing"
)

var exampleInput = []string{
	"0,9 -> 5,9",
	"8,0 -> 0,8",
	"9,4 -> 3,4",
	"2,2 -> 2,1",
	"7,0 -> 7,4",
	"6,4 -> 2,0",
	"0,9 -> 2,9",
	"3,4 -> 1,4",
	"0,0 -> 8,8",
	"5,5 -> 8,2",
}

func Test_part1(t *testing.T) {
	want := 5
	if got := part1(exampleInput); got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	want := 12
	if got := part2(exampleInput); got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}

func Test_splitCoordinates(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "0,9",
			args: args{
				input: "0,9",
			},
			want:  0,
			want1: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitCoordinates(tt.args.input)
			if got != tt.want {
				t.Errorf("splitCoordinates() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("splitCoordinates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
