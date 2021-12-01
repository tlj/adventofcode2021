package main

import "testing"

func Test_increaseCounter(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				data: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increaseCounter(tt.args.data); got != tt.want {
				t.Errorf("increaseCounter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumSlice(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "123",
			args: args{
				data: []int{1,2,3},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSlice(tt.args.data); got != tt.want {
				t.Errorf("sumSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_increseCounterOnSlidingWindow(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				data: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increaseCounterOnSlidingWindow(tt.args.data); got != tt.want {
				t.Errorf("increaseCounterOnSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}