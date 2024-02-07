package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var somma int
	numeri := LeggiNumeri()
	for _, n := range numeri {
		if Occorrenze(numeri, n) == 1 {
			somma += n
		} else {
			continue
		}
	}
	fmt.Println(somma)
}

func LeggiNumeri() (numeri []int) {
	for _, c := range os.Args {
		n, err := strconv.Atoi(c)
		if err == nil {
			numeri = append(numeri, n)
		} else {
			continue
		}
	}
	return numeri
}

func Occorrenze(numeri []int, n int) (occorrenze int) {
	for _, c := range numeri {
		if c == n {
			occorrenze++
		}
	}
	return occorrenze
}
