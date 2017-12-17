package main

import (
	"fmt"
	"os"

	day13 "github.com/oktopac/aoc-2017/day13/lib"
)

func main() {
	severity := day13.Severity(os.Args[1])
	fmt.Println(severity)
}
