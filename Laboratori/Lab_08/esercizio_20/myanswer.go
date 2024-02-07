package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// get the sequence from CLI input
	sequenza := make([]int, len(os.Args)-1)
	for i, c := range os.Args[1:] {
		sequenza[i], _ = strconv.Atoi(c)
	}

	fmt.Println("La somma è", Calcola(sequenza))
}

func Calcola(s1 []int) (somma int) {
	for i, _ := range s1 {
		if i%2 == 0 {
			somma += s1[i] * s1[i+1]
		}
	}
	return somma
}
