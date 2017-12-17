package day07

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadList(fname string) (map[string]string, error) {
	fd, err := os.Open(fname)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ret := make(map[string]string, 0)

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		parent := strings.Split(line[0], " ")[0]
		if len(line) == 1 {

		} else if len(line) == 2 {
			children := strings.Split(line[1], ", ")
			for _, child := range children {
				ret[child] = parent
			}
		} else {
			log.Fatalf("Oddness on line %s", line)
			return nil, nil
		}
	}
	return ret, nil
}

func FindTop(fname string) string {
	tree, err := ReadList(fname)
	if err != nil {
		panic(err)
	}
	var key string
	for key = range tree {
		break
	}
	for {
		if val, ok := tree[key]; ok {
			key = val
			continue
		}
		break
	}
	return key
}
