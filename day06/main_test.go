package main

import (
	"reflect"
	"testing"
)

func Test_increment(t *testing.T) {
	type args struct {
		fishes map[int]int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "Example, after day 1",
			args: args{
				fishes: map[int]int{3: 2, 4: 1, 1: 1, 2: 1},
			},
			want: map[int]int{2: 2, 0: 1, 3: 1, 1: 1},
		},
		{
			name: "Example, after day 2",
			args: args{
				fishes: map[int]int{2: 2, 0: 1, 3: 1, 1: 1},
			},
			want: map[int]int{1: 2, 2: 1, 6: 1, 0: 1, 8: 1},
		},
		{
			name: "Example, after day 3",
			args: args{
				fishes: map[int]int{1: 2, 2: 1, 6: 1, 0: 1, 8: 1},
			},
			want: map[int]int{0: 2, 1: 1, 5: 1, 6: 1, 7: 1, 8: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increment(tt.args.fishes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("increment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_incrementDays(t *testing.T) {
	type args struct {
		fishes map[int]int
		days   int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "Example, day 2",
			args: args{
				fishes: map[int]int{3: 2, 4: 1, 1: 1, 2: 1},
				days:   2,
			},
			want: map[int]int{1: 2, 2: 1, 6: 1, 0: 1, 8: 1},
		},
		{
			name: "Example, day 3",
			args: args{
				fishes: map[int]int{3: 2, 4: 1, 1: 1, 2: 1},
				days:   3,
			},
			want: map[int]int{0: 2, 1: 1, 5: 1, 6: 1, 7: 1, 8: 1},
		},
		{
			name: "Example, day 4",
			args: args{
				fishes: map[int]int{3: 2, 4: 1, 1: 1, 2: 1},
				days:   4,
			},
			want: map[int]int{6: 3, 0: 1, 4: 1, 5: 1, 7: 1, 8: 2},
		},
		{
			name: "Example, day 8",
			args: args{
				fishes: map[int]int{3: 2, 4: 1, 1: 1, 2: 1},
				days:   8,
			},
			want: map[int]int{2: 3, 3: 2, 0: 1, 1: 1, 4: 2, 5: 1},
		},
		{
			name: "Example, day 18",
			args: args{
				fishes: map[int]int{3: 2, 4: 1, 1: 1, 2: 1},
				days:   18,
			},
			want: map[int]int{6: 5, 0: 3, 4: 2, 5: 1, 1: 5, 2: 3, 3: 2, 7: 1, 8: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := incrementDays(tt.args.fishes, tt.args.days); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("incrementDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
