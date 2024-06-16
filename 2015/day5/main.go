package main

import (
	"fmt"
	"strings"
)

func checkVowels(s string) bool {
	count := 0

	for _, l := range s {
		if strings.ContainsAny(string(l), "aeiou") {
			count++
		}
	}

	return count <= 3
}

func nice(s string) bool {
	r := true

	switch {
	case strings.Contains(s, "ab"):
		r = false
	case strings.Contains(s, "cd"):
		r = false
	case strings.Contains(s, "pq"):
		r = false
	case strings.Contains(s, "xy"):
		r = false
	case checkVowels(s):
		r = true
	}

	return r
}

func main() {
	fmt.Println("2015 day5")
}
