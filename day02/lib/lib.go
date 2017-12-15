package day02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type spreadsheet [][]int

func ParseSpreadsheet(fname string) (spreadsheet, error) {
	fd, err := os.Open(fname)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ret := make(spreadsheet, 0)

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		digits := strings.Split(scanner.Text(), "\t")
		row := make([]int, len(digits))
		for i, d := range digits {
			dInt, err := strconv.Atoi(d)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			row[i] = dInt
		}
		ret = append(ret, row)
	}
	return ret, nil
}

func Checksum2(sheet spreadsheet) int {
	sum := 0
	for _, row := range sheet {
		lineResult := 0
		for _, item1 := range row {
			for _, item2 := range row {
				if item1 < item2 && item2%item1 == 0 {
					lineResult = item2 / item1
				}
			}
		}
		sum += lineResult
	}
	return sum
}

func Checksum(sheet spreadsheet) int {
	sum := 0
	for _, row := range sheet {
		min, max := row[0], row[0]
		for _, item := range row {
			if item > max {
				max = item
			}
			if item < min {
				min = item
			}
		}
		sum += max - min
	}
	return sum
}
