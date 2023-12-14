package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	inputString := os.Args[1]
	fmt.Printf("\ninpuststring: %s (len: %d)\n", inputString, len(inputString))
	primi := tagliaStringa(inputString)
	for i := 0; i < len(primi); i++ {
		fmt.Println(primi[i])
	}
}

func Primo(n int) bool {
	fmt.Printf("\nchecking if %v is prime\n", n)
	if n == 0 || n == 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			fmt.Printf("%d is not prime because it's divisible by %d, returning false\n", n, i)
			return false
		}
	}
	fmt.Printf("%d is prime, returning true\n", n)
	return true
}

func tagliaStringa(s string) (primi []int) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("\ncutting the pieces of %s\n", s)
		fmt.Printf("========== cutting point at %d (that corresponds to %c) ========== \n", i, s[i])
		for j := 1; j <= 3; j++ {
			if j+i > len(s) {
				continue
			}
			fmt.Printf("\nthe cut size is %d\n", j)
			pezzo := s[:i] + s[i+j:]
			fmt.Printf("the result piece is %s\n", pezzo)
			n, _ := strconv.Atoi(pezzo)
			if Primo(n) {
				fmt.Printf("\nappending %s to primi slice\n", pezzo)
				primi = append(primi, n)
			}
		}
	}
	fmt.Printf("\n########## the final primi slice is %v ##########\n", primi)
	sort.Ints(primi)
	return primi
}
