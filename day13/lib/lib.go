package day13

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func mustConv(in string) uint {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return uint(out)
}

func parseFirewall(fname string) (map[uint]uint, error) {
	fd, err := os.Open(fname)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ret := make(map[uint]uint, 0)

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		digits := strings.Split(scanner.Text(), ": ")
		ret[mustConv(digits[0])] = mustConv(digits[1])
	}

	return ret, nil
}

func doesHitScore(layerDepth uint, layerRange uint, offset uint) uint {
	layerFrequency := 2 * (layerRange - 1)

	scannerIndexAtDepth := (layerDepth + offset) % layerFrequency

	if scannerIndexAtDepth == 0 {
		return layerDepth * layerRange
	}
	return 0
}

func Severity(fname string, offset uint) uint {
	firewall, err := parseFirewall(fname)
	if err != nil {
		panic(err)
	}
	sum := uint(0)
	for layerDepth, layerRange := range firewall {
		sum += doesHitScore(layerDepth, layerRange, offset)
	}
	return sum
}

func LowestSeverity(fname string) uint {
	delay := uint(0)
	for delay = uint(0); delay < 1000000; delay++ {
		severity := Severity(fname, delay)
		if severity == 0 {
			break
		}
	}
	return delay
}
