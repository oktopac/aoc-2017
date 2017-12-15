package day01

import (
	"strconv"
	"strings"
)

func solve(challengeArrayInt []int, step int) (int, error) {
	// Start with a sum of zero
	sum := 0

	for i := 0; i < len(challengeArrayInt); i++ {
		val := challengeArrayInt[i]
		compareVal := challengeArrayInt[(i+step)%len(challengeArrayInt)]
		if compareVal == val {
			sum += compareVal
		}
	}

	return sum, nil
}

func parseChallenge(challenge string) ([]int, error) {
	challengeArray := strings.Split(challenge, "")
	var challengeArrayInt []int

	for i := 0; i < len(challengeArray); i++ {
		val, err := strconv.Atoi(challengeArray[i])

		if err != nil {
			return nil, err
		}
		challengeArrayInt = append(challengeArrayInt, val)
	}
	return challengeArrayInt, nil
}

func Solve(challenge string) (int, error) {

	challengeArrayInt, err := parseChallenge(challenge)
	if err != nil {
		return 0, err
	}

	return solve(challengeArrayInt, 1)
}

func Solve2(challenge string) (int, error) {

	challengeArrayInt, err := parseChallenge(challenge)
	if err != nil {
		return 0, err
	}

	return solve(challengeArrayInt, len(challengeArrayInt)/2)
}
