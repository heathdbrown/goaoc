package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Shape interface {
	surfaceArea()
}

type RectagularPrism struct {
	l int
	w int
	h int
}

func (r RectagularPrism) surfaceArea() int {
	return 2 * ((r.l * r.w) + (r.w * r.h) + (r.h * r.l))
}

func (r RectagularPrism) smallestSideArea() int {
	sides := make([]int, 3)

	sides[0] = (r.l * r.w)
	sides[1] = (r.w * r.h)
	sides[2] = (r.l * r.h)

	return slices.Min(sides)
}

func (r RectagularPrism) smallestPerimeter() int {
	sides := make([]int, 3)

	sides[0] = r.perimeter([]int{r.h, r.l})
	sides[1] = r.perimeter([]int{r.h, r.w})
	sides[2] = r.perimeter([]int{r.l, r.w})

	return slices.Min(sides)
}

func (r RectagularPrism) perimeter(dimensions []int) int {
	a := dimensions[0]
	b := dimensions[1]

	return 2 * (a + b)
}

func (r RectagularPrism) volume() int {
	return r.l * r.w * r.h
}

func parseLine(line string) (int, int) {
	dimensions := strings.Split(line, "x")
	length, err := strconv.Atoi(dimensions[0])
	if err != nil {
		log.Print("error with conversion")
	}

	width, err := strconv.Atoi(dimensions[1])
	if err != nil {
		log.Print("error with conversion")
	}

	height, err := strconv.Atoi(dimensions[2])
	if err != nil {
		log.Print("error with conversion")
	}

	r := RectagularPrism{
		l: length,
		w: width,
		h: height,
	}
	return (r.surfaceArea() + r.smallestSideArea()), (r.smallestPerimeter() + r.volume())
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println("Error opening file")
		os.Exit(1)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	sumPaper := 0
	sumRibbon := 0
	for _, line := range lines {
		paper, ribbon := parseLine(string(line))
		sumPaper += paper
		sumRibbon += ribbon
	}

	fmt.Printf("Part1 Sum of Sq ft: %d\n", sumPaper)
	fmt.Printf("Part2 Sum of Ribbon: %d\n", sumRibbon)
}
