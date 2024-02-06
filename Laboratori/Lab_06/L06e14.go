package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	if ÈPalindroma(s) {
		fmt.Println("Palindroma")
	} else {
		fmt.Println("Non palindroma")
	}
}

func ÈPalindroma(s string) bool {
	s = strings.ToLower(s)
	r := []rune(s)
	j := len(r) - 1

	for i := range r {
		if i >= j {
			break
		}

		if r[i] != r[j] {
			return false
		}

		j--
	}
	return true
}
