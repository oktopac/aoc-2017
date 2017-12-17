package main

import (
	"fmt"
	"os"

	day07 "github.com/oktopac/aoc-2017/day07/lib"
)

func main() {
	// Take the input as the first argument on the command line
	argsWithoutProg := os.Args[1:]

	// Run the solver
	ret := day07.FindTop(argsWithoutProg[0])

	fmt.Println(ret)
}
