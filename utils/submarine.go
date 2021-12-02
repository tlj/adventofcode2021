package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type SubmarineInstruction string

const (
	SubmarineInstructionForward = "forward"
	SubmarineInstructionUp      = "up"
	SubmarineInstructionDown    = "down"
)

type Submarine struct {
	HorizontalPosition int
	Depth              int
	Aim                int
	commands           []SubmarineCommand
}

type SubmarineCommand struct {
	Command string
	Unit    int
}

func NewSubmarine(horizontalPosition, depth, aim int) *Submarine {
	return &Submarine{
		HorizontalPosition: horizontalPosition,
		Depth:              depth,
		Aim:                aim,
	}
}

func (s *Submarine) Commands() []SubmarineCommand {
	return s.commands
}

func (s *Submarine) LoadInputStrings(commands []string) error {
	for k, c := range commands {
		if c == "" {
			continue
		}
		parts := strings.Split(c, " ")
		if len(parts) != 2 {
			return fmt.Errorf("'%s' on line %d split into %d parts, expected 2", c, k, len(parts))
		}
		units, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("'%s' on line %d is not an int", parts[1], k)
		}
		s.commands = append(s.commands, SubmarineCommand{Command: parts[0], Unit: units})
	}
	return nil
}

func (s *Submarine) RunFaultyCommand(command SubmarineCommand) error {
	switch command.Command {
	case SubmarineInstructionForward:
		s.HorizontalPosition += command.Unit
	case SubmarineInstructionDown:
		s.Depth += command.Unit
	case SubmarineInstructionUp:
		s.Depth -= command.Unit
	default:
		return fmt.Errorf("unsupported instruction %s", command.Command)
	}
	return nil
}

func (s *Submarine) RunFaulty() error {
	for k, c := range s.commands {
		if err := s.RunFaultyCommand(c); err != nil {
			return fmt.Errorf("error in command %d (%s %d): %s", k, c.Command, c.Unit, err)
		}
	}
	return nil
}

func (s *Submarine) RunCommand(command SubmarineCommand) error {
	switch command.Command {
	case SubmarineInstructionForward:
		s.HorizontalPosition += command.Unit
		s.Depth += s.Aim * command.Unit
	case SubmarineInstructionDown:
		s.Aim += command.Unit
	case SubmarineInstructionUp:
		s.Aim -= command.Unit
	default:
		return fmt.Errorf("unsupported instruction %s", command.Command)
	}
	return nil
}

func (s *Submarine) Run() error {
	for k, c := range s.commands {
		if err := s.RunCommand(c); err != nil {
			return fmt.Errorf("error in command %d (%s %d): %s", k, c.Command, c.Unit, err)
		}
	}
	return nil
}
