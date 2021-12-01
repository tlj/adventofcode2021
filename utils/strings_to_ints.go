package utils

import (
	"fmt"
	"strconv"
)

func StringsToInts(data []string) ([]int, error) {
	var res []int

	for k, d := range data {
		if d == "" {
			continue
		}
		i, err := strconv.Atoi(d)
		if err != nil {
			return nil, fmt.Errorf("error in key %d (%s): %s", k, d, err)
		}
		res = append(res, i)
	}

	return res, nil
}