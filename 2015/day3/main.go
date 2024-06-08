package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Directions map[string][]int

type node struct {
	data House
	next *node
}

type linkedList struct {
	head   *node
	length int
}

type House struct {
	location     []int
	visited      bool
	nextLocation []int
	prevLocation []int
}

func (l linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// func (h House) next() {
// }

// func (h House) prev() {
// }

func parseLine(line string) []string {
	return strings.Split(line, "")
}

func countHouses(directions string) int {
	dmap := map[string]int{"^": 1, "v": -1, "<": -1, ">": 1}
	count := 0
	dsum := 0
	d := parseLine(directions)
	for i, h := range d {
		if i == 0 {
			count++
		}
		count += dmap[h]
		dsum += dmap[h]
	}

	if dsum == 0 {
		return count + 1
	}
	return count
}

func (h House) applyDirection(d []int) {
	h.nextLocation[0] += d[0]
	h.nextLocation[1] += d[1]
}

// https://github.com/obalunenko/advent-of-code/blob/master/internal/puzzles/solutions/2015/day03/solution.go#L63
type grid struct {
	x int
	y int
}

func newGrid() grid {
	return grid{
		x: 0,
		y: 0,
	}
}

func main() {
	fmt.Println("2015 day3")

	directions := Directions{"^": []int{1, 0}, "v": []int{-1, 0}, ">": []int{0, 1}, "<": []int{0, -1}}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Panicln(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	// Starting line
	// houses := []House{}
	// for _, line := range lines {
	// 	for i, d := range parseLine(line) {
	// 		if i == 0 {
	// 			newHouse := House{location: []int{0, 0}, visited: true, nextLocation: []int{0, 0}, prevLocation: []int{0, 0}}
	// 			newHouse.applyDirection(directions[d])
	// 			fmt.Println(directions[d])
	// 			houses = append(houses, newHouse)
	// 		}
	// 	}
	// }

	// New line
	houses := []int{0, 0}
	for _, line := range lines {
		for i, d := range parseLine(line) {
			if i == 0 {
				fmt.Println("starting point")
			}
			houses = append(houses, directions[d]...)
		}
	}

	fmt.Printf("%+v", houses)
}
