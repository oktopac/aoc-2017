package main

import (
	day01 "aoc/day01/lib"
	"fmt"
	"log"
	"os"
)

func main() {
	// Take the input as the first argument on the command line
	argsWithoutProg := os.Args[1:]

	// Run the solver
	ret, err := day01.Solve(argsWithoutProg[0])

	if err != nil {
		log.Fatalln(err)
	}

	// Print the output
	fmt.Println(ret)
}
