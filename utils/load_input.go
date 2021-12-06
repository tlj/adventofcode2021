package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func LoadInput(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func LoadLinesIntoInts(filename string) ([]int, error) {
	lines, err := LoadInput(filename)
	if err != nil {
		return nil, err
	}
	return StringsToInts(lines)
}

func LoadLineIntoInts(filename string) ([]int, error) {
	lines, err := LoadInput(filename)
	if err != nil {
		return nil, err
	}
	if len(lines) == 0 || lines[0] == "" {
		return nil, fmt.Errorf("no lines or empty first line of input is not valie for LoadLineIntoInts()")
	}
	fields := strings.Split(lines[0], ",")

	return StringsToInts(fields)
}
