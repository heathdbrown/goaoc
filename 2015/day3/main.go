package main

import (
	"fmt"
	"os"
	"strings"
)

// https://dev.to/10xlearner/advent-of-code-perfectly-spherical-houses-in-a-vacuum-puzzle-3-n78
/*
Open file, read string into list of strings for directions

go through directions and capture the Coordinates in the Path

sort the path and look for unique Coordinates

the length of the unique path is the number of houses
*/

type Coordinate struct {
	x int
	y int
}

type DeliveryMan struct {
	position Coordinate
	path     Path
}

type Path []Coordinate

func parseLine(s string) []string {
	return strings.Split(s, "")
}

func unique(p Path) Path {
	var unique Path
	m := map[Coordinate]bool{}

	for _, v := range p {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}

func countHouses(directions []string) int {
	var santaPosition Coordinate

	var santaPath Path

	santaPath = append(santaPath, santaPosition)
	for _, d := range directions {
		switch d {
		case "^":
			santaPosition.y++
		case "v":
			santaPosition.y--
		case ">":
			santaPosition.x++
		case "<":
			santaPosition.x--

		}
		santaPath = append(santaPath, santaPosition)
	}

	return len(unique(santaPath))
}

func countSantaAndRoboHouses(directions []string) int {
	var santa, roboSanta DeliveryMan

	var presentPath Path

	santa.path = append(santa.path, santa.position)
	roboSanta.path = append(roboSanta.path, roboSanta.position)
	deliveryMan := &santa
	for _, d := range directions {
		switch d {
		case "^":
			deliveryMan.position.y++

		case "v":
			deliveryMan.position.y--

		case ">":
			deliveryMan.position.x++
		case "<":
			deliveryMan.position.x--

		}
		deliveryMan.path = append(deliveryMan.path, deliveryMan.position)
		switch {
		case deliveryMan == &santa:
			deliveryMan = &roboSanta
		case deliveryMan == &roboSanta:
			deliveryMan = &santa
		}

	}
	presentPath = append(presentPath, santa.path...)
	presentPath = append(presentPath, roboSanta.path...)
	return len(unique(presentPath))
}

func main() {
	fmt.Println("2015_day3")

	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Errorf("Erroring reading file")
	}

	directions := parseLine(string(file))

	fmt.Printf("Part1:%d\n", countHouses(directions))
	fmt.Printf("Part2:%d\n", countSantaAndRoboHouses(directions))
}
