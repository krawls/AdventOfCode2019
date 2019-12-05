package main

import (
	"fmt"
)

// global variables
var debugMessages = false // turn on verbose debug messages

func validatePart1(input int, lRange int, uRange int) bool {
	if !validateDigits(input) || !validateRange(input, lRange, uRange) || !validateAdjacent(input) || !validateNeverDecrease(input) {
		return false
	}
	return true
}

func validatePart2(input int, lRange int, uRange int) bool {
	if !validateDigits(input) || !validateRange(input, lRange, uRange) || !validateAdjacentExact(input) || !validateNeverDecrease(input) {
		return false
	}
	return true
}

func validateDigits(input int) bool {
	// validate if the given number is no more than 6 digits
	if input <= 999999 {
		return true
	}
	return false
}

func validateRange(input int, lRange int, uRange int) bool {
	// validate if the given number is within the given range
	if input >= lRange && input <= uRange {
		return true
	}
	return false
}

func validateAdjacent(input int) bool {
	// validate if any two adjacent digits are the same

	// copy input into intermediate variable so we can manipulate it
	var n = input

	// slice given input into digits
	a := make([]int, 0)
	for n != 0 {
		//since we're dividing, we need to prepend using a composite literal
		a = append([]int{n % 10}, a...)
		n /= 10
		if debugMessages {
			fmt.Printf("Current slice: %v\n", a)
		}
	}

	// check for adjacent equality
	for i := range a {
		// don't allow checking beyond end of slice
		if i == len(a)-1 {
			if debugMessages {
				fmt.Printf("Input: %d - did not find two adjacent values\n", input)
			}
			return false
		}
		if a[i] == a[i+1] {
			if debugMessages {
				fmt.Printf("Input: %d - found two adjacent values: %d and %d\n", input, a[i], a[i+1])
			}
			return true
		}
	}
	// shouldn't get here
	return false
}

func validateAdjacentExact(input int) bool {
	// validate if any two adjacent digits are the same
	// the two adjacent digits must be exactly a group of 2

	// copy input into intermediate variable so we can manipulate it
	var n = input

	// slice given input into digits
	a := make([]int, 0)
	for n != 0 {
		//since we're dividing, we need to prepend using a composite literal
		a = append([]int{n % 10}, a...)
		n /= 10
		if debugMessages {
			fmt.Printf("Current slice: %v\n", a)
		}
	}

	// check for exact adjacent equality
	var found = false
	for i := range a {
		if (i <= len(a)-3 && i > 0) && (a[i] == a[i+1] && a[i] != a[i+2] && a[i] != a[i-1]) { // check middle case
			found = true
		} else if (i == len(a)-2) && (a[i] == a[i+1] && a[i] != a[i-1]) { // check end edge case
			found = true
		} else if (i == 0) && (a[i] == a[i+1] && a[i] != a[i+2]) { // check beginning edge case
			found = true
		} else {
			found = false
		}
		if found {
			return true
		}
	}
	return false
}

func validateNeverDecrease(input int) bool {
	// validate if (left to right) the digits never decrease
	// only increase or stay the same

	// copy input into intermediate variable so we can manipulate it
	var n = input

	// slice given input into digits
	a := make([]int, 0)
	for n != 0 {
		//since we're dividing, we need to prepend using a composite literal
		a = append([]int{n % 10}, a...)
		n /= 10
		if debugMessages {
			fmt.Printf("Current slice: %v\n", a)
		}
	}

	// check for adjacent increase/equality
	for i := range a {
		// check for end of slice
		if i == len(a)-1 {
			if debugMessages {
				fmt.Printf("Input: %d - all values in increasing order\n", input)
			}
			return true
		}
		if a[i] > a[i+1] {
			if debugMessages {
				fmt.Printf("Input: %d - found two adjacent values non-increasing/non-equal values: %d and %d\n", input, a[i], a[i+1])
			}
			return false
		}
	}
	// shouldn't get here
	return false
}

func main() {
	// given constraints
	var lowerRange = 278384
	var upperRange = 824795

	var countPart1 = 0
	var countPart2 = 0

	for i := lowerRange; i <= upperRange; i++ {
		// Part 1
		resPart1 := validatePart1(i, lowerRange, upperRange)
		if debugMessages {
			fmt.Printf("Input: %d - Valid: %t\n", i, resPart1)
		}
		if resPart1 {
			countPart1++
		}

		// Part 2
		resPart2 := validatePart2(i, lowerRange, upperRange)
		if debugMessages {
			fmt.Printf("Input: %d - Valid: %t\n", i, resPart2)
		}
		if resPart2 {
			countPart2++
		}
	}
	fmt.Printf("(Part 1) Final count: %d\n", countPart1)
	fmt.Printf("(Part 2) Final count: %d\n", countPart2)
}
