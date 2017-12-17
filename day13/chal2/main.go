package main

import (
	"fmt"
	"os"

	day13 "github.com/oktopac/aoc-2017/day13/lib"
)

func main() {
	wait := day13.LowestSeverity(os.Args[1])
	fmt.Println(wait)
}
