package main

import "github.com/ctfrancia/AOC2024/day1"

func main() {
	total, err := day1.Day1()
	if err != nil {
		panic(err)
	}
	println(total)
}
