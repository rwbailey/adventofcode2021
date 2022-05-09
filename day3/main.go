package main

import (
	"fmt"
	"strconv"

	"github.com/rwbailey/advent2021/helpers"
)

func main() {
	part1()
	// part2()
}

func part1() {
	data := helpers.GetInput("3")
	fmt.Println(powerConsumption(data))
}

func part2() {
	// data := helpers.GetInput("3")
	// fmt.Println()
}

func powerConsumption(data string) int64 {
	strings := helpers.MustStringList(data, "\n")

	gamma, epsilon := "", ""
	for i := 0; i < 12; i++ {
		ones, zeroes := 0, 0
		for _, v := range strings {
			if string(v[i]) == "1" {
				ones++
			} else if string(v[i]) == "0" {
				zeroes++
			}
		}
		if ones > zeroes {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)

	return g * e
}
