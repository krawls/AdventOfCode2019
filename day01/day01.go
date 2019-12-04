package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// global variables
var debugMessages = false // turn on verbose debug messages

func calcModuleFuelRequirementPart1(mass int) int {
	return mass/3 - 2
}

func calcModuleFuelRequirementPart2(mass int) int {
	// using given mass, calculate fuel requirement
	// then, calculate fuel requirement of the fuel
	// keep calculating until fuel requirement is <= 0

	var total int

	for mass > 0 {
		// newMass allows us to "look forward" one calculation to see
		// if we need to make the calculation, or break and return the
		// current result
		newMass := mass/3 - 2
		if newMass <= 0 {
			break
		}
		total += newMass
		mass = newMass
	}
	return total
}

func main() {
	var totalFuelReqPart1 int
	var totalFuelReqPart2 int
	var inputFilePath = "input.txt"

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		// cumulatively add each module's fuel requirement, after calculating from
		// input file's given mass
		totalFuelReqPart1 += calcModuleFuelRequirementPart1(mass)
		totalFuelReqPart2 += calcModuleFuelRequirementPart2(mass)

		// optionally show verbose messages
		if debugMessages {
			fmt.Printf("(Part 1) Fuel required for mass of %v is %v\n", mass, calcModuleFuelRequirementPart1(mass))
			fmt.Printf("(Part 1) Current total fuel required: %v\n", totalFuelReqPart1)
			fmt.Printf("(Part 2) Fuel required for mass of %v is %v\n", mass, calcModuleFuelRequirementPart2(mass))
			fmt.Printf("(Part 2) Current total fuel required: %v\n", totalFuelReqPart2)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// show the final answers
	fmt.Printf("(Part 1) Total fuel required: %v\n", totalFuelReqPart1)
	fmt.Printf("(Part 2) Total fuel required: %v\n", totalFuelReqPart2)
}
