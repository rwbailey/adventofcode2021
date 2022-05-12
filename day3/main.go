package main

import (
	"fmt"
	"strconv"

	"github.com/rwbailey/advent2021/helpers"
)

func main() {
	part1()
	part2()
}

func part1() {
	data := helpers.GetInput("3")
	fmt.Println(powerConsumption(data))
}

func part2() {
	data := helpers.GetInput("3")
	fmt.Println(lifeSupport(data))
}

func lifeSupport(data string) int64 {
	strings := helpers.MustStringList(data, "\n")

	o2, co2 := lsr(0, strings, strings)

	a, _ := strconv.ParseInt(o2[0], 2, 64)
	b, _ := strconv.ParseInt(co2[0], 2, 64)

	return a * b
}

func lsr(n int, o2, co2 []string) ([]string, []string) {
	if len(o2) == 1 && len(co2) == 1 {
		return o2, co2
	}
	// oxy, carb := o2, co2

	ones, zeroes := []string{}, []string{}

	if len(o2) > 1 {
		for _, v := range o2 {
			if string(v[n]) == "1" {
				ones = append(ones, v)
			} else if string(v[n]) == "0" {
				zeroes = append(zeroes, v)
			}
		}
		if len(zeroes) > len(ones) {
			o2 = zeroes
		} else {
			o2 = ones
		}
	}

	ones, zeroes = []string{}, []string{}

	if len(co2) > 1 {
		for _, v := range co2 {
			if string(v[n]) == "1" {
				ones = append(ones, v)
			} else if string(v[n]) == "0" {
				zeroes = append(zeroes, v)
			}
		}
		if len(ones) < len(zeroes) {
			co2 = ones
		} else {
			co2 = zeroes
		}
	}

	return lsr(n+1, o2, co2)
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
