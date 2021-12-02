package utils

import (
	"reflect"
	"testing"
)

var (
	submarineDay02Example = []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	submarineDay02ExampleCommands = []SubmarineCommand{
		{Command: "forward", Unit: 5},
		{Command: "down", Unit: 5},
		{Command: "forward", Unit: 8},
		{Command: "up", Unit: 3},
		{Command: "down", Unit: 8},
		{Command: "forward", Unit: 2},
	}
)

func TestSubmarine_LoadInputStrings(t *testing.T) {
	type args struct {
		commands []string
	}
	tests := []struct {
		name    string
		args    args
		want    []SubmarineCommand
		wantErr bool
	}{
		{
			name: "example",
			args: args{
				commands: submarineDay02Example,
			},
			want: submarineDay02ExampleCommands,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Submarine{}
			if err := s.LoadInputStrings(tt.args.commands); (err != nil) != tt.wantErr {
				t.Errorf("LoadInputStrings() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := s.Commands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Commands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubmarine_RunFaultyCommand(t *testing.T) {
	type fields struct {
		HorizontalPosition int
		Depth              int
		commands           []SubmarineCommand
	}
	type args struct {
		command SubmarineCommand
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Submarine
		wantErr bool
	}{
		{
			name:   "forward 5",
			fields: fields{},
			args: args{
				command: SubmarineCommand{
					Command: SubmarineInstructionForward,
					Unit:    5,
				},
			},
			want:    Submarine{HorizontalPosition: 5},
			wantErr: false,
		},
		{
			name:   "down 5",
			fields: fields{},
			args: args{
				command: SubmarineCommand{
					Command: SubmarineInstructionDown,
					Unit:    5,
				},
			},
			want:    Submarine{Depth: 5},
			wantErr: false,
		},
		{
			name: "up 3",
			fields: fields{
				Depth: 5,
			},
			args: args{
				command: SubmarineCommand{
					Command: SubmarineInstructionUp,
					Unit:    3,
				},
			},
			want:    Submarine{Depth: 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Submarine{
				HorizontalPosition: tt.fields.HorizontalPosition,
				Depth:              tt.fields.Depth,
				commands:           tt.fields.commands,
			}
			if err := s.RunFaultyCommand(tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("RunFaultyCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := *s; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Submarine = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubmarine_RunFaulty(t *testing.T) {
	type fields struct {
		HorizontalPosition int
		Depth              int
		commands           []SubmarineCommand
	}
	tests := []struct {
		name    string
		fields  fields
		want    Submarine
		wantErr bool
	}{
		{
			name: "example",
			fields: fields{
				HorizontalPosition: 0,
				Depth:              0,
				commands: submarineDay02ExampleCommands,
			},
			want: Submarine{
				HorizontalPosition: 15,
				Depth:              10,
				commands: submarineDay02ExampleCommands,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Submarine{
				HorizontalPosition: tt.fields.HorizontalPosition,
				Depth:              tt.fields.Depth,
				commands:           tt.fields.commands,
			}
			if err := s.RunFaulty(); (err != nil) != tt.wantErr {
				t.Errorf("RunFaulty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSubmarine_RunCommand(t *testing.T) {
	type fields struct {
		HorizontalPosition int
		Depth              int
		Aim                int
		commands           []SubmarineCommand
	}
	type args struct {
		command SubmarineCommand
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Submarine
		wantErr bool
	}{
		{
			name:   "forward 5",
			fields: fields{},
			args: args{
				command: SubmarineCommand{
					Command: SubmarineInstructionForward,
					Unit:    5,
				},
			},
			want:    Submarine{HorizontalPosition: 5},
			wantErr: false,
		},
		{
			name: "forward 5, with aim",
			fields: fields{
				Aim: 2,
			},
			args: args{
				command: SubmarineCommand{
					Command: SubmarineInstructionForward,
					Unit:    5,
				},
			},
			want:    Submarine{HorizontalPosition: 5, Depth: 10, Aim: 2},
			wantErr: false,
		},
		{
			name:   "down 5",
			fields: fields{},
			args: args{
				command: SubmarineCommand{
					Command: SubmarineInstructionDown,
					Unit:    5,
				},
			},
			want:    Submarine{Aim: 5},
			wantErr: false,
		},
		{
			name: "forward 8",
			fields: fields{
				HorizontalPosition: 5,
				Depth: 0,
				Aim: 5,
			},
			args: args{
				command: SubmarineCommand{
					Unit:    8,
					Command: SubmarineInstructionForward,
				},
			},
			want:    Submarine{HorizontalPosition: 13, Depth: 40, Aim: 5},
			wantErr: false,
		},
		{
			name: "up 3",
			fields: fields{
				Aim: 5,
			},
			args: args{
				command: SubmarineCommand{
					Command: SubmarineInstructionUp,
					Unit:    3,
				},
			},
			want:    Submarine{Aim: 2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Submarine{
				HorizontalPosition: tt.fields.HorizontalPosition,
				Depth:              tt.fields.Depth,
				Aim:                tt.fields.Aim,
				commands:           tt.fields.commands,
			}
			if err := s.RunCommand(tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("RunFaultyCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := *s; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Submarine = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubmarine_Run(t *testing.T) {
	type fields struct {
		HorizontalPosition int
		Depth              int
		Aim                int
		commands           []SubmarineCommand
	}
	tests := []struct {
		name    string
		fields  fields
		want    Submarine
		wantErr bool
	}{
		{
			name: "example",
			fields: fields{
				HorizontalPosition: 0,
				Depth:              0,
				Aim:                0,
				commands: submarineDay02ExampleCommands,
			},
			want: Submarine{
				HorizontalPosition: 15,
				Depth:              60,
				Aim:                10,
				commands: []SubmarineCommand{
					{"forward", 5},
					{"down", 5},
					{"forward", 8},
					{"up", 3},
					{"down", 8},
					{"forward", 2},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Submarine{
				HorizontalPosition: tt.fields.HorizontalPosition,
				Depth:              tt.fields.Depth,
				Aim:                tt.fields.Aim,
				commands:           tt.fields.commands,
			}
			if err := s.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := *s; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Submarine = %v, want %v", got, tt.want)
			}
		})
	}
}
