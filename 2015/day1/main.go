package day1

import (
	"fmt"
	"log"
	"os"
)

func floorChange(s string) int {
	d := map[string]int{"(": 1, ")": -1}

	f := 0
	for _, c := range s {
		f += d[string(c)]
	}
	return f
}

func floorChange2(s string) int {
	seq := "_"
	c := '_'

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			if s[i] == '{' {
				c = '}'
			} else if s[i] == '[' {
				c = ']'
			} else {
				c = ')'
			}
			seq = string(c) + seq
		} else {
			if seq[0] != s[i] {
				return false
			} else {
				seq = seq[1:]
			}
		}
	}

	if seq == "_" {
		return true
	}

	return false
}

// ( - up
// ) - down
func main() {
	fs, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(floorChange(string(fs)))
}
