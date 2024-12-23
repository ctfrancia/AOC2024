package main

import "github.com/ctfrancia/AOC2024/day1"

func main() {
	total, err := day1.Day1()
	if err != nil {
		panic(err)
	}
	println("day1: ", total)

	similarityScore, err := day1.Day1Part2()
	if err != nil {
		panic(err)
	}
	println("day1 part 2: ", similarityScore)
}
