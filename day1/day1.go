package day1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
- sort the left and right column smallest to biggest
- compare the smallest of the left column to the smallest of the right column
- create a third column with the difference between the two
- add column 3 to the total
*/

func Day1() (int, error) {
	// read from a file
	f, err := os.Open("day1/input")
	if err != nil {
		fmt.Println("Error reading file")
		return 0, err
	}
	defer f.Close()

	// read each line of the file
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		col1 := strings.Split(line, "   ")[0]
		col2 := strings.Split(line, "   ")[1]

		fmt.Println("1", col1)
		fmt.Println("2", col2)
	}

	return 0, nil
}
