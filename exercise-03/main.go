package main

import (
	"fmt"
)

func main() {
	// sample
	array := [][]int{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 1, 1},
		{0, 0, 1, 1},
	}

	count := countRectangles(array)
	fmt.Println("Number of rectangles:", count)
}

func countRectangles(array [][]int) int {
	rows := len(array)
	if rows == 0 {
		return 0
	}

	cols := len(array[0])
	if cols == 0 {
		return 0
	}

	rectangles := 0

	allOnes := func(arr []int, start, end int) bool {
		for i := start; i <= end; i++ {
			if arr[i] != 1 {
				return false
			}
		}
		return true
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if array[i][j] == 1 {
				for k := i; k < rows; k++ {
					if array[k][j] == 1 {
						for l := j; l < cols; l++ {
							if array[i][l] == 1 && allOnes(array[k], j, l) {
								rectangles++
							}
						}
					}
				}
			}
		}
	}

	return rectangles
}