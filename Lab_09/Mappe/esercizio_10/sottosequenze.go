package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	sequenza := os.Args[1:]

	sottosequenze := GeneraSottosequenze(sequenza)

	StampaSottosequenze(sottosequenze, len(sequenza))
}

func GeneraSottosequenze(sequenza []string) (sottosequenze map[string]int) {
	sottosequenze = make(map[string]int)
	for i, c1 := range sequenza {
		for j, c2 := range sequenza[i+1:] {
			if c1 == c2 {
				sottosequenze[strings.Join(sequenza[i:i+j+2], " ")]++
			}
		}
	}
	return
}

func StampaSottosequenze(sottosequenze map[string]int, lunghezzaMassima int) {

	for lunghezza := lunghezzaMassima; lunghezza >= 3; lunghezza-- {

		for sequenza, occorrenze := range sottosequenze {
			if len(sequenza) == 2*lunghezza-1 {
				fmt.Printf("%s -> Occorrenze: %d\n", sequenza, occorrenze)
			}
		}

	}
}
