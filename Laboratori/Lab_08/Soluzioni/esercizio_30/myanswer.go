package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	inputString := os.Args[1]
	primi := tagliaStringa(inputString)
	for i := 0; i < len(primi); i++ {
		fmt.Println(primi[i])
	}
}

func Primo(n int) bool {
	if n == 0 || n == 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func tagliaStringa(s string) (primi []int) {
	for i := 0; i < len(s); i++ {
		for j := 1; j <= 3; j++ {
			if j+i > len(s) {
				continue
			}
			pezzo := s[:i] + s[i+j:]
			n, _ := strconv.Atoi(pezzo)
			if Primo(n) {
				primi = append(primi, n)
			}
		}
	}
	sort.Ints(primi)
	return primi
}
