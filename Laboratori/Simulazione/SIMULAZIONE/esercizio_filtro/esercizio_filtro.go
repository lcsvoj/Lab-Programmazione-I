package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	for i := 0; i < len(input); i++ {
		if i%2 == 1 {
			ripetizione, _ := strconv.Atoi(input[i])
			for j := 0; j < ripetizione; j++ {
				fmt.Print(input[i-1])
			}
		}
	}
	fmt.Println()
}
