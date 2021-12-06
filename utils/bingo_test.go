package utils

import (
	"reflect"
	"testing"
)

func TestBingoBoard_CreateFromLines(t *testing.T) {
	type args struct {
		sizeX int
		sizeY int
		lines []string
	}
	tests := []struct {
		name       string
		args       args
		wantFields [][]int
		wantErr    bool
	}{
		{
			name: "Example 1",
			args: args{
				lines: []string{
					"22 13 17",
					" 8  2 23",
				},
				sizeX: 3,
				sizeY: 2,
			},
			wantFields: [][]int{{22, 13, 17}, {8, 2, 23}},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bb := NewBingoBoard()
			if err := bb.CreateFromLines(tt.args.lines); (err != nil) != tt.wantErr {
				t.Errorf("CreateFromLines() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(bb.Fields, tt.wantFields) {
				t.Errorf("CreateFromLines() got = %v, want %v", bb.Fields, tt.wantFields)
			}
		})
	}
}

func TestBingoBoard_HasBingo(t *testing.T) {
	type fields struct {
		Fields      [][]int
		BingoFields map[int]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Bingo Vertical",
			fields: fields{
				Fields: [][]int{{22, 13, 17}, {8, 2, 23}},
				BingoFields: map[int]bool{22: true, 13: true, 17: true},
			},
			want: true,
		},
		{
			name: "Bingo Horizontal",
			fields: fields{
				Fields: [][]int{{22, 13, 17}, {8, 2, 23}},
				BingoFields: map[int]bool{22: true, 8: true},
			},
			want: true,
		},
		{
			name: "No Bingo",
			fields: fields{
				Fields: [][]int{{22, 13, 17}, {8, 2, 23}},
				BingoFields: map[int]bool{22: true, 2: true},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bb := &BingoBoard{
				Fields: tt.fields.Fields,
			}
			BingoFields = tt.fields.BingoFields
			if got := bb.HasBingo(); got != tt.want {
				t.Errorf("HasBingo() = %v, want %v", got, tt.want)
			}
		})
	}
}
