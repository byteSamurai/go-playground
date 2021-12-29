package main

import (
	"fmt"
)


func main() {

	a := []int{9, 8, 2, 0, 7, 53, 38924, 4}

	res := make(chan []int, 100)

	go mergeSort(a, res);
	fmt.Println(<- res)

	close(res)
}

func mergeSort(s []int, result chan []int) {

	if len(s) <= 1 {
		result <- s
		return
	}

	n := len(s) / 2

	rightSliceChannel := make(chan []int, 100)
	leftSliceChannel := make(chan []int, 100)

	// !!! TODO recursive calls with correct slices

	// !!! TODO get slices from each channel

	// !!! TODO close channels

	// !!! TODO call merge() and send result to result channel
}

func merge(leftSlice []int, rightSlice []int) (result []int) {

	result = make([]int, len(leftSlice) + len(rightSlice))
	leftIndex, rightIndex := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case leftIndex >= len(leftSlice):
			result[i] = rightSlice[rightIndex]
			rightIndex++
		case rightIndex >= len(rightSlice):
			result[i] = leftSlice[leftIndex]
			leftIndex++
		case leftSlice[leftIndex] < rightSlice[rightIndex]:
			result[i] = leftSlice[leftIndex]
			leftIndex++
		default:
			result[i] = rightSlice[rightIndex]
			rightIndex++
		}
	}

	return
}
