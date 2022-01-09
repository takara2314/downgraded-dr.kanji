package utils

import "strconv"

// AtoiSlice converts string slice to int slice
func AtoiSlice(s []string) ([]int, error) {
	result := make([]int, len(s))

	for i, str := range s {
		num, err := strconv.Atoi(str)
		if err != nil {
			return make([]int, 0), err
		}
		result[i] = num
	}

	return result, nil
}
