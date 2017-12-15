package main

import (
	"fmt"

	day15 "github.com/oktopac/aoc-2017/day15/lib"
)

func main() {

	g1 := day15.GeneratorA(277, 40e6)
	g2 := day15.GeneratorB(349, 40e6)

	fmt.Println(day15.CountMatches(g1, g2))
}
