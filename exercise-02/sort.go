package sorting

import (
	"sort"
)

//array of integers in asc
func SortInts(numbers []int) {
	sort.Ints(numbers)
}

// array of floats in asc
func SortFloats(numbers []float64) {
	sort.Float64s(numbers)
}

// array of strings in asc
func SortStrings(strings []string) {
	sort.Strings(strings)
}