package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func GetMD5Hash(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func CheckLeadingZeros(s string, z int) bool {
	l := s[:z]
	n := strings.Repeat("0", z)
	if l != n {
		return false
	}
	return true
}

func CheckNumber(key string, zeros int) int {
	var n int
	for i := 0; true; i++ {
		iStr := strconv.Itoa(i)
		h := GetMD5Hash(key + iStr)
		if !CheckLeadingZeros(h, zeros) {
			continue
		} else {
			n = i
			break
		}
	}
	return n
}

func main() {
	key := flag.String("key", "abcdef", "Prefix string to generage md5 hashes")
	flag.Parse()
	fmt.Printf("Part1: %d\n", CheckNumber(*key, 5))
	fmt.Printf("Part2: %d\n", CheckNumber(*key, 6))
}
