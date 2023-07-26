package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/linhnp511/go-23/exercise-02/sorting"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: sorttool <type> <values>")
		fmt.Println("Type: int | float | string")
		fmt.Println("Values: Comma-separated list of values")
		return
	}

	sortType := os.Args[1]
	values := os.Args[2]

	var sortedValues interface{}

	switch sortType {
	case "int":
		intValues := parseIntArray(values)
		sorting.SortInts(intValues)
		sortedValues = intValues
	case "float":
		floatValues := parseFloatArray(values)
		sorting.SortFloats(floatValues)
		sortedValues = floatValues
	case "string":
		stringValues := parseStringArray(values)
		sorting.SortStrings(stringValues)
		sortedValues = stringValues
	default:
		fmt.Println("Invalid type. Supported types: int, float, string")
		return
	}

	fmt.Println("Sorted result:", sortedValues)
}

func parseIntArray(input string) []int {
	strValues := strings.Split(input, ",")
	intValues := make([]int, len(strValues))

	for i, strVal := range strValues {
		intVal, err := strconv.Atoi(strings.TrimSpace(strVal))
		if err != nil {
			fmt.Println("Error parsing input:", err)
			os.Exit(1)
		}
		intValues[i] = intVal
	}

	return intValues
}

func parseFloatArray(input string) []float64 {
	strValues := strings.Split(input, ",")
	floatValues := make([]float64, len(strValues))

	for i, strVal := range strValues {
		floatVal, err := strconv.ParseFloat(strings.TrimSpace(strVal), 64)
		if err != nil {
			fmt.Println("Error parsing input:", err)
			os.Exit(1)
		}
		floatValues[i] = floatVal
	}

	return floatValues
}

func parseStringArray(input string) []string {
	strValues := strings.Split(input, ",")
	stringValues := make([]string, len(strValues))

	for i, strVal := range strValues {
		stringValues[i] = strings.TrimSpace(strVal)
	}

	return stringValues
}