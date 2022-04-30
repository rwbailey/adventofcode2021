package main

import (
	"fmt"

	"github.com/rwbailey/advent2021/helpers"
)

func main() {
	part1()
	part2()
}

func part1() {
	data := helpers.GetInput("1")
	fmt.Println(increase(data))
}

func part2() {
	data := helpers.GetInput("1")
	fmt.Println(increaseSlidingWindow(data))
}

func increase(data string) int {
	ints := helpers.MustIntList(data)
	increaseCount := 0
	for i := 1; i < len(ints); i++ {
		if ints[i] > ints[i-1] {
			increaseCount++
		}
	}
	return increaseCount
}

func increaseSlidingWindow(data string) int {
	ints := helpers.MustIntList(data)
	increaseCount := 0
	for i := 3; i < len(ints); i++ {
		if ints[i]+ints[i-1]+ints[i-2] > ints[i-1]+ints[i-2]+ints[i-3] {
			increaseCount++
		}
	}
	return increaseCount
}
