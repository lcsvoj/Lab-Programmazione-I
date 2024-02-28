package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Print("Inserisci la stringa: ")
	fmt.Scan(&s)

	var k int
	fmt.Print("Inserisci il parametro di ripetizione: ")
	fmt.Scan(&k)

	fmt.Println("Nuova stringa: ", ripetirore(s, k))
}

func ripetirore(s string, k int) (sR string) {
	for _, c := range s {
		sR += strings.Repeat(string(c), k)
	}
	return sR
}
