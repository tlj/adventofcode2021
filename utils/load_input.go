package utils

import (
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

func LoadInputIntoInts(filename string) ([]int, error) {
	lines, err := LoadInput(filename)
	if err != nil {
		return nil, err
	}
	return StringsToInts(lines)
}