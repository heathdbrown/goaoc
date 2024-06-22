package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

type instruction struct {
	Action          string
	SubAction       string
	StartCoordinate Coordinate
	EndCoordinate   Coordinate
}

type instructions []instruction

func parseCoordinates(c []string) (int, int) {
	x, err := strconv.Atoi(c[0])
	if err != nil {
		log.Fatal("Error with string conversion")
	}

	y, err := strconv.Atoi(c[1])
	if err != nil {
		log.Fatal("Error with string conversion")
	}

	return x, y
}

func parseLine(line string) instruction {
	l := strings.Split(line, " ")
	i := instruction{Action: l[0]}

	switch l[0] {
	case "turn":
		// turn on 0,0 through 999,999
		// 0    1  2   3       4
		// A    SA SC  -       EC
		i.SubAction = l[1]
		sc := strings.Split(l[2], ",")

		i.StartCoordinate.x, i.StartCoordinate.y = parseCoordinates(sc)

		ec := strings.Split(l[4], ",")

		i.EndCoordinate.x, i.EndCoordinate.y = parseCoordinates(ec)
	case "toggle":
		// toggle 0,0 through 999,0
		// 0      1   2       3
		sc := strings.Split(l[1], ",")

		i.StartCoordinate.x, i.StartCoordinate.y = parseCoordinates(sc)
		ec := strings.Split(l[3], ",")

		i.EndCoordinate.x, i.EndCoordinate.y = parseCoordinates(ec)
	}

	return i
}

func main() {
	fmt.Println("2015 day6")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error with opening issues")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var i instructions
	for scanner.Scan() {
		i = append(i, parseLine(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, order := range i {
		fmt.Println(order)
	}
}
