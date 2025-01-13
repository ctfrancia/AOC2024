package day2

import (
	"bufio"
	"fmt"
	"github.com/ctfrancia/AOC2024/helpers"
	"math"
	"strconv"
	"strings"
)

func Day2() (int, error) {
	safeReports := 0
	f, err := helpers.ReadFile("../day2/input.txt")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// read each line of the file
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		report := scanner.Text()
		levels := strings.Split(report, " ")

		// Convert string slices to integers
		var numbers []int
		for _, level := range levels {
			current, err := strconv.Atoi(level)
			if err != nil {
				fmt.Println("Error converting string to int:", level, err)
				continue
			}
			numbers = append(numbers, current)
		}
		// Flag to track if the report is safe
		isSafe := true
		var direction int // 0 means no direction, 1 means increasing, -1 means decreasing

		// Loop through the numbers and check the difference between consecutive numbers
		for i := 0; i < len(numbers)-1; i++ {
			current := numbers[i]
			next := numbers[i+1]

			// Calculate the difference
			diff := math.Abs(float64(next - current))

			// If difference is 0 or greater than 2, it's unsafe
			if diff < 1 || diff > 3 {
				isSafe = false
				break
			}

			// Track direction: if direction has changed from increasing to decreasing or vice versa, it's unsafe
			if next > current {
				if direction == -1 { // previous was decreasing
					isSafe = false
					break
				}
				direction = 1
			} else if next < current {
				if direction == 1 { // previous was increasing
					isSafe = false
					break
				}
				direction = -1
			}
		}

		// If the report is safe, increment safeReports
		if isSafe {
			safeReports++
		}
	}

	return safeReports, nil
}
