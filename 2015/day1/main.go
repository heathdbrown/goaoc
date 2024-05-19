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
// ( - up
// ) - down
func main() {
	fs, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(floorChange(string(fs)))
}
