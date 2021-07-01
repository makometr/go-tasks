package task30

import "errors"

// Task30 delete elen from slice
func Task30(slice []int, i int) ([]int, error) {
	if i >= len(slice) {
		return slice, errors.New("Index out of bound")
	}
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1], nil
}
