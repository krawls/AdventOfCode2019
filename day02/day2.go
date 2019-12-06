package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// global variables
var debugMessages = false // turn on verbose debug messages

func add(s *[]int, pos int) {
	// input 1 position offset by +1
	i1 := (*s)[pos+1]

	// input 2 position offset by +2
	i2 := (*s)[pos+2]

	// output position offset by +3
	o := (*s)[pos+3]

	// write value into output location in program pointer
	(*s)[o] = (*s)[i1] + (*s)[i2]
}

func multiply(s *[]int, pos int) {
	// input 1 position offset by +1
	i1 := (*s)[pos+1]

	// input 2 position offset by +2
	i2 := (*s)[pos+2]

	// output position offset by +3
	o := (*s)[pos+3]

	// write value into output location in program pointer
	(*s)[o] = (*s)[i1] * (*s)[i2]
}

func process(s *[]int) {
	// process opcode program
	for i := 0; i < len(*s); i = i + 4 {
		switch (*s)[i] {
		case 1:
			if debugMessages {
				fmt.Printf("Opcode: Add")
			}
			add(s, i)
		case 2:
			if debugMessages {
				fmt.Printf("Opcode: Multiply")
			}
			multiply(s, i)
		case 99:
			if debugMessages {
				fmt.Printf("Opcode: Halt")
			}
			return
		default:
			if debugMessages {
				fmt.Printf("Opcode: UNKNOWN")
			}
			//panic(fmt.Sprintf("Opcode: UNKNOWN"))
			return
		}
	}
}

func main() {

	var inputFilePath = "input.txt"

	b, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// tokenize string into slice
	s := strings.Split(string(b), ",")

	// convert []string to []int to save casting headache down the road
	progInit := []int{}

	for _, i := range s {
		j, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		progInit = append(progInit, j)
	}

	// copy initial program into part 1 program to save a new fresh instance for use in part 2
	progPart1 := make([]int, len(progInit))
	copy(progPart1, progInit)

	// (Part 1) replace position 1 with value 12 and replace position 2 with value 2
	progPart1[1] = 12
	progPart1[2] = 2

	process(&progPart1)

	if debugMessages {
		fmt.Printf("Program: %v", progPart1)
	}

	// show me the answer
	fmt.Printf("(Part 1) Value at position 0: %v\n", progPart1[0])

	// (Part 2) brute force to find values for positions 1 and 2 to get desired output
	// max values are length of program
	progPart2 := make([]int, len(progInit))
	output := 19690720
	max := len(progInit)
	done := false
	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			// important: reinitialize program every try
			copy(progPart2, progInit)
			progPart2[1] = i
			progPart2[2] = j
			process(&progPart2)
			if progPart2[0] == output {
				done = true
			}
			if done {
				break
			}
		}
		if done {
			break
		}
	}
	//show me the answers
	fmt.Printf("(Part 2) Value at position 0: %v\n", progPart2[0])
	fmt.Printf("(Part 2) Noun: %v Verb: %v\n", progPart2[1], progPart2[2])
	fmt.Printf("(Part 2) Answer: 100*noun+verb: %v\n", 100*progPart2[1]+progPart2[2])
}
