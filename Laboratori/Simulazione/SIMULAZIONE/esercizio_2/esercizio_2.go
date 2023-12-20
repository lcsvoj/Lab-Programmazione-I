package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	stringa := os.Args[1]
	sottostringhe := creaSottostringhe(stringa)

	// conta quante volte ogni sottostringha appare
	occorrenze := make(map[string]int)
	for _, c := range sottostringhe {
		occorrenze[c]++		
	}
	
	// stampa la relazione delle sottostringhe con il respettivo numero di occorrenze
	for _, c := range sottostringhe {		
		fmt.Printf("%s -> Occorrenze: %d\n", c, occorrenze[c])
	}
}

func creaSottostringhe(stringa string) (sottostringhe []string) {
	/* riceve la stringa input e restituice una slice con tutte le sottostringhe che soddisfano
	le condizione e quindi devono essere analizzati e, forse, stampati */

	for i := 0; i < len(stringa)-3; i++ {
		for j := len(stringa); j >= i; j-- {
			sottostringa := stringa[i:j]
			if condizioniSoddisfati(sottostringa) {
				sottostringhe = append(sottostringhe, sottostringa)
			}
		}
	}

	sort.Slice(sottostringhe, func(i, j int) bool { return len(sottostringhe[i]) > len(sottostringhe[j]) })
	return sottostringhe
}

func condizioniSoddisfati(sottostringa string) bool {
	// define se la sottostringa soddisfa le condizioni di stampatura

	if (len(sottostringa) < 3) || (sottostringa[0] != sottostringa[len(sottostringa)-1]) {
		return false
	}
	return true
}

func contaOccorrenze(sottostringa string, sottostringhe []string) (occorrenze int) {
	for i := 0; i < len(sottostringhe); i++ {
		if sottostringhe[i] == sottostringa {
			occorrenze++
		}
	}
	return occorrenze
}
