package main

import (
	"fmt"
	"os"
)

func main() {
	stringa := os.Args[1]
	sottostringhe := creaSottostringhe(stringa)
	ordine := creaOrdine(sottostringhe, stringa)

	// stampa la relazione delle sottostringhe con il respettivo numero di occorrenze
	fmt.Println(ordine)
	for _, sottostringa := range ordine {
		fmt.Printf("%s -> Occorrenze: %d\n", sottostringa, sottostringhe[sottostringa])
	}
}

func creaSottostringhe(stringa string) map[string]int {
	/* riceve la stringa input e restituice una mappa con tutte le sottostringhe che soddisfano
	le condizione e quindi devono essere analizzati e, forse, stampati */

	sottostringhe := make(map[string]int)
	for i := 0; i < len(stringa)-3; i++ {
		for j := len(stringa); j >= i; j-- {
			sottostringa := stringa[i:j]
			if condizioniSoddisfati(sottostringa) {
				sottostringhe[sottostringa]++
			}
		}
	}

	return sottostringhe
}

func condizioniSoddisfati(sottostringa string) bool {
	// define se la sottostringa soddisfa le condizioni di stampatura

	if (len(sottostringa) < 3) || (sottostringa[0] != sottostringa[len(sottostringa)-1]) {
		return false
	}
	return true
}

func creaOrdine(sottostringhe map[string]int, stringa string) (ordine []string) {
	// Crea una slice con l'ordine di stampatura

	// Riempi la slice ordine con le sottostringhe da ordinare
	for sottostringa, _ := range sottostringhe {
		ordine = append(ordine, sottostringa)
	}

	// ordina le sottostringhe nella slice ordine
	for i := 0; i < len(ordine); i++ {
		for j := 0; j < i; j++ {
			if len(ordine[i]) > len(ordine[j]) {
				ordine[i], ordine[j] = ordine[j], ordine[i]
			} else if len(ordine[i]) == len(ordine[j]) {
				for _, c := range stringa {
					if string(c) == ordine[j] {
						ordine[j], ordine[i] = ordine[i], ordine[j]
						break
					} else {
						continue					
					}
				}
			}
		}
	}

	return ordine
}
