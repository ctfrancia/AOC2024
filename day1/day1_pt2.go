package day1

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Day1Part2() (int, error) {
	// read from a file
	var score int
	f, err := os.Open("day1/input")
	if err != nil {
		return 0, err
	}
	defer f.Close()
	orderedLeft := []int{}
	orderedRight := []int{}

	// read each line of the file
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		col1 := strings.Split(line, "   ")[0]
		num1, err := strconv.Atoi(col1)
		if err != nil {
			return 0, err
		}
		orderedLeft = append(orderedLeft, int(num1))
		col2 := strings.Split(line, "   ")[1]
		num2, err := strconv.Atoi(col2)
		if err != nil {
			return 0, err
		}
		orderedRight = append(orderedRight, int(num2))
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// not the most efficient way to do this, but it works
	for i := 0; i < len(orderedLeft); i++ {
		score += similarityScore(orderedLeft[i], orderedRight)
	}

	return score, nil
}

func similarityScore(num int, list []int) int {
	multiplier := 0
	for _, n := range list {
		if n == num {
			multiplier += 1
		}
	}

	return num * multiplier
}
