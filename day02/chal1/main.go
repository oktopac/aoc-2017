package main

import (
	"fmt"
	"os"

	day02 "github.com/oktopac/aoc-2017/day02/lib"
)

func main() {
	sheet, _ := day02.ParseSpreadsheet(os.Args[1])

	result := day02.Checksum(sheet)

	fmt.Println(result)
}
