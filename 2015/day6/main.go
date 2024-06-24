package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// based on the following https://www.kaggle.com/code/jaeihn/advent-of-code-2015-day-6-solutions
func main() {
	lights := make([]bool, 1000000)
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error with reading file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	r, err := regexp.Compile(`\d+`)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		// [222 12 856 241]
		coord := r.FindAllString(scanner.Text(), -1)

		x_i, err := strconv.Atoi(coord[0])
		if err != nil {
			log.Fatal("Error with string conversion to integer")
		}
		y_i, err := strconv.Atoi(coord[1])
		if err != nil {
			log.Fatal("Error with string conversion to integer")
		}
		x_f, err := strconv.Atoi(coord[2])
		if err != nil {
			log.Fatal("Error with string conversion to integer")
		}
		y_f, err := strconv.Atoi(coord[3])
		if err != nil {
			log.Fatal("Error with string conversion to integer")
		}

		switch {
		case strings.HasPrefix(scanner.Text(), "turn on"):
			for i := x_i; i <= x_f; i++ {
				for j := y_i; j <= y_f; j++ {
					lights[j*1000+i] = true
				}
			}
		case strings.HasPrefix(scanner.Text(), "turn off"):
			for i := x_i; i <= x_f; i++ {
				for j := y_i; j <= y_f; j++ {
					lights[j*1000+i] = false
				}
			}
		case strings.HasPrefix(scanner.Text(), "toggle"):
			for i := x_i; i <= x_f; i++ {
				for j := y_i; j <= y_f; j++ {
					lights[j*1000+i] = !lights[j*1000+i]
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0
	for i := 0; i < 1000000; i++ {
		if lights[i] == true {
			count += 1
		}
	}
	// for line := range string(file) {
	// 	fmt.Println(line)
	// }

	// var lights [100000000]bool

	// fmt.Println(lights)
	fmt.Println(count)
}
