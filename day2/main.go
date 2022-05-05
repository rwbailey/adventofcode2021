package main

import (
	"fmt"
	"strings"

	"github.com/rwbailey/advent2021/helpers"
)

func main() {
	part1()
	part2()
}

func part1() {
	data := helpers.GetInput("2")
	fmt.Println(coordProduct(data))
}

func part2() {
	data := helpers.GetInput("2")
	fmt.Println(aimCoordProduct(data))
}

func coordProduct(data string) int {
	X, Y := 0, 0
	list := helpers.MustStringList(data, "\n")
	for _, v := range list {
		d := strings.Split(v, " ")
		switch d[0] {
		case "forward":
			X += helpers.MustInt(d[1])
		case "down":
			Y += helpers.MustInt(d[1])
		case "up":
			Y -= helpers.MustInt(d[1])
		}
	}
	return X * Y
}

func aimCoordProduct(data string) int {
	x, y, aim := 0, 0, 0
	list := helpers.MustStringList(data, "\n")
	for _, v := range list {
		d := strings.Split(v, " ")
		switch d[0] {
		case "forward":
			x += helpers.MustInt(d[1])
			y += aim * helpers.MustInt(d[1])
		case "down":
			aim += helpers.MustInt(d[1])
		case "up":
			aim -= helpers.MustInt(d[1])
		}
	}
	return x * y
}
