package main

import (
	"fmt"
	"os"
)

type Sottostringa struct {
	stringa    string
	occorrenze int
}

var maggioreLunghezza int

func main() {
	// Leggi la stringa
	input := os.Args[1]
	sottostringhe := TrovaSottostringhe(input)

	lunghezzaDaStampare := maggioreLunghezza
	for i := 0; i < len(sottostringhe); i++ {
		for _, el := range sottostringhe {
			if len(el.stringa) == lunghezzaDaStampare {
				fmt.Printf("%v -> Occorrenze: %v\n", sottostringhe[el.stringa].stringa, sottostringhe[el.stringa].occorrenze)
			}
		}
		lunghezzaDaStampare--
	}

}

func TrovaSottostringhe(stringa string) map[string]Sottostringa {

	sottostringhe := make(map[string]Sottostringa, 0)

	for i := range stringa {
		for j := len(stringa); j > i; j-- {

			sotto := Sottostringa{stringa: stringa[i:j], occorrenze: 1}

			if len(sotto.stringa) > maggioreLunghezza {
				maggioreLunghezza = len(sotto.stringa)
			}

			if len(sotto.stringa) >= 3 && sotto.stringa[0] == sotto.stringa[len(sotto.stringa)-1] {

				if _, ok := sottostringhe[sotto.stringa]; ok {
					sotto.occorrenze++
				}
				sottostringhe[sotto.stringa] = sotto
			}
		}
	}
	return sottostringhe
}
