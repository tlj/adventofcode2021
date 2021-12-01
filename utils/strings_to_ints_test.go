package utils

import (
	"reflect"
	"testing"
)

func Test_stringsToInts(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "Everything is fine",
			args: args{
				data: []string{"1", "2", "3"},
			},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name: "Ignore empty",
			args: args{
				data: []string{"1", "2", "3", ""},
			},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringsToInts(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringsToInts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringsToInts() got = %v, want %v", got, tt.want)
			}
		})
	}
}
