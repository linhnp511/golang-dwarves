package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Format: <first_name> <last_name> <country_code>")
		return
	}

	firstName := os.Args[1]
	lastName := os.Args[2]
	countryCode := strings.ToLower(os.Args[3])

	reorderedName := reorderName(firstName, lastName, countryCode)
	fmt.Println("Reordered: ", reorderedName)
}

func reorderName(firstName, lastName, countryCode string) string {
	switch countryCode {
	case "us":
		return firstName + " " + lastName
	case "vn":
		return lastName + " " + firstName
	default:
		return "invalid country code"
	}
}